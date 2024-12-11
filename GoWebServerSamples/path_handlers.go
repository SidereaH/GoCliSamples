package main

import (
	"fmt"
	"net/http"
	"path" //← Импорт пакета path для обработки URL-путей
	"strings"
)

// go run .\path_handlers.go
func main() {
	pr := newPathResolver()     // ←  Получение экземпляра маршрутизатора
	pr.Add("GET /hello", hello) //  Отображение путей в функции
	pr.Add("* /goodbye/*", goodbye)
	http.ListenAndServe(":4000", pr) // ←  Передача маршрутизатора HTTP-серверу
}
func newPathResolver() *pathResolver {
	// Создание нового инициализированного объекта pathResolver
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler //Добавление путей для внутреннего поиска
}
func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path       // ← Объединение метода и пути для проверки
	for pattern, handlerFunc := range p.handlers { //←  Обход зарегистрированных путей
		if ok, err := path.Match(pattern, check); ok && err == nil { //←  Проверка соответствия текущего пути одному из зарегистрированных
			handlerFunc(res, req) //←  Вызов функции для обработки пути
			return
		} else if err != nil {
			fmt.Fprint(res, err)
		}
	}
	http.NotFound(res, req) //←  Если для пути не нашлось соответствия, вернуть сообщение о том, что страница не найдена
}
func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Andreyy"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}
func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Andreyy"
	}
	fmt.Fprint(res, "Goodbye ", name)
}
