package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct{}

var singleInstance *Singleton
var one sync.Once

func GetSingletonObj() *Singleton {
	one.Do(func() {
		fmt.Println("Singleton created")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%d\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
