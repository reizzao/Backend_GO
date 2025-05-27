package create

import (
	edit "github.com/reizzao/backendgo/entities/example2/edit"
	types "github.com/reizzao/backendgo/entities/example2/type"
)

func Create(model types.Example2Model) types.Example2Model {
	res := prepareCreate(model.Request, edit.EditRepoinuseExample2)
	return res
}
