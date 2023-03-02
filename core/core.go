package core

import (
	"fmt"
	"log"
	"net/http"
)

/*
* basic config for request to api
 */
type BaseConf struct {
	Token string
	Ip    string
	Ipv6  string
	Port  uint16
}

/*
* check system ip with request to api.dnslab.ir
* and set ip4 and ip6 and port for config
 */
func (c *BaseConf) CheckIP() error {
	res, err := request(IP)
	if err != nil {
		return err
	}

	var data ipInfo
	if err := decodeBodyJson(res, &data); err != nil {
		return err
	}

	c.SetUpIp(data.IPv4, data.IPv6, 0)
	return nil
}

/*
* ping your ip for check get response or not
 */
func (c *BaseConf) Ping() (*ipOrHostPing, error) {
	url := fmt.Sprintf(PING, c.Ip)
	res, err := request(url)
	if err != nil {
		return nil, err
	}

	var data ipOrHostPing
	if err := decodeBodyJson(res, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

/*
* check one port form on ip is open or not
 */
func (c *BaseConf) OpenPort() (bool, error) {
	url := fmt.Sprintf(OPEN, c.Ip, c.Port)
	res, err := request(url)
	if err != nil {
		res.Body.Close()
		return false, err
	}
	return boolPars(res)
}

/*
* get response with out pars data with
* response models type query!
 */
func (c *BaseConf) DnsLookup(q string) (*http.Response, error) {
	url := fmt.Sprintf(DNSLOOKUP, c.Ip, q)
	return request(url)
}

/*
* Reverse lookup ip/host and return string
 */
func (c *BaseConf) ReverseLookup() (string, error) {
	url := fmt.Sprintf(REVERSELookup, c.Ip)
	res, err := request(url)
	if err != nil {
		return "", err
	}
	return bodyToString(res)
}

/*
* update ip and check response is ok or not
* this needed valid token
* and run with timing
* for example every 1 hour run this
 */
func (c *BaseConf) UpdateIp() (bool, error) {
	url := fmt.Sprintf(DNSREF, c.Token)
	res, err := request(url)
	if err != nil {
		return false, err
	} else if res.StatusCode != 200 {
		return false, fmt.Errorf("token not found")
	}
	return true, nil
}

/*
* setup new ip/port for config
* if IPv4 is empty call method CheckIP()
* if port equal 0 set port 80
 */
func (c *BaseConf) SetUpIp(IPv4, IPv6 string, port uint16) {
	if IPv4 == "" {
		if err := c.CheckIP(); err != nil {
			log.Panic(err)
		}
		return
	}

	c.Ip = IPv4
	c.Ipv6 = IPv6

	c.Port = port
	if port == 0 {
		c.Port = 80
	}
}
