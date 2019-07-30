package util

import (
	"crypto/tls"
	"github.com/hashicorp/go-retryablehttp"
	"net/http"
)

func GetDefaultClient() *retryablehttp.Client {

	// request https, allowing for insecure requests
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := retryablehttp.NewClient()
	client.HTTPClient.Transport = tr
	return client
}
