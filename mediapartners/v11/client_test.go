package v11

import (
	"os"
	"testing"
)

func TestClient_GetCompanyInformation(t *testing.T) {

	i := Client{
		AuthToken: os.Getenv("IMPACT_AUTHTOKEN"),
		SID:       os.Getenv("IMPACT_SID"),
	}

	r, err := i.GetCompanyInformation()
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Fatal("response was null")
	}
}
