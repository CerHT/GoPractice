package goroutine

import "testing"

func Test_goroutine(t *testing.T) {
	go sayHello("hello")
	sayHello("world")
}

func Test_goroutineGosched(t *testing.T) {
	go sayWorld("world")
	sayWorld("hello")
}
