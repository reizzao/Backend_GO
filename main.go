package main

import (
	"fmt"

	tt_create_example1 "github.com/reizzao/backendgo/entities/example1/usecases/create"
	tt_create_example2 "github.com/reizzao/backendgo/entities/example2/usecases/create"
)

func main() {
	fmt.Println("Hello Main!")

	tt_create_example1.TesterCreate()
	tt_create_example2.TesterCreate()
}
