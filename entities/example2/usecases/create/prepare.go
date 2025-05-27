package create

import (
	edit "github.com/reizzao/backendgo/entities/example2/edit"
	types "github.com/reizzao/backendgo/entities/example2/type"
)

func prepareCreate(request types.RequestCreateExample2, repoinuse types.RepositoryExample2) types.Example2Model {
	model := types.Example2Model{
		ID:      "todo make ID",
		Request: request,
	}

	saved := edit.EditRepoinuseExample2.Save(model)

	return saved
}
