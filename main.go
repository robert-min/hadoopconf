package main

import (
	"net/http"

	"github.com/hadoopconf/go-server/router"
	"github.com/urfave/negroni"
)

func main() {
	m := router.MakeHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	http.ListenAndServe(":8080", n)
}
