package example1

type ImplRepositoryPostgresExample1 struct{}

var varImplRepositoryPostgresExample1 = ImplRepositoryMemoryExample1{}

func (impl ImplRepositoryPostgresExample1) createExample1(model Example1Model) Example1Model {
	// todo: implementar - salvar na base de dados - Postgres
	return model
}