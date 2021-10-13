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

func envOrFail(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("key not found: %s", key)
	}
	return os.Getenv(key)
}

func envOptional(key string, f func(s string)) {
	if os.Getenv(key) != "" {
		f(os.Getenv(key))
	}
}

func testClient() (client *Client) {
	client = &Client{Config: &Config{}}
	client.Config.Username = os.Getenv("BITBUCKET_USERNAME")
	client.Config.Password = os.Getenv("BITBUCKET_PASSWORD")
	envOptional("BITBUCKET_BASE_URL", func(s string) { client.Config.BaseURL = s })
	envOptional("BITBUCKET_WORKSPACE", func(s string) { client.Config.Workspace = &s })
	envOptional("BITBUCKET_REPOSITORY", func(s string) { client.Config.Repository = &s })
	return
}

// TestIntegration runs a full pipeline and grabs the pipeline output
func TestIntegration(t *testing.T) {
	client := testClient()

	targetBuilder := builders.Target()
	envOptional("BITBUCKET_TAG", func(s string) {
		r, e := client.GetTag(model.GetTagRequest{
			Tag: s,
		})

		if e != nil {
			panic(e)
		}

		targetBuilder.Tag(s, r.Target.Hash)
	})
	envOptional("BITBUCKET_PATTERN", func(s string) { targetBuilder.Pattern(s) })

	request := model.PostPipelineRequest{
		Pipeline: builders.Pipeline().Target(targetBuilder.Build()).Build(),
	}

	pipeline, err := client.RunPipeline(request)

	assert.Nil(t, err, "failed to run pipeline")
	assert.NotNil(t, pipeline, "pipeline should not be nil")
	assert.NotNil(t, pipeline.UUID, "UUID expected")
}
