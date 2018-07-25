package codeclimate

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func MakeRequest(client *Client, method string, path string) (string, error) {
	httpClient := &http.Client{}
	postData := make([]byte, 100)
	targetUrl := client.ApiBasePath + path
	req, err := http.NewRequest(strings.ToUpper(method), targetUrl, bytes.NewReader(postData))

	if err != nil {
		return "", err
	}

	authHeader := "Token token=" + client.ApiToken
	req.Header.Add("Accept", "application/vnd.api+json")
	req.Header.Add("Authorization", authHeader)
	resp, err := httpClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	bodyString := string(bodyBytes)
	return bodyString, nil
}
