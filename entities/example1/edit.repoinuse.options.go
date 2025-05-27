package example1

type Example1OptionsRepoinuse struct {
	Json   RepositoryExample1
	Memory RepositoryExample1
}

var optionsRepoinuseExample1 = Example1OptionsRepoinuse{
	Json:   jsonRepositoryExample1,
	Memory: memoryRepositoryExample1,
}

// editable repoinuse # important
var editRepoinuseExample1 = optionsRepoinuseExample1.Memory
