package builders

import "github.com/adaptavist/bitbucket_pipelines_client/model"

type target struct {
	model.PipelineTarget
}

func (t *target) Ref(refType, name string) *target {
	t.PipelineTarget.Type = "pipeline_ref_target"
	t.PipelineTarget.RefType = refType
	t.PipelineTarget.RefName = name
	return t
}

func (t *target) Tag(name, commit string) *target {
	return t.Ref("tag", name).Commit(commit)
}

func (t *target) Branch(name string) *target {
	return t.Ref("branch", name)
}

func (t *target) Commit(commit string) *target {
	t.PipelineTarget.Commit = &model.PipelineTargetCommit{
		Type: "commit",
		Hash: commit,
	}
	return t
}

func (t *target) Pattern(pattern string) *target {
	t.Selector = &model.PipelineTargetSelector{
		Type:    "custom",
		Pattern: pattern,
	}
	return t
}

func (t *target) Build() model.PipelineTarget {
	return t.PipelineTarget
}

func Target() *target {
	return new(target)
}
