package main

import (
	"net/http"
	routes "petopia-server/router"
)

func main() {
	router := routes.NewRouter()
	http.ListenAndServe(":3000", router)
}
