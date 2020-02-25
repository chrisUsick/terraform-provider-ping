package pingidentityapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/resty.v1"
)

type Client struct {
	baseURL  string
	username string
	password string
	*resty.Client
}

type ClientError struct {
	FullResponse *resty.Response
	Body         map[string]interface{}
}

func (e *ClientError) Error() string {
	return string(e.FullResponse.Body())
}

type IClient interface {
	Get(path string) (map[string]interface{}, error)
	Post(path string, body map[string]interface{}) (map[string]interface{}, error)
	Put(path string, body map[string]interface{}) (map[string]interface{}, error)
	Delete(path string) (map[string]interface{}, error)
}

type Configuration struct {
	BaseURL   string
	Username  string
	Password  string
	XSRFHeader string
	Transport http.RoundTripper
}

func NewClient(config *Configuration) *Client {
	client := resty.New()
	if config.Transport != nil {
		client.SetTransport(config.Transport)
	}
	if config.XSRFHeader != "" {
		client.SetHeader("X-Xsrf-Header", config.XSRFHeader)
	} else {
		client.SetHeader("X-Xsrf-Header", "PingAccess")
	}
	client.SetHeader("Accept", "application/json")
	client.SetHeader("Content-Type", "application/json")
	client.SetBasicAuth(config.Username, config.Password)
	client.SetRESTMode()
	return &Client{
		baseURL: config.BaseURL,
		Client:  client,
	}
}

type errorJson struct {
	json map[string]interface{}
}

func (e *errorJson) Error() string {
	s, err := json.Marshal(e.json)
	if s == nil {
		return fmt.Sprintf("Error marshalling json: %v", err)
	}
	return string(s)
}

func (c *Client) Get(path string) (map[string]interface{}, error) {
	resp, err := c.R().Get(c.baseURL + path)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	json.Unmarshal(resp.Body(), &m)
	if resp.StatusCode() != 200 {
		return nil, &ClientError{Body: m, FullResponse: resp}
	}
	return m, err
}

func (c *Client) Post(path string, body map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.R().SetBody(body).Post(c.baseURL + path)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	json.Unmarshal(resp.Body(), &m)
	if resp.StatusCode() != 200 {
		return nil, &errorJson{m}
	}
	return m, err
}

func (c *Client) Put(path string, body map[string]interface{}) (map[string]interface{}, error) {
	resp, err := c.R().SetBody(body).Put(c.baseURL + path)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	json.Unmarshal(resp.Body(), &m)
	if resp.StatusCode() != 200 {
		return nil, &ClientError{Body: m, FullResponse: resp}
	}
	return m, err
}

func (c *Client) Delete(path string) (map[string]interface{}, error) {
	resp, err := c.R().Delete(c.baseURL + path)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	json.Unmarshal(resp.Body(), &m)
	if resp.StatusCode() != 200 {
		return nil, &errorJson{m}
	}
	return m, err
}
