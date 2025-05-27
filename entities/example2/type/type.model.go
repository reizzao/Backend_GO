package types

type Example2Model struct {
	ID      string
	Request RequestCreateExample2
}

type RequestCreateExample2 struct {
	Nome      string
	SobreNome string
}

