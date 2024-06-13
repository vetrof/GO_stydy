package main

import (
	"fmt"
	"net/http"
	"path"
)

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello2)
	http.ListenAndServe(":8888", pr)

}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	for pattern, handleFunc := range p.handlers {
		if ok, err := path.Match(pattern, check); ok && err == nil {
			handleFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
		}
	}
	http.NotFound(res, req)
}

func hello2(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "yohoho"
	}
	fmt.Fprint(res, "hello", name)

}
