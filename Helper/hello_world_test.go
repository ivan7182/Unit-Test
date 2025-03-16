package helper

import (
	"fmt"
	"testing"
)

// /Fail
func TestHelloWorldIvan(t *testing.T) {
	value := HelloWorld("ivan")

	if value != "hello ivan" {
		t.Fail()
	}
	fmt.Println("hello ivan done")
}

// FailNow
func TestHelloWorldSatria(t *testing.T) {
	value := HelloWorld("satria")

	if value != "hello satria" {
		t.FailNow()
	}
	fmt.Println("hello satria done")
}

func TestHelloWorldJoko(t *testing.T) {
	value := HelloWorld("joko")

	if value != "Hello Joko" {
		t.Error("value tidak 'Hello Joko'")
	}
}
