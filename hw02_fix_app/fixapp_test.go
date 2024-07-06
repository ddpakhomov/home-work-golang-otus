package main

import (
	"testing"

	"github.com/ddpakhomov/home-work-golang-otus/hw02_fix_app/types"
)

func TestEmployeeString(t *testing.T) {
	employee := types.Employee{UserID: 1, Age: 30, Name: "Alice", DepartmentID: 2}
	expectedOutput := "User ID: 1; Age: 30; Name: Alice; Department ID: 2; "

	actualOutput := employee.String()

	if actualOutput != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, actualOutput)
	}
}
