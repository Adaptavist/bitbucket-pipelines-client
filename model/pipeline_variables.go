package model

// PipelineVariable for a BitBucket spec
type PipelineVariable struct {
	Key     string `json:"key" yaml:"key"`
	Value   string `json:"value" yaml:"value"`
	Secured bool   `json:"secured" yaml:"secured"`
}

// PipelineVariables for a BitBucket spec
type PipelineVariables []PipelineVariable

func Variable(key string, value string, secured bool) PipelineVariable {
	return PipelineVariable{key, value, secured}
}
