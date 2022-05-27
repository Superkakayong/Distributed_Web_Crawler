package infra

import (
	"io/ioutil"
	"net/http"
)

type Retriever struct{}

func (Retriever) Get(url string) string {
	response, err := http.Get("https://www.imooc.com")

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bytes, _ := ioutil.ReadAll(response.Body)

	return string(bytes)
}
