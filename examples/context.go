package main

import (
	"time"

	"golang.org/x/net/context"
)

func newContext() (context.Context, context.CancelFunc, error) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err == nil {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}

	return ctx, cancel, nil
}
