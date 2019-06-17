package decorator

import "testing"

func Double(n int) int {
	return n * 2
}

func TestLogDecorator(t *testing.T) {
	f := LogDecorator(Double)

	f(5)

}

