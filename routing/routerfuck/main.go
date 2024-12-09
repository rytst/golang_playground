package main

import (
	"fmt"
	"net/http"
	"regex"
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

func main() {

}
