package gocrud

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"

	"github.com/vardius/gorepo"
	"github.com/vardius/goserver"
	"golang.org/x/net/context"
)

type viewAction struct{}

func (act *viewAction) Handle(ctx context.Context, w http.ResponseWriter, req *http.Request, c *goserver.Context, repo gorepo.Repository, t reflect.Type) {
	var (
		data interface{}
		id   int64
		err  error
	)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id, err = strconv.ParseInt(c.Params["id"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = doGet(ctx, repo, id, func(v interface{}, err error) error {
		if err != nil {
			return err
		}
		data = v

		return nil
	})

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	} else if err = json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func doGet(ctx context.Context, repo gorepo.Repository, id int64, f func(v interface{}, err error) error) error {
	c := make(chan error, 1)
	go func() {
		c <- f(repo.Get(id))
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
	Register("view", &viewAction{})
}
