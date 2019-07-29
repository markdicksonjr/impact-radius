package impact_radius

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/hashicorp/go-retryablehttp"
	"io/ioutil"
	"net/http"
)

type Client struct {
	SID       string
	AuthToken string
}

func (i *Client) GetCatalogs() (*CatalogsResponse, error) {

	// request https, allowing for insecure requests
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := retryablehttp.NewClient()
	client.HTTPClient.Transport = tr

	url := "https://products.api.impactradius.com/Mediapartners/" + i.SID + "/Catalogs"
	req, err := retryablehttp.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(i.SID, i.AuthToken)
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New("failed impactradius request")
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	catalogsResponse := CatalogsResponse{}
	if err := json.Unmarshal(responseData, &catalogsResponse); err != nil {
		return nil, err
	}

	return &catalogsResponse, nil
}

func (i *Client) GetCatalogItems(id string) (*CatalogItemsResponse, error) {

	// request https, allowing for insecure requests
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := retryablehttp.NewClient()
	client.HTTPClient.Transport = tr

	url := "https://products.api.impactradius.com/Mediapartners/" + i.SID + "/Catalogs/" + id + "/Items"
	req, err := retryablehttp.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(i.SID, i.AuthToken)
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New("failed impactradius request")
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	catalogItemsResponse := CatalogItemsResponse{}
	if err := json.Unmarshal(responseData, &catalogItemsResponse); err != nil {
		return nil, err
	}

	return &catalogItemsResponse, nil
}
