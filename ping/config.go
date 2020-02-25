package ping

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/chrisUsick/pingidentityapi"
	api "github.com/chrisUsick/pingidentityapi"
)

type Config struct {
	Username           string
	Password           string
	BaseURL            string
	XSRFHeader         string
	InsecureSkipVerify bool
}

func (c Config) Client(client pingidentityapi.IClient) interface{} {
	if client == nil {
		var tr http.RoundTripper = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}
		tr.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: c.InsecureSkipVerify}
		client := api.NewClient(&api.Configuration{c.BaseURL, c.Username, c.Password, c.XSRFHeader, tr})
		return client
	}
	return client
}
