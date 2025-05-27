package create

import (
	"fmt"

	"github.com/reizzao/backendgo/entities/example1/data"
)

func TesterCreate() {
	fmt.Println(
		Create(data.SeedCreateExample1),
	)
}