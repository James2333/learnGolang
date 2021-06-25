package go_test_test
import (
	"testing"
	"learn101/go_test"
)

func TestLoop(t *testing.T) {
	t.Log("Loop:", go_test.Loop(uint64(32)))
}

func TestFactorial(t *testing.T) {
	t.Log("Factorial:", go_test.Factorial(uint64(32)))
}

func BenchmarkLoop(b *testing.B) {

	for i := 0; i < b.N; i++ {
		go_test.Loop(uint64(40))
	}
}

func BenchmarkFactorial(b *testing.B) {

	for i := 0; i < b.N; i++ {
		go_test.Factorial(uint64(40))
	}
}
