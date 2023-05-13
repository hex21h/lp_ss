package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var (
	user string = "Andrew"
	pass string = "290"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello %s\n", user)
	fmt.Println(time.Now(), "дернули ручку корневую")
}

func HomeHandlerWithArg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Args: %v\n", vars["args"])
	fmt.Println(time.Now(), "дернули ручку корневую с аргументом: ", vars["args"])
}

func ExitHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "для выхода передайте пароль в виде числа вида ./exit/123")
	fmt.Println(time.Now(), "дернули ручку выхода")
}

func ExitHandlerWithArg(w http.ResponseWriter, r *http.Request) {
	passTry := mux.Vars(r)
	fmt.Println(time.Now(), "дернули ручку выхода с аргументами")
	if passTry["pass"] == pass {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "пароль верный сервер будет отключен")
		fmt.Println(time.Now(), "пароль верный - выключаем сервер")
		os.Exit(0)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "пароль не верный сервер не верный сервер продолжает работать")
		fmt.Println(time.Now(), "пароль не верный - продолжаем работать")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/{args}", HomeHandlerWithArg)
	router.HandleFunc("/exit/", ExitHandler)
	router.HandleFunc("/exit/{pass:[0-9]+}", ExitHandlerWithArg)
	// http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
