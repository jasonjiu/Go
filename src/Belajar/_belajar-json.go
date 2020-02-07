package main

import (
	"encoding/json"
	"fmt"
)

// User is a representation of a user
type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	// var jsonString = `{"name":"jason jiu", "age": 23}`
	// var jsonString = `[
	// 	{"name":"jason jiu", "age": 23},
	// 	{"Name": "ethan hunt", "Age": 32}
	// 	]`
	var object = []User{{"jason jiu", 27}, {"ethan hunt", 28}}

	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var jsonString = string(jsonData)
	fmt.Println(jsonString)

	// // var jsonData = []byte(jsonString)
	// // var data []User
	// // 	var err = json.Unmarshal([]byte(jsonString), &data)
	// // 	if err != nil {
	// // 	fmt.Println(err.Error())
	// // 	return
	// // }

	// fmt.Println("user :", data[0].FullName)
	// fmt.Println("age :", data[0].Age)
	// fmt.Println("user :", data[1].FullName)
	// fmt.Println("age :", data[1].Age)

}
