package model

import "fmt"

const (
	RefTypeBranch = "branch"
	RefTypeTag    = "tag"
)

// PipelineTarget spec for a Pipeline to run
type PipelineTarget struct {
	Type     string                  `json:"type,omitempty"`
	RefType  string                  `json:"ref_type,omitempty"`
	RefName  string                  `json:"ref_name,omitempty"`
	Selector *PipelineTargetSelector `json:"selector,omitempty"`
	Commit   *PipelineTargetCommit   `json:"commit,omitempty"`
}

type PipelineTargetCommit struct {
	Type string `json:"type"`
	Hash string `json:"hash"`
}

// PipelineTargetSelector description what branch/tag and spec we are to execute
type PipelineTargetSelector struct {
	Type    string `json:"type"`
	Pattern string `json:"pattern"`
}

func (t PipelineTarget) String() string {
	return fmt.Sprintf("%s/%s", t.RefName, t.Selector.Pattern)
}

func (t PipelineTarget) GetTargetDescriptor() string {
	return t.String()
}
