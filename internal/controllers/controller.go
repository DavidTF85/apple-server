package controllers

import (
	"net/http"
	"strings"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) HandleRequests(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")[1:]
	n := len(p)

	switch {
	case n == 3 && p[2] == "version" && r.Method == "GET":
		c.getVersion(w, r)
	case n == 3 && p[2] == "hi" && r.Method == "POST":
		c.postHello(w, r)
	default:
		http.NotFound(w, r)
	}
}
