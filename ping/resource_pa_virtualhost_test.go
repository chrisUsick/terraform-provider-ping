package ping

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/chrisUsick/pingidentityapi"
	resty "gopkg.in/resty.v1"

	"github.com/chrisUsick/terraform-provider-ping/ping/mocks"
	"github.com/stretchr/testify/mock"

	r "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func response404() *pingidentityapi.ClientError {
	return &pingidentityapi.ClientError{
		FullResponse: &resty.Response{
			RawResponse: &http.Response{
				StatusCode: 404,
			},
		},
		Body: map[string]interface{}{},
	}
}
func TestCreateVirtualhost(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		IsUnitTest: true,
		Providers: CreateProvider(func(mockClient *mocks.IClient) {
			mockClient.On("Get", "virtualhosts/1").Return(map[string]interface{}{
				"id": 1,
			}, nil)
			mockClient.On("Post", mock.AnythingOfType("string"), mock.Anything).
				Return(map[string]interface{}{
					"id": 1.0,
				}, nil)
			mockClient.On("Delete", mock.Anything).
				Return(map[string]interface{}{
					"msg": "Operation successful.",
				}, nil)
		}),
		Steps: []r.TestStep{
			r.TestStep{
				Config: testCreateVirtualhostConfig(),
				Check: func(s *terraform.State) error {
					instanceState := s.RootModule().Resources["ping_virtualhost.test"].Primary
					id := instanceState.ID
					if id != "1" {
						return fmt.Errorf("Didn't set correct ID")
					}

					if instanceState.Attributes["port"] != "3000" {
						return fmt.Errorf("Incorrect Port")
					}
					if host := instanceState.Attributes["host"]; host != "test" {
						return fmt.Errorf("Inccorect host `%s`", host)
					}
					return nil
				},
			},
		},
	})
}

func TestUpdateVirtualhost(t *testing.T) {
	r.UnitTest(t, r.TestCase{
		IsUnitTest: true,
		Providers: CreateProvider(func(mockClient *mocks.IClient) {
			mockClient.On("Get", "virtualhosts/1").Return(map[string]interface{}{
				"id":   1.0,
				"host": "test",
				"port": 3000.0,
			}, nil)
			mockClient.On("Post", "virtualhosts", mock.MatchedBy(func(body map[string]interface{}) bool {
				return body["port"] == 3000
			})).
				Return(map[string]interface{}{
					"id": 1.0,
				}, nil)
			mockClient.On("Put", "virtualhosts/1", mock.MatchedBy(func(body map[string]interface{}) bool {
				return body["port"] == 4000
			})).
				Return(map[string]interface{}{
					"id": 1.0,
				}, nil)
			mockClient.On("Delete", mock.Anything).
				Return(map[string]interface{}{
					"msg": "Operation successful.",
				}, nil)
		}),
		Steps: []r.TestStep{
			r.TestStep{
				Config:  testCreateVirtualhostConfig(),
				Destroy: false,
			},
			r.TestStep{
				Config:  testUpdateVirtualhostConfig(),
				Destroy: false,
				Check: func(s *terraform.State) error {
					instanceState := s.RootModule().Resources["ping_virtualhost.test"].Primary
					if instanceState.ID != "1" {
						return fmt.Errorf("Wrong ID: `%s`", instanceState.ID)
					}
					if port := instanceState.Attributes["port"]; port != "4000" {
						return fmt.Errorf("Wrong port: `%v`", port)
					}
					return nil
				},
			},
		},
	})
}

func testCreateVirtualhostConfig() string {
	return fmt.Sprintf(`
		provider "ping" {
			username             = "Administrator"
			password             = "Testpassword1"
			base_url             = "https://192.168.33.111:9000/pa-admin-api/v3/"
			xsrf_header			 = "PingAccess"
			insecure_skip_verify = true
		}
		resource "ping_virtualhost" "test" {
			host = "test"
			port = 3000
		}
		`)
}

func testUpdateVirtualhostConfig() string {
	return fmt.Sprintf(`
		provider "ping" {
			username             = "Administrator"
			password             = "Testpassword1"
			base_url             = "https://192.168.33.111:9000/pa-admin-api/v3/"
			xsrf_header			 = "PingAccess"
			insecure_skip_verify = true
		}
		resource "ping_virtualhost" "test" {
			host = "test"
			port = 4000
		}
		`)
}
