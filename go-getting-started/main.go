package main

import (
	"net/http"

	"github.com/jackaitken/go-core-language-pluralsight/go-getting-started/models/go-getting-started/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}