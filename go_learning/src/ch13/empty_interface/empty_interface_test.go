package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("String", s)
		return
	}
	fmt.Println("Unknown Type")
}

func TypeAssertion2(p interface{}) {
	switch t := p.(type) {
	case int:
		fmt.Println("Integer", t)
	case string:
		fmt.Println("String", t)
	default:
		fmt.Println("Unknown Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")

	TypeAssertion2(10)
	TypeAssertion2("10")
}
