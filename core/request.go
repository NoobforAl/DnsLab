package core

import (
	"fmt"
	"net/http"
)

func request(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("can't connect to api")
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("err : server response with this StatusCode %v", res.StatusCode)
	}
	return res, nil
}
