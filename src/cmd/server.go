package main

import (
	"knull/api/routes"
	"net/http"
)

func main() {
	r := routes.BaseRouter()

	http.ListenAndServe(":3000", r)
}
