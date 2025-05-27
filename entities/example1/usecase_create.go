package example1

func Create(model Example1Model) Example1Model {
	res := prepareCreate(model.Request, editRepoinuseExample1)
	return res
}
