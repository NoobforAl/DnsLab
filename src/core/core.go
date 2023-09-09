package core

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/NoobforAl/DnsLab/src/contract"
)

/*
* basic config for request to api
 */
type BaseConf struct {
	retryTime  time.Duration
	retryCount uint

	log   contract.Logger
	token string
}

func New(
	token string,
	log contract.Logger,
	retryTime time.Duration,
	retryCount uint,
) *BaseConf {
	return &BaseConf{
		token:      token,
		log:        log,
		retryTime:  retryTime,
		retryCount: retryCount,
	}
}

/*
* check system ip with request to api.dnslab.link
* and set ip4 and ip6 and port for config
 */
func (c *BaseConf) CheckIP() (map[string]any, error) {
	res, err := c.request(IP)
	if err != nil {
		return nil, err
	}

	data := make(map[string]any)
	if err := c.decodeBodyJson(res, &data); err != nil {
		return nil, err
	}

	return data, nil
}

/*
* ping your ip for check get response or not
 */
func (c *BaseConf) Ping(ipOrHost string) (map[string]any, error) {
	url := fmt.Sprintf(PING, ipOrHost)
	res, err := c.request(url)
	if err != nil {
		return nil, err
	}

	data := make(map[string]any)
	if err := c.decodeBodyJson(res, &data); err != nil {
		return nil, err
	}

	return data, nil
}

/*
* check one port form on ip is open or not
 */
func (c BaseConf) OpenPort(ipOrHost string, prot uint16) (bool, error) {
	url := fmt.Sprintf(OPEN, ipOrHost, prot)
	res, err := c.request(url)
	if err != nil {
		return false, err
	}
	s, err := c.bodyToString(res)
	return strings.EqualFold(s, "true"), err
}

/*
* get response with out pars data with
* response models type query!
 */
func (c BaseConf) DnsLookup(ipOrHost, q string) (map[string]any, error) {
	url := fmt.Sprintf(DNS_LOOKUP, ipOrHost, q)
	res, err := c.request(url)
	if err != nil {
		return nil, err
	}

	data := make(map[string]any)
	if err := c.decodeBodyJson(res, &data); err != nil {
		return nil, err
	}

	return data, nil
}

/*
* Reverse lookup ip/host and return string
 */
func (c BaseConf) ReverseLookup(ip string) (string, error) {
	url := fmt.Sprintf(REVERSELookup, ip)
	res, err := c.request(url)
	if err != nil {
		return "", err
	}
	return c.bodyToString(res)
}

/*
* update ip and check response is ok or not
* this needed valid token
* and run with timing
* for example every 1 hour run this
 */
func (c BaseConf) UpdateIp() (bool, error) {
	url := fmt.Sprintf(DNS_REF, c.token)
	res, err := c.request(url)
	if err != nil {
		return false, err
	} else if res.StatusCode != 200 {
		txt, _ := c.bodyToString(res)
		err = fmt.Errorf("body-> %s", txt)
		return false, errors.Join(AUTH_ERR, err)
	}
	return true, nil
}
