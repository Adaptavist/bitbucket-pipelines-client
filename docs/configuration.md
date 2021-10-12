# Configuration

To use this CLI, you will need an App password with `pipelines:read` and `repository:admin` scopes.

Having `repository:admin` feels excessive but, this scope is required to POST to the Pipelines endpoint. I would have
hoped Atlassian used `pipelines:write` instead but unfortunately, not.

## Config initialization

```go
client := Client{
    Config: &Config{
        Username: "username",
        Password: "password",
    }
}

// from here you can run operations against the API
pipeline, err := client.GetPipeline(model.GetPipelineRequest{
	Workspace: *string,
	Repository: *string,
	Pipeline: {
        UUID: *string
    }
})
```
