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

func CheckToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("SessionID")
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "проверка токена - токен нен найжен")
		fmt.Println(time.Now(), "проверка токена - токен нен найжен")
	} else {
		w.WriteHeader(http.StatusOK)
		var cookievalue = cookie.Value
		w.Write([]byte(cookievalue))
		fmt.Println(time.Now(), "проверка токена", cookievalue)
	}

}

func GetToken(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "SessionID", Value: "123", Path: "/", MaxAge: 0, HttpOnly: true}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "выдача токена: имя - %s | значение - %s", cookie.Name, cookie.Value)
	fmt.Println(time.Now(), "выдача токена", cookie)
}

func DeleteToken(w http.ResponseWriter, r *http.Request) {
	// cookie, err := r.Cookie("SessionID")
	// if err != nil {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	fmt.Fprintf(w, "удаление токена - токен нен найжен")
	// 	fmt.Println(time.Now(), "удаление токена - токен нен найжен")
	// } else {
	cookie := http.Cookie{
		Name:       "SessionID",
		Value:      "deleted",
		Path:       "/",
		Domain:     "",
		Expires:    time.Now(),
		RawExpires: "",
		MaxAge:     -1,
		Secure:     false,
		HttpOnly:   true,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "выдача токена: имя - %s | значение - %s", cookie.Name, cookie.Value)
	fmt.Println(time.Now(), "выдача токена", cookie)
	// cookie.Value = "Unuse"
	// cookie.Expires = time.Now()
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "удаление токена")
	// fmt.Println(time.Now(), "удаление токена")
	// }
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/{args}", HomeHandlerWithArg)
	router.HandleFunc("/exit/", ExitHandler)
	router.HandleFunc("/exit/{pass:[0-9]+}", ExitHandlerWithArg)
	router.HandleFunc("/token/", CheckToken)
	router.HandleFunc("/token/get/", GetToken)
	router.HandleFunc("/token/delete/", DeleteToken)
	// http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
