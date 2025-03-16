package helper

import (
	"fmt"
	"runtime"
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

// Assert
func TestHelloWorldAssert(t *testing.T) {
	value := HelloWorld("ivan")
	assert.Equal(t, "Hello ivan", value, "tidak sama dengan 'hello ivan'")
	fmt.Println("done dengan test assert")
}

// Require
func TestHelloWorldRequire(t *testing.T) {
	value := HelloWorld("ivan")
	require.Equal(t, "Hello ivan", value, "tidak sama dengan hello ivan")
	fmt.Println("Test Hello World with require done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "ivan" {
		t.Skip("can not on winndows")
	}

	value := HelloWorld("skip")
	require.Equal(t, "hello ivan", value, "tidak sama valuenya")
}

//After Before Test

func TestMain(m *testing.M) {
	//Before
	fmt.Println("Before unit test")
	m.Run()

	//after
	fmt.Println("After unit test")
}

//Test Subtest

func TestSubTest(t *testing.T) {
	t.Run("ivan", func(t *testing.T) {
		value := HelloWorld("ivan")
		assert.Equal(t, "Hello joko", value, "tidak sama dengan value")
	})

	t.Run("satria", func(t *testing.T) {
		value := HelloWorld("satria")
		require.Equal(t, "Hello satria", value, "tidak sama dengan value")
	})
}

//Table Test

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "ivan",
			request:  "ivan",
			expected: "Hello ivan",
		},
		{
			name:     "satria",
			request:  "satria",
			expected: "Hello satria",
		},
		{
			name:     "joko",
			request:  "joko",
			expected: "Hello joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value := HelloWorld(test.request)
			require.Equal(t, test.expected, value)
		})
	}
}
