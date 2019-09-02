package productdata

import (
	"github.com/hashicorp/go-retryablehttp"
	"github.com/markdicksonjr/impact-radius/util"
	"net/http"
)

type Client struct {
	SID       string
	AuthToken string
}

type GetCatalogItemsParams struct {
	Id      string
	PageUri *string
}

func (i *Client) GetCatalogs() (*CatalogsResponse, error) {
	client := util.GetDefaultClient()

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

	catalogsResponse := CatalogsResponse{}
	if h, err := util.HandleResponse(res, &catalogsResponse); err != nil {
		return nil, err
	} else {
		catalogsResponse.RateLimitInfo.FromHeader(h)
	}

	return &catalogsResponse, nil
}

func (i *Client) GetCatalogItems(params GetCatalogItemsParams) (*CatalogItemsResponse, error) {
	client := util.GetDefaultClient()

	var url string
	if params.PageUri != nil {
		url = "https://products.api.impactradius.com" + *params.PageUri
	} else {
		url = "https://products.api.impactradius.com/Mediapartners/" + i.SID + "/Catalogs/" + params.Id + "/Items"
	}

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

	catalogItemsResponse := CatalogItemsResponse{}
	if h, err := util.HandleResponse(res, &catalogItemsResponse); err != nil {
		return nil, err
	} else {
		catalogItemsResponse.RateLimitInfo.FromHeader(h)
	}

	return &catalogItemsResponse, nil
}

func (r *RateLimitInfo) FromHeader(h http.Header) {
	r.ConcurrentLimit = h.Get("X-Ratelimit-Concurrentlimit")
	r.ConcurrentRemaining = h.Get("X-Ratelimit-Concurrentremaining")
	r.Limit = h.Get("X-Ratelimit-Limit")
	r.Remaining = h.Get("X-Ratelimit-Remaining")
}