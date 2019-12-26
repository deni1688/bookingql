package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

var client = &http.Client{}

type FleetsterAPI struct {
	Token string
}

func (f *FleetsterAPI) Get(endpoint string) ([]byte, error) {
	req, err := http.NewRequest("GET", "https://release.fleetster.de"+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("authorization", f.Token)

	requestDump, err := httputil.DumpRequest(req, false)
	fmt.Println(string(requestDump))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 399 {
		errMsg := fmt.Sprintf("Failed to fetch endpoint %s with status code: %d", endpoint, resp.StatusCode)
		return nil, errors.New(errMsg)
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
