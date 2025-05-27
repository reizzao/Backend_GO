package edit

import (
	"github.com/reizzao/backendgo/entities/example1/repositories"
	types "github.com/reizzao/backendgo/entities/example1/type"
)

type Example1OptionsRepoinuse struct {
	Json   types.RepositoryExample1
	Memory types.RepositoryExample1
}

var optionsRepoinuseExample1 = Example1OptionsRepoinuse{
	Json:   repositories.JsonRepositoryExample1,
	Memory: repositories.MemoryRepositoryExample1,
}

// editable repoinuse # important
var EditRepoinuseExample1 = optionsRepoinuseExample1.Memory
