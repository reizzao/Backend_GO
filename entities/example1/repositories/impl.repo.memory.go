package repositories

import types "github.com/reizzao/backendgo/entities/example1/type"

var MemoryRepositoryExample1 = types.RepositoryExample1{

	Save: func(model types.Example1Model) types.Example1Model {
		// todo: implementar - salvar na base de dados - Json
		return model
	},
}
