package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	API           string = "https://api.dnslab.ir/"
	IP            string = API + "IP/"
	PING          string = API + "IP/IPHavePing?hostOrIPAddress=%v"
	OPEN          string = API + "IP/IsIPAndPortOpen?hostOrIPAddress=%v&port=%v"
	DNSLOOKUP     string = API + "IP/DNSLookup?hostOrIPAddress=%v"
	REVERSELookup string = API + "IP/ReverseLookup?IPAddress=%v"
	DNSREF        string = API + "DNS/UDIBT?k=%v"

	ENTERIP  string = "if enter your ip checked!\nEnter ip or enter: "
	ENTERDNS string = "Enter dns: "
)

type BaseConf struct {
	Token string
	Ip    string
	Ipv6  string
	Port  uint16
}

type response struct {
	IPv4 string `json:"iPv4"`
	IPv6 string `json:"iPv6"`
}

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

func getIpPort(getPort bool, text string) (string, uint16) {
	var ip string
	var port uint16

	fmt.Print(text)
	fmt.Scanln(&ip)

	if getPort {
		fmt.Print("Enter port or enter:")
		fmt.Scanln(&port)
	}
	return ip, port
}

func decodeBodyJson(res *http.Response) (*response, error) {
	r := &response{}
	de := json.NewDecoder(res.Body)

	if err := de.Decode(&r); err != nil {
		return nil, fmt.Errorf("err : %v", err)
	}

	return r, nil
}

func (c *BaseConf) CheckIP() error {
	res, err := request(IP)
	if err != nil {
		return err
	}

	r, err := decodeBodyJson(res)
	if err != nil {
		return err
	}

	c.Ip = r.IPv4
	c.Ipv6 = r.IPv6
	c.Port = 80
	return err
}

func (c *BaseConf) setUpNewIp(getPort bool, text string) {
	ip, port := getIpPort(getPort, text)
	if ip != "" {
		c.Ip = ip
	}

	if port != 0 {
		c.Port = port
	}
}

func (c *BaseConf) Ping() (bool, error) {
	c.setUpNewIp(false, ENTERIP)
	url := fmt.Sprintf(PING, c.Ip)
	res, err := request(url)
	if err != nil {
		return false, err
	}
	defer c.CheckIP()
	return boolPars(res)
}

func (c *BaseConf) OpenPort() (bool, error) {
	c.setUpNewIp(true, ENTERIP)
	url := fmt.Sprintf(OPEN, c.Ip, c.Port)
	res, err := request(url)
	if err != nil {
		return false, err
	}
	defer c.CheckIP()
	return boolPars(res)
}

func (c *BaseConf) DnsLookup() (string, error) {
	c.setUpNewIp(false, ENTERDNS)
	url := fmt.Sprintf(DNSLOOKUP, c.Ip)
	res, err := request(url)
	if err != nil {
		return "", err
	}
	defer c.CheckIP()
	return bodyToString(res)
}

func (c *BaseConf) ReverseLookup() (string, error) {
	c.setUpNewIp(false, ENTERIP)
	url := fmt.Sprintf(DNSLOOKUP, c.Ip)
	res, err := request(url)
	if err != nil {
		return "", err
	}
	defer c.CheckIP()
	return bodyToString(res)
}

func (c *BaseConf) UpdateIp() (bool, error) {
	url := fmt.Sprintf(DNSREF, c.Token)
	res, err := request(url)
	if err != nil {
		return false, err
	}
	defer c.CheckIP()
	return boolPars(res)
}

func (c *BaseConf) Show(text, v any, err error) {
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(text, v)
	}
}
