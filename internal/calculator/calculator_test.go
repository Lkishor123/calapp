package calculator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

// Mock object for CalFuncInterface
type MockCalFunc struct {
	mock.Mock
}

func (m *MockCalFunc) Add() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockCalFunc) Sub() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockCalFunc) Mul() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockCalFunc) Div() int {
	args := m.Called()
	return args.Int(0)
}

func TestCalculator_Add(t *testing.T) {
	// Creating mock object
	mockCalFunc := new(MockCalFunc)

	// Set Expectations
	mockCalFunc.On("Add").Return(5)

	// Create the CalObj with the mock
	cal := CalObj{CalIntObj: mockCalFunc}

	// Test the Add function
	result := cal.Calculator(2, 3, "ADD")
	fmt.Println("Results: ", result)
	if result != 5 {
		t.Errorf("Expected 5 but got %d", result)
	}

	mockCalFunc.AssertExpectations(t)
}

func TestCalculator_Sub(t *testing.T) {
	mockCalFunc := new(MockCalFunc)
	mockCalFunc.On("Sub").Return(1)

	cal := CalObj{CalIntObj: mockCalFunc}

	result := cal.Calculator(3, 2, "SUB")
	if result != 1 {
		t.Errorf("Expected 1 but got %d", result)
	}

	mockCalFunc.AssertExpectations(t)
}

func TestCalculator_Mul(t *testing.T) {
	mockCalFunc := new(MockCalFunc)
	mockCalFunc.On("Mul").Return(6)

	cal := CalObj{CalIntObj: mockCalFunc}

	result := cal.Calculator(2, 3, "MUL")
	if result != 6 {
		t.Errorf("Expected 6 but got %d", result)
	}

	mockCalFunc.AssertExpectations(t)
}

func TestCalculator_Div(t *testing.T) {
	mockCalFunc := new(MockCalFunc)
	mockCalFunc.On("Div").Return(2)

	cal := CalObj{CalIntObj: mockCalFunc}

	result := cal.Calculator(6, 3, "DIV")
	if result != 2 {
		t.Errorf("Expected 2 but got %d", result)
	}

	mockCalFunc.AssertExpectations(t)
}

func TestCalculator_DivisionByZero(t *testing.T) {
	mockCalFunc := new(MockCalFunc)
	mockCalFunc.On("Div").Return(-1)

	cal := CalObj{CalIntObj: mockCalFunc}

	result := cal.Calculator(6, 0, "DIV")
	if result != -1 {
		t.Errorf("Expected -1 (division by zero) but got %d", result)
	}

	mockCalFunc.AssertExpectations(t)
}

func TestCalculator_InvalidOps(t *testing.T) {
	mockCalFunc := new(MockCalFunc)
	cal := CalObj{CalIntObj: mockCalFunc}

	result := cal.Calculator(2, 3, "INVALID")
	if result != -1 {
		t.Errorf("Expected -1 for invalid operation but got %d", result)
	}
}
