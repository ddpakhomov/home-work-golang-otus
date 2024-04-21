package main

import (
	"fmt"

	"github.com/ddpakhomov/home-work-golang-otus/hw02_fix_app/printer"
	"github.com/ddpakhomov/home-work-golang-otus/hw02_fix_app/reader"
	"github.com/ddpakhomov/home-work-golang-otus/hw02_fix_app/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	} else {
		fmt.Printf("")
	}

	staff, err = reader.ReadJSON(path, -1)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
