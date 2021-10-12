package model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// Pipeline containing details we need to start a pipeline
type Pipeline struct {
	UUID        *string            `json:"uuid,omitempty"`
	BuildNumber *int               `json:"build_number,omitempty"`
	CompletedOn *string            `json:"completed_on,omitempty"`
	State       *State             `json:"state,omitempty"`
	Target      *PipelineTarget    `json:"target,omitempty"`
	Variables   *PipelineVariables `json:"variables,omitempty"`
}

func (p Pipeline) Equal(c Pipeline) bool {
	return reflect.DeepEqual(p, c)
}

func (p Pipeline) Empty() bool {
	return p.Equal(Pipeline{})
}

func (p Pipeline) String() string {
	return strings.ToLower(fmt.Sprintf("%s %s", *p.UUID, *p.State))
}

func (p Pipeline) ToJSON() ([]byte, error) {
	return json.Marshal(p)
}
