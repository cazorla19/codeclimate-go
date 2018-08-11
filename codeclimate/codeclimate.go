package codeclimate

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

const apiBaseUrl string = "https://api.codeclimate.com/v1"

type Client struct {
	ApiToken    string
	ApiBasePath string
}

type Option func(*Client) error

func (c *Client) parseOptions(opts ...Option) error {
	// Range over each options function and apply it to our API type to
	// configure it. Options functions are applied in order, with any
	// conflicting options overriding earlier calls.
	for _, option := range opts {
		err := option(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewClient(token string, opts ...Option) (*Client, error) {
	client := &Client{ApiToken: token, ApiBasePath: apiBaseUrl}
	if err := client.parseOptions(opts...); err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) MakeRequest(method string, path string, postData []byte) (string, error) {
	httpClient := &http.Client{}
	targetUrl := c.ApiBasePath + path
	reqMethod := strings.ToUpper(method)
	data := bytes.NewReader(postData)
	req, requestErr := http.NewRequest(reqMethod, targetUrl, data)

	if requestErr != nil {
		return "", requestErr
	}

	authHeader := "Token token=" + c.ApiToken
	req.Header.Add("Accept", "application/vnd.api+json")
	req.Header.Add("Authorization", authHeader)

	if reqMethod == "POST" {
		req.Header.Add("Content-Type", "application/vnd.api+json")
	}

	resp, responseErr := httpClient.Do(req)
	defer resp.Body.Close()

	if responseErr != nil {
		return "", responseErr
	}

	bodyBytes, ioErr := ioutil.ReadAll(resp.Body)

	if ioErr != nil {
		return "", ioErr
	}

	bodyString := string(bodyBytes)
	return bodyString, nil
}
