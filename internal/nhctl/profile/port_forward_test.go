package profile

import (
	"errors"
	"fmt"
	"nocalhost/pkg/nhctl/log"
	"testing"
)

func TestParse(t *testing.T) {
	_, _, err := GetPortForwardForString("z")
	if err == nil {
		t.Error(err)
	} else {
		log.Info(err.Error())
	}
}

func TestParseOverload(t *testing.T) {
	l, r, err := GetPortForwardForString("65536")
	if err != nil {
		log.Info(err.Error())
	} else {
		t.Error(errors.New(fmt.Sprintf("err: %v:%v", l, r)))
	}
}

func TestParseOverload1(t *testing.T) {
	l, r, err := GetPortForwardForString("-65535")
	if err != nil {
		log.Info(err.Error())
	} else {
		t.Error(errors.New(fmt.Sprintf("err: %v:%v", l, r)))
	}
}

func TestParseSingle(t *testing.T) {
	l, r, err := GetPortForwardForString("8080")
	if err != nil {
		log.Info(err.Error())
		t.Error(err)
	}

	if l != 8080 || r != 8080 {
		t.Error(errors.New("err"))
	}
}

func TestParseComplete(t *testing.T) {
	l, r, err := GetPortForwardForString("8080:80")
	if err != nil {
		log.Info(err.Error())
		t.Error(err)
	}

	if l != 8080 || r != 80 {
		t.Error(errors.New("err"))
	}
}

func TestParseRandom(t *testing.T) {
	l, r, err := GetPortForwardForString(":80")
	if err != nil {
		log.Info(err.Error())
		t.Error(err)
	}

	if l < 0 || r != 80 {
		t.Error(errors.New("err"))
	}
}
