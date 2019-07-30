package mediapartners

import (
	"github.com/hashicorp/go-retryablehttp"
	"github.com/markdicksonjr/impact-radius/util"
)

type Client struct {
	SID       string
	AuthToken string
}

func (i *Client) GetCompanyInformation() (*CompanyInfoResponse, error) {
	client := util.GetDefaultClient()

	url := "https://products.api.impactradius.com/Mediapartners/" + i.SID + "/CompanyInformation"
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

	companyInfoResponse := CompanyInfoResponse{}
	if err := util.HandleResponse(res, &companyInfoResponse); err != nil {
		return nil, err
	}

	return &companyInfoResponse, nil
}
