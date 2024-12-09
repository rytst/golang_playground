package main

import (
	"fmt"
	"net/http"
	"regexp"
)

// Path parameter
type routerParam map[string]string

type routerFunc func(routerParam, http.ResponseWriter, *http.Request)

type routerItem struct {
	method  string // Request method
	matcher *regexp.Regexp
	fnc     routerFunc
}

type router struct {
	items []routerItem
}

// Add item of "GET"
func (rt *router) GET(prefix string, fnc routerFunc) {
	rt.items = append(rt.items, routerItem{
		method:  http.MethodGet,
		matcher: regexp.MustCompile(prefix),
		fnc:     fnc,
	})
}

// Add item of "POST"
func (rt *router) POST(prefix string, fnc routerFunc) {
	rt.items = append(rt.items, routerItem{
		method:  http.MethodPost,
		matcher: regexp.MustCompile(prefix),
		fnc:     fnc,
	})
}

func (rt *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, v := range rt.items {
		if v.method == r.Method && v.matcher.MatchString(r.RequestURI) {
			match := v.matcher.FindStringSubmatch(r.RequestURI)
			param := make(routerParam)
			for i, name := range v.matcher.SubexpNames() {
				param[name] = match[i]
			}
			v.fnc(param, w, r)
			return
		}
	}
}

func main() {
	rt := router{}

	rt.GET(`^/$`, func(p routerParam, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})

	rt.GET(`^/(?P<name>\w+)$`, func(p routerParam, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello: %v\n", p["name"])
	})

	rt.GET(`^/fuck/(?P<fuck>\w+)/(?P<name>\w+)$`, func(p routerParam, w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "f**k %v %v\n", p["fuck"], p["name"])
	})

	rt.POST(`^/api$`, func(p routerParam, w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/json")
		fmt.Fprintln(w, `{"status": "OK"}`)
	})

    // Start server
	http.ListenAndServe(":8080", &rt)
}
