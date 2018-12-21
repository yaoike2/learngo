package mock

type FakeRetriever struct {
	Contents string
}

func (r *FakeRetriever) Get(url string) string {
	return r.Contents
}

//func (r Retriever) Get(url string) string {
//	return r.Contents
//}
