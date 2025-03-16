package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

// Error
func TestHelloWorldJoko(t *testing.T) {
	value := HelloWorld("joko")

	if value != "Hello Joko" {
		t.Error("value tidak 'Hello Joko'")
	}
	fmt.Println("hello joko done")
}

func TestHelloWorldEko(t *testing.T) {
	value := HelloWorld("Eko")

	if value != "Hello Eko" {
		t.Fatal("value tidak 'Hello Eko'")
	}
	fmt.Println("hello Eko done")
}

func TestHelloWorldAssert(t *testing.T) {
	value := HelloWorld("ivan")
	assert.Equal(t, "Hello ivan", value, "tidak sama dengan 'hello ivan'")
	fmt.Println("done dengan test assert")
}

func TestHelloWorldRequire(t *testing.T) {
	value := HelloWorld("ivan")
	require.Equal(t, "Hello ivan", value, "tidak sama dengan hello ivan")
	fmt.Println("Test Hello World with require done")
}
