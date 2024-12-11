package main

import (
	"fmt"
	"net/http"
	"strings"
)

// go run .\multiple_handlers.go
func main() {
	http.HandleFunc("/hello", hello)
	//Регистрация обработчиков URL-адресов
	//http://localhost:4000/goodbye/YourName
	http.HandleFunc("/goodbye/", goodbye)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":4000", nil) //← Запуск веб-сервера, подключенного к порту 8080
}
func hello(res http.ResponseWriter, req *http.Request) { //←  Функция-обработчик для пути /hello
	query := req.URL.Query()
	//Получение имени из строки запроса
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) { //←  Функция-обработчик для пути /goodbye/
	path := req.URL.Path
	//Выборка имени из строки запроса
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Goodbye ", name)
}
func homePage(res http.ResponseWriter, req *http.Request) { //←  Функция-обработчик для домашней и неопознанной страницы
	if req.URL.Path != "/" { //Проверка соответствия пути домашней или неопознанной странице
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}
