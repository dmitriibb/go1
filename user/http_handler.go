package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

func HandlerUser(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		switch r.Method {
		case "PUT":
			saveUser(w, r)
		case "GET":
			getAllUsers(w)
		case "DELETE":
			deleteUser(w, r)
		default:
			fmt.Fprintf(w, "Error - unsupported request method %v", r.Method)
		}
		wg.Done()
	}()
	wg.Wait()
}
func saveUser(w http.ResponseWriter, r *http.Request) {
	user, err := jsonToUser(r)
	if err != nil {
		fmt.Printf("ERROR - %v\n", err)
		fmt.Fprintf(w, "ERROR - %v", err)
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

func deleteUser(w http.ResponseWriter, r *http.Request) {
	user, err := jsonToUser(r)
	if err != nil {
		fmt.Printf("ERROR - %v\n", err)
		fmt.Fprintf(w, "ERROR - %v", err)
		return
	}
	deleted := Delete(user)
	fmt.Fprintf(w, "User deleted: %v", deleted)
}

func jsonToUser(r *http.Request) (user *User, err error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	user = &User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
