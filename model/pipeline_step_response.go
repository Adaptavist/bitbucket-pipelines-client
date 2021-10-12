package model

import "encoding/json"

// PipelineStep on a running/complete spec
type PipelineStep struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	State State  `json:"state"`
}

// String representation of a Step
func (s PipelineStep) String() string {
	return s.Name
}

// ToJSON marshals the Step to JSON
func (s PipelineStep) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}

// StepsResponse response
type StepsResponse struct {
	Page   int           `json:"page"`
	Values PipelineSteps `json:"values"`
	Next   *string       `json:"next"`
}

// PipelineSteps is a list of steps
type PipelineSteps []PipelineStep

// Filter Steps using a callable function
func (s PipelineSteps) Filter(test func(step PipelineStep) bool) PipelineSteps {
	return FilterSteps(s, test)
}

// FilterSteps filters using a callable function
func FilterSteps(steps PipelineSteps, test func(PipelineStep) bool) (results PipelineSteps) {
	for _, s := range steps {
		if test(s) {
			results = append(results, s)
		}
	}
	return
}
