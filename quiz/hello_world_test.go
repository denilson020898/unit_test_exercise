package quiz_test

import (
	"fmt"
	"runtime"
	"testing"
	"unit_test_exercise/quiz"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Sebelum test dieksekusi")

	m.Run()

	fmt.Println("Sesudah test dieksekusi")
}

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

func TestSkip(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("unit test only runs on Mac")
	}
	result := quiz.HelloWorld("Denil")
	require.Equal(t, result, "Hello Denil", "they should be equal")
}

func TestSubTest(t *testing.T) {
	t.Run("Denil", func(t *testing.T) {
		result := quiz.HelloWorld("Denil")
		require.Equal(t, result, "Hello Denil", "They should be equal")
	})
	t.Run("Another", func(t *testing.T) {
		result := quiz.HelloWorld("Another")
		require.Equal(t, result, "Hello Another", "They should be equal")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Denilson",
			request:  "Denilson",
			expected: "Hello Denilson",
		}, {
			name:     "Denil",
			request:  "Denil",
			expected: "Hello Denil",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := quiz.HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Denilson)",
			request: "Denilson",
		},
		{
			name:    "HelloWorld(Denil)",
			request: "Denil",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				quiz.HelloWorld(benchmark.request)
			}
		})
	}
	// b.Run("Denil", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		quiz.HelloWorld("Denil")
	// 	}
	// })
}
