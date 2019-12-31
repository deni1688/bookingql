package services

import (
	"encoding/json"
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

func (f *FleetsterAPI) Get(endpoint string, model interface{}) error {
	base := os.Getenv("SERVER")

	req, err := http.NewRequest("GET", base+endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Add("authorization", f.Token)
	dumpRequestInfo(req)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode > 399 {
		errMsg := fmt.Sprintf("Failed to fetch endpoint %s with status code: %d", endpoint, resp.StatusCode)
		return errors.New(errMsg)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &model)
}

func (f *FleetsterAPI) GetByIds(entity string, ids []string, model interface{}) error {
	query := buildQuery(ids)
	return f.Get(entity+query, &model)
}

func buildQuery(ids []string) string {
	query := "?"

	for i, k := range ids {
		if i > 0 && i <= len(ids)-1 {
			query += "&"
		}
		query += fmt.Sprintf("_id[$in][%d]=%s", i, k)
	}
	return query
}

func dumpRequestInfo(req *http.Request) {
	if b, _ := strconv.ParseBool(os.Getenv("DUMPREQ")); b {
		dumpReq, _ := httputil.DumpRequest(req, false)
		fmt.Println(string(dumpReq))
	}
}
