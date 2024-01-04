package main

import (
	"net/http"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Router struct {
	routes []*Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) add(method string, path string, handler http.HandlerFunc) *Router {
	route := &Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}
	r.routes = append(r.routes, route)
	return r
}

func (r *Router) GET(path string, handler http.HandlerFunc) {
	r.add(http.MethodGet, path, handler)
}

func (r *Router) POST(path string, handler http.HandlerFunc) {
	r.add(http.MethodPost, path, handler)
}

func (r *Router) DELETE(path string, handler http.HandlerFunc) {
	r.add(http.MethodDelete, path, handler)
}

func (r Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.Path == req.URL.Path && route.Method == req.Method {
			route.Handler(res, req)
			return
		}
	}
	http.NotFound(res, req)
}
