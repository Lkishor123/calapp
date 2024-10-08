package calculator

import "fmt"

type ServiceCalObj struct {
	A   int
	B   int
	Ops string
}

type CalFuncInterface interface {
	Add() int
	Sub() int
	Mul() int
	Div() int
}

type CalObj struct {
	CalIntObj CalFuncInterface
}

func (scal *ServiceCalObj) Add() int {
	if scal.Ops != "ADD" {
		fmt.Println("Incorrect Func Call")
		return -1
	}
	return scal.A + scal.B
}

func (scal *ServiceCalObj) Sub() int {
	if scal.Ops != "SUB" {
		fmt.Println("Incorrect Func Call")
		return -1
	}
	return scal.A - scal.B
}

func (scal *ServiceCalObj) Mul() int {
	if scal.Ops != "MUL" {
		fmt.Println("Incorrect Func Call")
		return -1
	}
	return scal.A * scal.B
}

func (scal *ServiceCalObj) Div() int {
	if scal.Ops != "DIV" {
		fmt.Println("Incorrect Func Call")
		return -1
	}
	return scal.A / scal.B
}

func (cal *CalObj) Calculator(A int, B int, Ops string) int {
	// cal.CalIntObj = &ServiceCalObj{A: A, B: B, Ops: Ops}
	if cal.CalIntObj == nil {
		cal.CalIntObj = &ServiceCalObj{A: A, B: B, Ops: Ops}
	}

	switch Ops {
	case "ADD":
		return cal.CalIntObj.Add()
	case "SUB":
		return cal.CalIntObj.Sub()
	case "MUL":
		return cal.CalIntObj.Mul()
	case "DIV":
		return cal.CalIntObj.Div()

	}
	return -1
}
