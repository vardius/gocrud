package gocrud

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/vardius/gorepo"
	"github.com/vardius/goserver"
	"golang.org/x/net/context"
)

type removeAction struct{}

func (act *removeAction) Handle(ctx context.Context, w http.ResponseWriter, req *http.Request, c *goserver.Context, repo gorepo.Repository, t reflect.Type) {
	var (
		err error
		id  int64
	)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id, err = strconv.ParseInt(c.Params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = doRemove(ctx, repo, id, func(v interface{}, err error) error {
		return err
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func doRemove(ctx context.Context, repo gorepo.Repository, id int64, f func(v interface{}, err error) error) error {
	c := make(chan error, 1)
	go func() {
		c <- f(repo.Remove(id))
	}()
	select {
	case <-ctx.Done():
		<-c
		return ctx.Err()
	case err := <-c:
		return err
	}
}

func init() {
	Register("remove", &removeAction{})
}
