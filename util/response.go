package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

func HandleResponse(res *http.Response, obj interface{}) error {
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New("failed impactradius request, status code = " + strconv.Itoa(res.StatusCode))
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(responseData, obj); err != nil {
		return err
	}

	return nil
}
