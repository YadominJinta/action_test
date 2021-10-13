package main

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Yadomin"
	want := regexp.MustCompile("Hello, " + name)
	msg := Hello(name)
	if !want.MatchString(msg) {
		t.Fatal("Hello(name) test failed")
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""
	msg := Hello(name)
	if msg != "" {
		t.Fatal("Hello(Empty) test failed, should return empty string")
	}
}
