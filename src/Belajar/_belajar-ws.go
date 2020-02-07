package main

import "fmt"
import "encoding/json"
import "net/http"

type students struct{
	ID string
	Name string
	Grade int
}

var data = []students{
	students{"E001", "ethan", 21},
    students{"W001", "wick", 22},
    students{"B001", "bourne", 23},
    students{"B002", "bond", 23},
}

func users(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST"{
		var result, err = json.Marshal(data)
		if err !=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

	func user(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type","application/json")
		if r.Method == "POST"{
			var id = r.FormValue("id")
			var result []byte
			var err error

			for _, each := range data{
				if each.ID == id{
					result, err = json.Marshal(each)
					if err != nil{
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					w.Write(result)
					return
				}

				http.Error(w, "", http.StatusBadRequest)
			}
		}
	}

	func main(){
		http.HandleFunc("/users", users)
		http.HandleFunc("/user", user)

		fmt.Println("starting web server at http://localhost:8080/")
		http.ListenAndServe(":8080", nil)

	}