package quiz_test

import (
	"fmt"
	"testing"
	"unit_test_exercise/quiz"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld_Assert(t *testing.T) {
	result := quiz.HelloWorld("Budi")
	assert.Equal(t, result, "Hello Budi", "they should be equal")
	fmt.Println("hehe", "done")
}

func TestHelloWorld_Require(t *testing.T) {
	result := quiz.HelloWorld("Budi")
	require.Equal(t, result, "Hello Budi", "they should be equal")
	fmt.Println("hehe", "no require")
}

func TestHelloWorld(t *testing.T) {
	result := quiz.HelloWorld("Denilson")
	if result != "Hello Denilson" {
		// unit test failed
		// panic("Result is not 'Hello Denilson'")
		t.Error("Result must be 'Hello Denilson'")
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorld_Another(t *testing.T) {
	result := quiz.HelloWorld("Another")
	if result != "Hello Another" {
		// unit test failed
		// panic("Result is not 'Hello Another'")
		t.Fatal("Result must be 'Hello Another'")
	}
	fmt.Println("TestHelloWorldAnother")
}
