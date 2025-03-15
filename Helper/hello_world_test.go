package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	value := HelloWorld("ivan")

	if value != "Hello ivan" {
		panic("Value tidak 'Hello ivan")
	}
}
