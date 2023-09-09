package core

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/NoobforAl/DnsLab/src/contract"
	"github.com/NoobforAl/DnsLab/src/core"
	"github.com/kr/pretty"
	"github.com/sirupsen/logrus"
)

const (
	example1 = "example.com"
	example2 = "8.8.8.8"
)

var bc contract.DnslabAPi

func TestMain(m *testing.M) {
	// setup token
	// export token_app="token"

	bc = core.New(
		os.Getenv("token_app"),
		logrus.New(),
		time.Second, 2)
	os.Exit(m.Run())
}

func TestCheckIp(t *testing.T) {
	data, err := bc.CheckIP()
	if err != nil {
		t.Error(err)
	}

	t.Logf("%# v", pretty.Formatter(data))
}

func TestPing(t *testing.T) {
	data, err := bc.Ping(example1)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%# v", pretty.Formatter(data))
}

func TestOpenPort(t *testing.T) {
	isOpen, err := bc.OpenPort(example2, 53)
	if err != nil {
		t.Error(err)
	}

	if !isOpen {
		t.Error("8.8.8.8 port 53 is open!")
	}
}

func TestDnsLookup(t *testing.T) {
	querys := []string{"1", "2", "5", "6", "16"}
	for _, v := range querys {
		data, err := bc.DnsLookup(example1, v)
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("%# v", pretty.Formatter(data))
	}
}

func TestReverseLookup(t *testing.T) {
	msg, err := bc.ReverseLookup(example2)
	if err != nil {
		t.Error(err)
	}

	if msg == "" {
		t.Error("nothing find!")
	}
}

func TestUpdateIp(t *testing.T) {
	isOK, err := bc.UpdateIp()
	if err != nil {
		t.Error(err)
	}

	if !isOK {
		t.Error("ip not updated check token!")
	}
}
