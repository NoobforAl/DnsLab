package core

import (
	"testing"

	"github.com/NoobforAl/DnsLab/core"
)

/*
* setup basic config for testCase
* creat one token form `dnslab` for test
 */
var bc = core.BaseConf[core.Config, []core.DnsLookUpQT1]{
	Token: "",
	Ip:    "8.8.8.8",
	Ipv6:  "",
	Port:  443,
}

func TestPing(t *testing.T) {
	data, err := bc.Ping()
	if err != nil {
		t.Error(err)
	}

	if !data.Success {
		t.Error("Can't Ping 8.8.8.8!")
	}
}

func TestOpenPort(t *testing.T) {
	isOpen, err := bc.OpenPort()
	if err != nil {
		t.Error(err)
	}

	if !isOpen {
		t.Error("8.8.8.8 port 80 not open!")
	}
}

// ! this test case only test query type 1
func TestDnsLookup(t *testing.T) {
	bc.Ip = "google.com"
	_, err := bc.DnsLookup("1")

	if err != nil {
		t.Error(err)
	}
}

func TestReverseLookup(t *testing.T) {
	msg, err := bc.ReverseLookup()
	if err != nil {
		t.Error(err)
	}

	if msg == "" {
		t.Error("nothing find!")
	}
}

// ! this test case needed token
func TestUpdateIp(t *testing.T) {
	isOK, err := bc.UpdateIp()
	if err != nil {
		t.Error(err)
	}

	if isOK {
		t.Error("ip not updated check token!")
	}
}
