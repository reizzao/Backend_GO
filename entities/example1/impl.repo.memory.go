package example1

type ImplRepositoryMemoryExample1 struct{}

var varImplRepositoryMemoryExample1 = ImplRepositoryMemoryExample1{}

func (impl ImplRepositoryMemoryExample1) createExample1(model Example1Model) Example1Model {
	// todo: implementar - salvar na base de dados - Memory
	return model
}