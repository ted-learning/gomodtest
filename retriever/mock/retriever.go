package mock

import "fmt"

type Retriever struct {
	Content string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Mock Retriever: {contents= %s}", r.Content)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Content = form["Contents"]
	return "ok"
}

func (r *Retriever) Get(_ string) string {
	return r.Content
}
