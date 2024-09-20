package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	Name string `json:"name"`
}

var userCache = make(map[int]User)

var mutex sync.RWMutex

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", getHomePage)

	mux.HandleFunc("POST /users", createUser)

	mux.HandleFunc("GET /users/{id}", getUser)

	mux.HandleFunc("DELETE /users/{id}", deleteUser)

	fmt.Println("Server is listening at :8080")

	http.ListenAndServe(":8080", mux)
}

func getHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Abhinav Pandey")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	if user.Name == "" {
		http.Error(
			w,
			"User Name is required field",
			http.StatusBadRequest,
		)
		return
	}

	mutex.Lock()
	userCache[len(userCache)+1] = user
	mutex.Unlock()

	w.WriteHeader(http.StatusCreated)
}

func getUser(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)
		return
	}

	mutex.RLock()
	user, ok := userCache[id]
	mutex.RUnlock()

	if !ok {
		http.Error(
			w,
			"User not found",
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(user)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)
		return
	}

	if _, ok := userCache[id]; !ok {
		http.Error(
			w,
			"User not found",
			http.StatusNotFound,
		)
		return
	}

	mutex.Lock()
	delete(userCache, id)
	mutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
