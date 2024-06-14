package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(50 * time.Millisecond)
	return "Done"
}

func otherTask() {
	fmt.Println("Other task")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Other task done")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	// add buffer to
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(100 * time.Millisecond):
		t.Error("timed out")
	}
}
