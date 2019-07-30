package productdata

import (
	"os"
	"testing"
)

func TestClient_GetCatalogs(t *testing.T) {

	i := Client{
		AuthToken: os.Getenv("IMPACT_AUTHTOKEN"),
		SID:       os.Getenv("IMPACT_SID"),
	}

	r, err := i.GetCatalogs()
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Fatal("response was null")
	}
}

func TestClient_GetCatalogItems(t *testing.T) {

	i := Client{
		AuthToken: os.Getenv("IMPACT_AUTHTOKEN"),
		SID:       os.Getenv("IMPACT_SID"),
	}

	r, err := i.GetCatalogItems("530")
	if err != nil {
		t.Fatal(err)
	}

	if r == nil {
		t.Fatal("response was null")
	}
}