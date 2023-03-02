package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// * read Response Body and convert to string
func bodyToString(res *http.Response) (string, error) {
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("err :  %v", err)
	}
	return string(b), nil
}

/*
* check response body is bool or not
* first parse with bodyToString()
* secund check is equal to true or not
* if value not "true" return false
* this func only use for some request then not response json
* ex: openPort() method use this func
 */
func boolPars(res *http.Response) (bool, error) {
	b, err := bodyToString(res)

	defer res.Body.Close()

	if b == "true" {
		return true, err
	}
	return false, err
}

/*
* decode response json with some struct
* list all struct in core/models.go
 */
func decodeBodyJson[T response](res *http.Response, data *T) error {
	de := json.NewDecoder(res.Body)

	defer res.Body.Close()

	if err := de.Decode(&data); err != nil {
		return fmt.Errorf("err : %v", err)
	}
	return nil
}

/*
* request get url and request to url
* if not get error or not get bad status code ( only 200 is ok! )
* return response
 */
func request(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("can't connect to api")
	}

	if res.StatusCode != 200 {
		text, _ := bodyToString(res)
		res.Body.Close()
		return nil, fmt.Errorf("err : server response with this StatusCode %d\nresponse value:%s",
			res.StatusCode, text)
	}
	return res, nil
}
