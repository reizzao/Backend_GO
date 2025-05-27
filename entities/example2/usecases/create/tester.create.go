package create

import (
	"fmt"

	"github.com/reizzao/backendgo/entities/example2/data"
)

func TesterCreate() {
	fmt.Println(
		Create(data.SeedCreateExample2),
	)
}