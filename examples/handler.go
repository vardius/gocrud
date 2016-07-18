package main

import (
	"/env"
	"net/http"
	"reflect"
	"time"

	"github.com/vardius/crud"
	"github.com/vardius/goapi"
	"github.com/vardius/gorepo"
)

func NewHandler(hName, rName string, t reflect.Type) http.HandlerFunc {
	hdl, err := crud.Get(hName)
	if err != nil {
		panic(err)
	}

	repo, err := gorepo.Get(rName)
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request, c *goapi.Context) {
		start := time.Now()
		ctx, cancel, err := newContext()
		if err != nil {
			panic(err)
		}
		defer cancel()
		hdl.Handle(ctx, w, r, c, repo, t)
		env.Log.Info(ctx, "%s\t%s\t%d", r.Method, r.RequestURI, time.Since(start))
	}
}
