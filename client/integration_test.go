// +build integration

package client

import (
	"github.com/adaptavist/bitbucket_pipelines_client/model"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"

	"github.com/adaptavist/bitbucket_pipelines_client/builders"
)

const testWorkspace = "workspace"
const testRepoSlug = "bitbucket-pipelines-runner"
const testTag = "v0.0.0"
const testCommit = "b28c6cad5a664eb1a5fcd49c3bb256dd426cf731"
const testPattern = "example"

func envOrFail(key string) string {
	if os.Getenv(key) == "" {
		log.Fatatf("key not found: %s", key)
	}
	return os.Getenv(key)
}

func envOptional(key string, f func(s string)) {
	if os.Getenv(key) == "" {
		f(os.Getenv(key))
	}
}

func testClient() (client *Client) {
	client = &Client{Config: &Config{}}
	client.Config.Username = os.Getenv("BITBUCKET_USERNAME")
	client.Config.Password = os.Getenv("BITBUCKET_PASSWORD")
	envOptional("BITBUCKET_BASE_URL", func(s string) { client.Config.BaseURL = s })
	envOptional("BITBUCKET_WORKSPACE", func(s string) { client.Config.Workspace = s })
	envOptional("BITBUCKET_REPOSITORY", func(s string) { client.Config.Repository = s })
	return
}

func testTarget() model.PipelineTarget {
	return builders.Target().
		Tag(testTag, testCommit).
		Pattern(testPattern).
		Build()
}

func testPayload() model.Pipeline {
	return builders.Pipeline().Target(testTarget()).Build()
}

// TestIntegration runs a full pipeline and grabs the pipeline output
func TestIntegration(t *testing.T) {
	client := testClient()

	pipeline, err := client.RunPipeline(model.PostPipelineRequest{
		Workspace:  &testWorkspace,
		Repository: &testRepoSlug,
		Pipeline:   testPayload(),
	})

	assert.Nil(t, err, "failed to run pipeline")
	assert.NotNil(t, pipeline.UUID, "UUID expected")
}
