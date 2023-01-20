package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func bodyToString(res *http.Response) (string, error) {
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("err :  %v", err)
	}
	return string(b), nil
}

func boolPars(res *http.Response) (bool, error) {
	b, err := bodyToString(res)
	if string(b) == "true" {
		return true, err
	}
	return false, err
}

func decodeBodyJson[T response](res *http.Response, data *T) error {
	de := json.NewDecoder(res.Body)
	if err := de.Decode(&data); err != nil {
		return fmt.Errorf("err : %v", err)
	}
	return nil
}
