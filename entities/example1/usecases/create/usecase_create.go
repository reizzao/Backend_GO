package create

import (
	edit "github.com/reizzao/backendgo/entities/example1/edit"
	types "github.com/reizzao/backendgo/entities/example1/type"
)

func Create(model types.Example1Model) types.Example1Model {
	res := prepareCreate(model.Request, edit.EditRepoinuseExample1)
	return res
}
