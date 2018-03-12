package ping

import (
	api "github.com/chrisUsick/pingidentityapi"
	"github.com/chrisUsick/terraform-provider-ping/ping/mocks"
	"github.com/hashicorp/terraform/terraform"
)

type FakeClient struct {
}

const (
	GET    HTTPMethod = 1
	POST   HTTPMethod = 2
	PUT    HTTPMethod = 3
	DELETE HTTPMethod = 4
)

type Request struct {
	method      HTTPMethod
	path        string
	body        map[string]interface{}
	resultBody  map[string]interface{}
	resultError error
}

func (c *FakeClient) Get(path string) (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}

func (c *FakeClient) Post(path string, body map[string]interface{}) (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}

func (c *FakeClient) Put(path string, body map[string]interface{}) (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}

func (c *FakeClient) Delete(path string) (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}

func createClient(configureMock func(mock *mocks.IClient)) api.IClient {
	client := &mocks.IClient{}
	configureMock(client)
	return client
}

type HTTPMethod int

func CreateProvider(configureMock func(mock *mocks.IClient)) map[string]terraform.ResourceProvider {
	return map[string]terraform.ResourceProvider{
		"ping": ProviderFactory(createClient(configureMock)),
	}
}
