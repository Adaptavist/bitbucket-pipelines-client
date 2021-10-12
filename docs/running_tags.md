# Running tags

Running tags requires you to know a couple of things up front. 1, the tag. 2, the commit hash. Without these, you're not
going to get very far. Ideally, the end-users shouldn't care about the commit hash of the tag, so we should look it up
on their behalf.

## Getting tags

```bash
curl \
    --location \
    --header "Authorization: Basic ${BASE64_AUTH_STR}" \
    --request GET "https://api.bitbucket.org/2.0/repositories/${SPACE}/${REPO}/refs/tags?q=name=%22${TAG}%22"
```

We want the commit hash of the tag, which "should" be in `values[0].target.hash` but also ensure that
`values[0].name` is the tag you're looking for. Use the example below to get an idea of what we're looking for

```json
// truncated output to show only the fields we care about.
{
    "pagelen": 10,
    "values": [
        {
            "name": "v0.1.0",
            "target": {
                "hash": "e28c6cad5a664eb1a5fcd46c3bb256dd426ef721"
            }
        }
    ],
    "page": 1
}
```

## Running the tags pipeline

Let's run the pipeline now we've got the tag and the commit hash.

```bash
curl \
    --location \
    --header "Authorization: Basic ${BASE64_AUTH_STR}" \
    --header 'Content-Type: application/json' \
    --request POST "https://api.bitbucket.org/2.0/repositories/${SPACE}/${REPO}/pipelines/" \
    --data-raw '{
    "target": {
        "ref_name": "v0.1.0",
        "ref_type": "tag",
        "type": "pipeline_ref_target",
        "commit": { "type": "commit", "hash": "e28c6cad5a664eb1a5fcd46c3bb256dd426ef721" }
    }
}'
```

Without that commit hash, the Bitbucket will not be able to find the pipeline.