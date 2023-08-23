package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userID"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todo Todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo)
}

func Demo2() {
	todo := Todo{1, 2, "Alışveriş yapılacak", false}
	jsontodo, err := json.Marshal(todo)
	response, err := http.Post("https://jsonplaceholder.typicode.com/todos",
		"application/json;charset=utf-8",
		bytes.NewBuffer(jsontodo))
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyByte, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyByte)
	fmt.Println(bodyString)

	var todoresponse Todo
	json.Unmarshal(bodyByte, &todoresponse)
	fmt.Println(todoresponse)
}
