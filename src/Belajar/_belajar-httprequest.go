package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"bytes"
	"net/url"
)

var baseURL = "http://localhost:8080"

type student struct {
	ID    string
	Name  string
	Grade int
}

func fetchUser() ([]student, error) {
	var err error
	var client = &http.Client{}
	var data []student

	request, err := http.NewRequest("POST", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchUserFormData(ID string)(student, error){
	var err error
	var client = &http.Client{}
	var data student

	var param = url.Values{}
	param.Set("id", ID)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/user", payload)
	if err != nil{
		return data, err
	}
	request.Header.Set("Content-Type","application/x-www-form-urlencoded")
	
	response, err := client.Do(request)
	if err != nil{
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil{
		return data, err
	}
	return data, nil
}

func main(){
	var users, err = fetchUser()
	var user1, err1 = fetchUserFormData("E001")
	if err != nil{
		fmt.Println("Error!", err.Error())
		return
	}
	if err1 != nil{
		fmt.Println("Error lagi !", err1.Error())
		return
	}

	for _, each := range users{
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	}

	fmt.Printf("ID(1): %s\t Name(1): %s\t Grade(1): %d\n", user1.ID, user1.Name, user1.Grade)

}
