package main

import (
	"Distributed_Web_Crawler/testing"
	"fmt"
)

func getRetriever() retriever {
	return testing.Retriever{}
}

type retriever interface {
	Get(string) string
}

func main() {
	//fmt.Printf("%s\n", retrieve("https://imooc.com"))

	var r retriever = getRetriever()
	//retriever := getRetriever()

	fmt.Println(r.Get("https://imooc.com"))
}
