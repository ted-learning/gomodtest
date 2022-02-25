package test

type Retriever struct {
}

func (Retriever) Get(_ string) string {
	return "fake content"
}
