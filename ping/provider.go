package ping

import (
	"github.com/chrisUsick/pingidentityapi"
	"github.com/hashicorp/terraform/helper/schema"
)

func ProviderFactory(client pingidentityapi.IClient) *schema.Provider {
	if client == nil {
		return provider()
	}
	p := provider()
	p.ConfigureFunc = providerConfigureFactory(client)
	return p
}

func provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"USERNAME",
				}, nil),
				Description: "Username",
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"PASSWORD",
				}, nil),
				Description: "Password",
			},
			"base_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"BASE_URL",
				}, nil),
			},
			"xsrf_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"XSRF_HEADER",
				}, nil),
			},
			"insecure_skip_verify": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"INSECURE_SKIP_VERIFY",
				}, false),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ping_server":      resourceServer(),
			"ping_virtualhost": resourcePAVirtualhost(),
		},
	}
}

func providerConfigureFactory(client pingidentityapi.IClient) func(*schema.ResourceData) (interface{}, error) {
	if client == nil {
		return func(d *schema.ResourceData) (interface{}, error) {
			return providerConfigure(nil, d)
		}
	}
	return func(d *schema.ResourceData) (interface{}, error) {
		return providerConfigure(client, d)
	}
}
func providerConfigure(client pingidentityapi.IClient, d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Username:           d.Get("username").(string),
		Password:           d.Get("password").(string),
		BaseURL:            d.Get("base_url").(string),
		XSRFHeader:          d.Get("xsrf_header").(string),
		InsecureSkipVerify: d.Get("insecure_skip_verify").(bool),
	}
	return config.Client(client), nil
}
