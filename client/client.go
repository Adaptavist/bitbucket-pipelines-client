package client

import (
	"fmt"
	"github.com/adaptavist/bitbucket_pipelines_client/model"
	http "github.com/hashicorp/go-retryablehttp"
	"time"
)

type Client struct {
	Config *Config
	client *http.Client
}

// PostPipeline will post a pipeline but will not wait for the pipeline to complete.
func (h Client) PostPipeline(request model.PostPipelineRequest) (pipeline *model.Pipeline, err error) {
	workspace, err := h.Config.GetWorkspace(request.Workspace)
	if err != nil {
		return pipeline, err
	}

	repository, err := h.Config.GetRepository(request.Repository)
	if err != nil {
		return pipeline, err
	}

	url := fmt.Sprintf("%s/2.0/repositories/%v/%v/pipelines/", h.Config.GetBaseURL(), *workspace, *repository)
	err = h.postUnmarshalled(url, request.Pipeline, &pipeline)
	if err != nil {
		return pipeline, err
	}

	return pipeline, err
}

// RunPipeline will POST the pipeline but will monitor its status until completion.
func (h Client) RunPipeline(request model.PostPipelineRequest) (pipeline *model.Pipeline, err error) {
	pipeline, err = h.PostPipeline(request)

	if err != nil {
		return
	}

	for {
		if pipeline.State.Result.OK() {
			break
		}
		time.Sleep(time.Second * 5) // TODO
		pipeline, err = h.GetPipeline(model.GetPipelineRequest{
			Workspace:  request.Workspace,
			Repository: request.Repository,
			Pipeline:   pipeline,
		})
	}

	return
}

// GetPipeline gets current status of a running pipeline
func (h Client) GetPipeline(request model.GetPipelineRequest) (pipeline *model.Pipeline, err error) {
	workspace, err := h.Config.GetWorkspace(request.Workspace)
	if err != nil {
		return pipeline, err
	}

	repository, err := h.Config.GetRepository(request.Repository)
	if err != nil {
		return pipeline, err
	}

	url := fmt.Sprintf("%s/2.0/repositories/%v/%v/pipelines/%v", h.Config.GetBaseURL(), *workspace, *repository, *request.Pipeline.UUID)
	err = h.getUnmarshalled(url, &pipeline)
	if err != nil {
		return nil, err
	}
	return
}

// GetPipelineSteps will return all steps and their status associated with a given pipeline.
func (h Client) GetPipelineSteps(request model.GetPipelineRequest) (steps model.PipelineSteps, err error) {
	workspace, err := h.Config.GetWorkspace(request.Workspace)
	if err != nil {
		return steps, err
	}

	repository, err := h.Config.GetRepository(request.Repository)
	if err != nil {
		return steps, err
	}

	url := fmt.Sprintf("%s/2.0/repositories/%v/%v/pipelines/%v/steps/", h.Config.GetBaseURL(), *workspace, *repository, *request.Pipeline.UUID)
	var result model.StepsResponse
	err = h.getUnmarshalled(url, &result)
	steps = append(steps, result.Values...)
	if err != nil {
		return
	}

	for {
		if result.Next == nil || *result.Next == "" {
			break
		}
		err = h.getUnmarshalled(*result.Next, &result)
		if err != nil {
			return
		}
		steps = append(steps, result.Values...)
	}
	return
}

func (h Client) GetPipelineStepLog(request model.GetPipelineStepRequest) ([]byte, error) {
	w, r, err := h.Config.GetWorkspaceAndRepository(request.Workspace, request.Repository)

	if err != nil {
		return []byte{}, err
	}

	url := fmt.Sprintf("%s/2.0/repositories/%s/%s/pipelines/%s/steps/%s/log", h.Config.GetBaseURL(), *w, *r, *request.Pipeline.UUID, request.PipelineStep.UUID)
	response, err := h.get(url)
	return response, err
}

func (h Client) GetTag(request model.GetTagRequest) (response *model.TagResponse, err error) {
	w, r, err := h.Config.GetWorkspaceAndRepository(request.Workspace, request.Repository)
	if err != nil {
		return nil, err
	}
	url:= fmt.Sprintf("%s/2.0/repositories/%s/%s/refs/tags/%s", h.Config.GetBaseURL(), *w, *r, request.Tag)
	err = h.getUnmarshalled(url, &response)
	return
}