package model

type PostPipelineRequest struct {
	Workspace *string
	Repository *string
	*Pipeline
}

type GetPipelineRequest struct {
	Workspace *string
	Repository *string
	*Pipeline
}

type GetPipelineStepRequest struct {
	Workspace *string
	Repository *string
	*Pipeline
	*PipelineStep
}

type GetTagRequest struct {
	Workspace *string
	Repository *string
	Tag string
}