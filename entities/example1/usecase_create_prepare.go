package example1

func prepareCreate(request RequestCreateExample1) Example1Model {
	model := Example1Model{
		ID:      "todo make ID",
		Request: request,
	}

	saved := varImplRepositoryMemoryExample1.createExample1(model)
	// saved := varImplRepositoryPostgresExample1.createExample1(model)

	return saved
}

// todo : implementar modo de poder escolher o repositorio que vai salvar
// prepare receber o repoinuse

