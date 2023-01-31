package core

import (
	"fmt"
)

const (
	API           string = "https://api.dnslab.ir/"
	IP            string = API + "IP/"
	PING          string = API + "IP/Ping?hostOrIPAddress=%v"
	OPEN          string = API + "IP/IsIPAndPortOpen?hostOrIPAddress=%v&port=%v"
	DNSLOOKUP     string = API + "IP/DNSLookup?hostOrIPAddress=%v"
	REVERSELookup string = API + "IP/ReverseLookup?IPAddress=%v"
	DNSREF        string = API + "DNS/UDIBT?k=%v"

	ENTER_IP  string = "if enter, your ip checked!\nEnter ip or Host: "
	ENTER_DNS string = "Enter dns: "
)

func (c *BaseConf) checkIP() error {
	res, err := request(IP)
	if err != nil {
		return err
	}

	var data ipInfo
	if err := decodeBodyJson(res, &data); err != nil {
		return err
	}

	c.setUpIp(data.IPv4, data.IPv6, 0)
	return nil
}

func (c *BaseConf) ping() (*ipOrHostPing, error) {
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

func (c *BaseConf) openPort() (bool, error) {
	url := fmt.Sprintf(OPEN, c.Ip, c.Port)
	res, err := request(url)
	if err != nil {
		return false, err
	}
	return boolPars(res)
}

func (c *BaseConf) dnsLookup() (string, error) {
	url := fmt.Sprintf(DNSLOOKUP, c.Ip)
	res, err := request(url)
	if err != nil {
		return "", err
	}
	return bodyToString(res)
}

func (c *BaseConf) reverseLookup() (string, error) {
	url := fmt.Sprintf(DNSLOOKUP, c.Ip)
	res, err := request(url)
	if err != nil {
		return "", err
	}
	return bodyToString(res)
}

func (c *BaseConf) updateIp() (bool, error) {
	url := fmt.Sprintf(DNSREF, c.Token)
	res, err := request(url)
	if err != nil {
		return false, err
	}
	return boolPars(res)
}
