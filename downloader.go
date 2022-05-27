package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func retrieve(url string) string {
	response, err := http.Get("https://www.imooc.com")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bytes, _ := ioutil.ReadAll(response.Body)

	return string(bytes)
}

func main() {
	//fmt.Printf("%s\n", retrieve("https://imooc.com"))
	fmt.Println(retrieve("https://imooc.com"))
}
