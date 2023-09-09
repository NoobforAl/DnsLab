package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	retry "github.com/avast/retry-go"
)

// * read Response Body and convert to string
func (c BaseConf) bodyToString(res *http.Response) (string, error) {
	c.log.Debug("response to string")

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.Join(DECODE_ERR, err)
	}

	c.log.Debug(b)
	return string(b), nil
}

/*
* decode response json with some struct
* list all struct in core/models.go
 */
func (c BaseConf) decodeBodyJson(
	res *http.Response,
	data *map[string]any,
	datas *[]map[string]any,
) error {
	c.log.Debug("decode json body to map")

	de := json.NewDecoder(res.Body)
	var err error

	defer res.Body.Close()

	if data == nil {
		err = de.Decode(&datas)
	} else {
		err = de.Decode(&data)
	}

	if err != nil {
		text, _ := c.bodyToString(res)
		err = fmt.Errorf("%s, body: %s", err.Error(), text)
		return errors.Join(DECODE_ERR, err)
	}
	return nil
}

/*
* request get url and request to url
* if not get error or not get bad status code ( only 200 is ok! )
* return response
 */
func (c BaseConf) request(url string) (*http.Response, error) {
	var res *http.Response
	var err error

	c.log.Debugf("request to url: %s", url)

	err = retry.Do(func() error {
		res, err = http.Get(url)
		if err != nil {
			return CONNECTION_ERR
		}

		if res.StatusCode != 200 {
			text, _ := c.bodyToString(res)
			defer res.Body.Close()

			err = fmt.Errorf("StatusCode %d\nbody: %s", res.StatusCode, text)
			return errors.Join(REQUEST_ERR, err)
		}
		return nil
	}, retry.Delay(c.retryTime), retry.Attempts(c.retryCount))

	return res, err
}
