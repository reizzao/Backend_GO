package create

import (
	edit "github.com/reizzao/backendgo/entities/example1/edit"
	types "github.com/reizzao/backendgo/entities/example1/type"
)

func prepareCreate(request types.RequestCreateExample1, repoinuse types.RepositoryExample1) types.Example1Model {
	model := types.Example1Model{
		ID:      "todo make ID",
		Request: request,
	}

	saved := edit.EditRepoinuseExample1.Save(model)

	return saved
}
