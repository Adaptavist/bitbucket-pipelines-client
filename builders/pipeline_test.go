package builders

import "testing"

func TestPipelineVariable(t *testing.T) {
	p := Pipeline().
		Variable("VAR_NAME", "Value", true).
		Build()

	if len(*p.Variables) != 1 {
		t.Error("Pipeline should a variable")
	}
}

func TestPipelineVariables(t *testing.T) {
	p := Pipeline().
		Variable("VAR_NAME_1", "Value1", true).
		Variable("VAR_NAME_2", "Value2", true).
		Build()

	if len(*p.Variables) != 2 {
		t.Error("Pipeline should have 2 variables")
	}
}
