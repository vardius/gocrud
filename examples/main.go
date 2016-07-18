package main

import (
	"/env"
	"net/http"

	"golang.org/x/net/context"
)

func main() {
	env.Log.Critical(context.TODO(), "%s", http.ListenAndServe(":8080", env.Server))
}
