package services

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
)

var client = &http.Client{}

type FleetsterAPI struct {
	Token string
}

func (f *FleetsterAPI) Get(endpoint string) ([]byte, error) {
	base := os.Getenv("SERVER")

	req, err := http.NewRequest("GET", base+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("authorization", f.Token)

	dumpRequestInfo(err, req)

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

func dumpRequestInfo(err error, req *http.Request) {
	if b, _ := strconv.ParseBool(os.Getenv("DUMPREQ")); b {
		dumpReq, _ := httputil.DumpRequest(req, false)
		fmt.Println(string(dumpReq))
	}
}
