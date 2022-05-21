package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func UnmarshalUserModel(data []byte) (UserModel, error) {
	var r UserModel
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UserModel) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UserModel struct {
	Page       int64   `json:"page"`
	PerPage    int64   `json:"per_page"`
	Total      int64   `json:"total"`
	TotalPages int64   `json:"total_pages"`
	Data       []Datum `json:"data"`
	Support    Support `json:"support"`
}

type Datum struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type Support struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

func getUser1() <-chan *UserModel {
	fmt.Println("Get user 1")

	c := make(chan *UserModel)

	go func() {
		_, resp, err := fasthttp.Get(nil, "https://reqres.in/api/users?page=1")
		if err != nil {
			log.Fatalln(err)
		}

		var userModel *UserModel
		json.Unmarshal(resp, &userModel)
		c <- userModel
	}()

	return c
}
func getUser2() <-chan *UserModel {
	fmt.Println("Get user 2")

	c := make(chan *UserModel)

	go func() {
		_, resp, err := fasthttp.Get(nil, "https://reqres.in/api/users?page=2")
		if err != nil {
			log.Fatalln(err)
		}

		var userModel *UserModel
		json.Unmarshal(resp, &userModel)
		c <- userModel
	}()

	return c
}

func main() {
	value1 := <- getUser1()
	value2 := <- getUser2()

	fmt.Println("Get data success")

	for _, val := range value1.Data {
		fmt.Println(val.FirstName)
	}
	
	for _, val := range value2.Data {
		fmt.Println(val.FirstName)
	}
}
