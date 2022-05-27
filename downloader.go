package main

import (
	"Distributed_Web_Crawler/infra"
	"fmt"
)

func getRetriever() infra.Retriever {
	return infra.Retriever{}
}

func main() {
	//fmt.Printf("%s\n", retrieve("https://imooc.com"))

	var retriever infra.Retriever = getRetriever()
	//retriever := getRetriever()

	fmt.Println(retriever.Get("https://imooc.com"))
}
