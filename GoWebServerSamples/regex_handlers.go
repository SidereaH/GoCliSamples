package main

import (
	"fmt"
	"net/http"
	"regexp" //← Импорт пакета для работы с регулярными выражениями
	"strings"
)

// go run .\regex_handlers.go
func main() {
	rr := newPathResolver()
	rr.Add("GET /hello", hello) // Регистрация путей и функций
	rr.Add("(GET|HEAD) /goodbye(/?[A-Za-z0-9]*)?", goodbye)
	http.ListenAndServe(":4000", rr)
}
func newPathResolver() *regexResolver {
	return &regexResolver{
		handlers: make(map[string]http.HandlerFunc),
		cache:    make(map[string]*regexp.Regexp),
	}
}

type regexResolver struct {
	handlers map[string]http.HandlerFunc
	cache    map[string]*regexp.Regexp //← Сохранение скомпилированных регулярных выражений для повторного использования
}

func (r *regexResolver) Add(regex string, handler http.HandlerFunc) {
	r.handlers[regex] = handler
	cache, _ := regexp.Compile(regex)
	r.cache[regex] = cache
}
func (r *regexResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	// Поиск и вызов функции-обработчика
	for pattern, handlerFunc := range r.handlers {
		if r.cache[pattern].MatchString(check) == true {
			handlerFunc(res, req)
			return
		}
	}
	http.NotFound(res, req) // ← Если соответствие не найдено, вернуть ошибку Page Not Found
}
func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Andrey"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}
func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := ""
	if len(parts) > 2 {
		name = parts[2]
	}
	if name == "" {
		name = "Andrey"
	}
	fmt.Fprint(res, "Goodbye ", name)
}
