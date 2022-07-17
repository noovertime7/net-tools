package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	PostMethod = "POST"
	GETMethod  = "GET"
)

func CurlCheck(url, method, content string) error {
	switch method {
	case PostMethod:
		buf := new(bytes.Buffer)
		encoder := json.NewEncoder(buf)
		err := encoder.Encode(content)
		if err != nil {
			return err
		}
		resp, err := http.Post(url, "application/json", buf)
		if err != nil {
			return err
		}
		bodyRes, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("服务端回应:", string(bodyRes))
	case GETMethod:
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		bodyRes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println("服务端回应:", string(bodyRes))
	}
	return nil
}
