package builders

import "testing"

func TestTargetTag(t *testing.T) {
	p := Target().Tag("v1", "commit-hash").Build()

	if p.Commit == nil {
		t.Errorf("p.Commit should not be nil when using Tag")
	}
}
