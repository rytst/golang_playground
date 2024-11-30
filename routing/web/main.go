package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gocraft/web"
)

type AppContext struct {
	HelloCount int
}

func (c *AppContext) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HelloCount = 3
	next(rw, req)
}

func main() {

}
