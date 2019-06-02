package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWitParameter(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}

func (m MyType) MethodWitParameter(f float64) {
	fmt.Println("MethodWitParameter called wuth", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}

func (my MyType) MethodNotInterface() {
	fmt.Println("MethodNotInterface called")
}
