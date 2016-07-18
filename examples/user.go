package main

import (
	"reflect"

	env "github.com/vardius/gocrud/examples/enviroment"
	"github.com/vardius/gorepo"
)

type (
	User struct {
		Id    int64  `json:"id" column:"id"`
		Email string `json:"email" column:"email"`
	}
)

func init() {
	t := reflect.TypeOf(User{})

	gorepo.Register("user", gorepo.NewSQL(env.DB, t))

	env.Server.GET("/users", NewHandler("list", "user", t))
	env.Server.GET("/users/:id:[0-9]+", NewHandler("view", "user", t))
	env.Server.POST("/users", NewHandler("save", "user", t))
	env.Server.DELETE("/users/:id:[0-9]+", NewHandler("remove", "user", t))
}
