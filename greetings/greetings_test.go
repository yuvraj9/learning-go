package greetings

import (
	"testing"
	"regexp"
)

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if err == nil || message != "" {
		t.Fatalf(`Hello("") should have returned error. Message was %q and error returned = %v`, message, err)
	}
}

func TestHello(t *testing.T) {
	name := "yuvraj"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("yuvraj")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("yuvraj") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}