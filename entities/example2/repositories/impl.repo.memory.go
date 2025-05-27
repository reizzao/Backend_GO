package repositories

import types "github.com/reizzao/backendgo/entities/example2/type"

var MemoryRepositoryExample2 = types.RepositoryExample2{

	Save: func(model types.Example2Model) types.Example2Model {
		// todo: implementar - salvar na base de dados - Json
		return model
	},
}
