package core

import "errors"

var (
	CONNECTION_ERR = errors.New("can't connect to api")
	REQUEST_ERR    = errors.New("request error: ")
	DECODE_ERR     = errors.New("decode data error: ")
	AUTH_ERR       = errors.New("Auth error: ")
)
