package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	fmt.Println("goroutine test")
}

func sayHello(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sayWorld(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
