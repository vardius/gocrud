package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/vardius/goapi"
	"github.com/vardius/gocrud"
	"github.com/vardius/gorepo"
	"golang.org/x/net/context"
)

type removeAction struct{}

func (act *removeAction) Handle(ctx context.Context, w http.ResponseWriter, req *http.Request, c *goapi.Context, repo gorepo.Repository, t reflect.Type) {
	var (
		ids  []int64
		err  error
		data interface{}
	)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err = body.Close(); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, ids); err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	err = doRemove(ctx, repo, ids, func(v interface{}, err error) error {
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

func doRemove(ctx context.Context, repo gorepo.Repository, ids []int64, f func(v interface{}, err error) error) error {
	c := make(chan error, 1)
	go func() {
		c <- f(repo.Remove(ids...))
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
	gocrud.Register("remove", &removeAction{})
}
