package main

import (
	"net/http"

	env "github.com/vardius/gocrud/examples/enviroment"
	"golang.org/x/net/context"
)

func main() {
	env.Log.Critical(context.TODO(), "%s", http.ListenAndServe(":8080", env.Server))
}
