package test

import (
	"github.com/antony0016/go-test/a"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	if a.HelloWorld() != "HelloWorld" {
		t.Errorf("HelloWorld test failed")
	}
}
