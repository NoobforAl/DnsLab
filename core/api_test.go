package core

import (
	"os"
	"testing"
)

var bc BaseConf

func TestMain(m *testing.M) {
	bc = BaseConf{
		Token: "", // your token
		Ip:    "8.8.8.8",
		Ipv6:  "",
		Port:  443,
	}
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestPing(t *testing.T) {
	data, err := bc.ping()
	if err != nil {
		t.Error(err)
	}

	if !data.Success {
		t.Error("Can't Ping 8.8.8.8!")
	}
}

func TestOpenPort(t *testing.T) {
	isOpen, err := bc.openPort()
	if err != nil {
		t.Error(err)
	}

	if !isOpen {
		t.Error("8.8.8.8 port 80 not open!")
	}
}

func TestDnsLookup(t *testing.T) {
	bc.Ip = "google.com"
	_, err := bc.dnsLookup("1")

	if err != nil {
		t.Error(err)
	}
}

func TestReverseLookup(t *testing.T) {
	msg, err := bc.reverseLookup()
	if err != nil {
		t.Error(err)
	}

	if msg == "" {
		t.Error("nothing find!")
	}
}

func TestUpdateIp(t *testing.T) {
	isOK, err := bc.updateIp()
	if err != nil {
		t.Error(err)
	}

	if isOK {
		t.Error("ip not updated check token!")
	}
}
