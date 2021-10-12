package builders

import "github.com/adaptavist/bitbucket_pipelines_client/model"

type pipeline struct {
	model.Pipeline
}

func (b *pipeline) Target(t model.PipelineTarget) *pipeline {
	b.Pipeline.Target = &t
	return b
}

func (b *pipeline) Variables(v model.PipelineVariables) *pipeline {
	b.Pipeline.Variables = &v
	return b
}

func (b *pipeline) Variable(key, value string, secured bool) *pipeline {
	n := append(*b.Pipeline.Variables, model.Variable(key, value, secured))
	b.Pipeline.Variables = &n
	return b
}

func (b *pipeline) Build() *model.Pipeline {
	return &b.Pipeline
}

func Pipeline() *pipeline {
	return &pipeline{Pipeline: model.Pipeline{Variables: &model.PipelineVariables{}}}
}
