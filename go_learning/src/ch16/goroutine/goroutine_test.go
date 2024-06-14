package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// notice! do not write below in older go earlier than 1.22
		//go func() {
		//	fmt.Println(i)
		//}()

		go func(i int) {
			fmt.Printf("goroutine %d\n", i)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}
