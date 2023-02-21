package main

import (
	"fmt"
	"backend/data"
)

var contactRepository ContactRepository = NewContactRepository(data.GetConnection())
func main() {
	fmt.Println("hello world")
}