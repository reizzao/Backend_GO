package edit

import (
	"github.com/reizzao/backendgo/entities/example2/repositories"
	types "github.com/reizzao/backendgo/entities/example2/type"
)

type Example2OptionsRepoinuse struct {
	Json   types.RepositoryExample2
	Memory types.RepositoryExample2
}

var optionsRepoinuseExample2 = Example2OptionsRepoinuse{
	Json:   repositories.JsonRepositoryExample2,
	Memory: repositories.MemoryRepositoryExample2,
}

// editable repoinuse # important
var EditRepoinuseExample2 = optionsRepoinuseExample2.Memory
