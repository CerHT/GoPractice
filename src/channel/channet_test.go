package channel

import "testing"

func Test_Main(t *testing.T) {
	Main()
}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < 10000; i++ {

	}
	Main()
}
