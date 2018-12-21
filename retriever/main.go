package main

import (
	"fmt"

	"imooc.com/yaob/learngo/retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	//r = mock.FakeRetriever{"this is a fake imooc.com"}
	//fmt.Printf("%T %v\n", r, r)

	r = &mock.FakeRetriever{"this is a fake imooc.com"}
	fmt.Printf("%T %v\n", r, r)

	if mockFakeRetriever, ok := r.(*mock.FakeRetriever); ok {
		fmt.Println(mockFakeRetriever.Contents)
	} else {
		fmt.Println("not a *mock.FakeRetriever")
	}

	fmt.Println(
		"Try a session with mockRetriever")
	fmt.Println(session(&mockRetriever))
	//r = real2.Retriever{}
	//fmt.Printf("%T %v\n", r, r)

	//r = &real2.Retriever{}
	//fmt.Printf("%T %v\n", r, r)

	//var r2 real2.Retriever
	//r2 = real2.Retriever{}
	//fmt.Printf("%T %v\n", r2, r2)
	//fmt.Println(download(r))
}
