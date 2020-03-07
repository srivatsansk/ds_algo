package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)
type Numverify struct {
	UserId	int `json:"userId"`
  	Id int `json:"id"`
  	Title string `json:"title"`
  	Completed bool `json:"completed"`
}
func main() {

	/*myval := "abc"
	tmpval := url.QueryEscape(myval)*/

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/1")
	
	req, _ := http.NewRequest("GET", url, nil)		
	client := &http.Client{}

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	var record Numverify
	json.NewDecoder(resp.Body).Decode(&record)

	fmt.Println(record.Title)
}