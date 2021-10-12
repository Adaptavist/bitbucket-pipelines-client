package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	retry "github.com/hashicorp/go-retryablehttp"
)

func (h *Client) getHttpClient() *retry.Client {
	if h.client == nil {
		h.client = retry.NewClient()
	}
	return h.client
}

func (h *Client) doRequest(r *retry.Request) (*http.Response, error) {
	r.SetBasicAuth(h.Config.Username, h.Config.Password)
	r.Header.Add("Content-Type", "application/json")
	return h.getHttpClient().Do(r)
}

func (h Client) get(url string) (resp []byte, err error) {
	req, err := retry.NewRequest("GET", url, nil)

	if err != nil {
		return
	}

	res, err := h.doRequest(req)

	if err != nil {
		return
	}

	err = hasError(res.StatusCode)

	if err != nil {
		err = fmt.Errorf("failed to GET (%s) %s", url, err)
		return
	}

	resp, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return
	}

	return
}

// hasError returns an error if 40x or 50x codes are given
func hasError(s int) (err error) {
	if s >= 400 && s < 600 {
		err = fmt.Errorf("received %s", strconv.Itoa(s))
	}
	return
}

// post an BasicAuth authenticated resource
func (h Client) post(url string, data interface{}) (resp []byte, err error) {
	reqData, err := json.Marshal(data)

	if err != nil {
		return
	}

	reqBody := bytes.NewBuffer(reqData)
	req, err := retry.NewRequest("POST", url, reqBody)

	if err != nil {
		return
	}

	res, err := h.doRequest(req)

	if err != nil {
		return
	}

	httpError := hasError(res.StatusCode)

	resp, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return
	}

	if httpError != nil {
		err = fmt.Errorf("failed to POST (%s) %s - %s", url, httpError, string(resp))
		return
	}

	return
}

// postUnmarshalled makes a POST HTTP request and unmarshalls the data
func (h Client) postUnmarshalled(url string, data interface{}, target interface{}) (err error) {
	resp, err := h.post(url, data)

	if err != nil {
		return
	}

	err = json.Unmarshal(resp, target)

	return
}

// getUnmarshalled makes a GET HTTP request and unmarshalls the data
func (h Client) getUnmarshalled(url string, targetPtr interface{}) (err error) {
	resp, err := h.get(url)

	if err != nil {
		err = fmt.Errorf("%s - %s", err.Error(), string(resp))
		return
	}

	err = json.Unmarshal(resp, targetPtr)

	return
}
