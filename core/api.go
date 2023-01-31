package core

import (
	"fmt"
)

const (
	API           string = "https://api.dnslab.ir/"
	IP            string = API + "IP/"
	PING          string = API + "IP/Ping?hostOrIPAddress=%v"
	OPEN          string = API + "IP/IsIPAndPortOpen?hostOrIPAddress=%v&port=%v"
	DNSLOOKUP     string = API + "DNSLookUp/Query?Query=%v&queryType=%v"
	REVERSELookup string = API + "DNSLookUp/QueryReverse?IPAddress=%v"
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

func (c *BaseConf) dnsLookup(q string) (data qTypes, err error) {
	url := fmt.Sprintf(DNSLOOKUP, c.Ip, q)
	res, err := request(url)
	if err != nil {
		return
	}

	switch q {
	case "1":
		return data, decodeBodyJson(res, &data.Q1)
	case "2":
		return data, decodeBodyJson(res, &data.Q2)
	case "5":
		return data, decodeBodyJson(res, &data.Q5)
	case "6":
		return data, decodeBodyJson(res, &data.Q6)
	case "16":
		return data, decodeBodyJson(res, &data.Q16)
	default:
		return data, fmt.Errorf("not found query")
	}
}

func (c *BaseConf) reverseLookup() (string, error) {
	url := fmt.Sprintf(REVERSELookup, c.Ip)
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
