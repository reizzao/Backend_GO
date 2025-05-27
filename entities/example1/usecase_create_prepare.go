package example1

func prepareCreate(request RequestCreateExample1, repoinuse RepositoryExample1) Example1Model {
	model := Example1Model{
		ID:      "todo make ID",
		Request: request,
	}

	saved := editRepoinuseExample1.save(model)

	return saved
}
