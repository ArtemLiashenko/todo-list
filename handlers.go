package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resRow, rowErr := RepoAllTodo()
	if rowErr != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		if err := json.NewEncoder(w).Encode(resRow); err != nil {
			panic(err)
		}
	}

}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	resRow, rowErr := RepoFindTodo(todoId)

	if rowErr != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		if err := json.NewEncoder(w).Encode(resRow); err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {

	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(http.StatusConflict)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusConflict)
	}
	t, err := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, t)
	}
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	if err := RepoDeleteTodo(todoId); err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}
