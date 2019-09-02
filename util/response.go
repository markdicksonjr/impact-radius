package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// HandleResponse will interpret the response body and marshal it into each reference provided (if any).  This is
// especially handy when a response could return either data OR an error (and perhaps even a 2XX status code in each case).
func HandleResponse(res *http.Response, objs ...interface{}) (http.Header, error) {
	if res.Body != nil {
		defer res.Body.Close()

		responseData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return res.Header, err
		}

		// if there are no instances to unmarshal into, return now
		if len(objs) == 0 {
			if res.StatusCode != 200 {
				return res.Header, errors.New("failed impactradius request, status code = " + strconv.Itoa(res.StatusCode))
			}

			return res.Header, nil
		}

		// unmarshal into each inbound instance
		for i := range objs {
			if err := json.Unmarshal(responseData, objs[i]); err != nil {
				return res.Header, err
			}
		}
	}

	// if not a 200 status code, return an error (and implicitly, any structure in the response body, if objs provided)
	if res.StatusCode != 200 {
		return res.Header, errors.New("failed impactradius request, status code = " + strconv.Itoa(res.StatusCode))
	}

	return res.Header, nil
}
