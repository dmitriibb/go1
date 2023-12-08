package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HandlerUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		saveUser(w, r)
	} else if r.Method == "GET" {
		getAllUsers(w)
	} else {
		fmt.Fprintf(w, "Error - unsupported request method %v", r.Method)
		return
	}
}
func saveUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error - ", err)
		fmt.Fprintf(w, "Error - %v", err)
		return
	}
	user := &User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		fmt.Println("Error - ", err)
		fmt.Fprintf(w, "Error - %v", err)
		return
	}
	Save(user)
}
func getAllUsers(w http.ResponseWriter) {
	all := GetAll()
	err := json.NewEncoder(w).Encode(all)
	if err != nil {
		fmt.Println("Error - ", err)
		fmt.Fprintf(w, "Error - %v", err)
		return
	}
}
