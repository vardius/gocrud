package gocrud

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/vardius/gorepo"
	"github.com/vardius/goserver"
	"golang.org/x/net/context"
)

type listAction struct{}

func (act *listAction) Handle(ctx context.Context, w http.ResponseWriter, req *http.Request, c *goserver.Context, repo gorepo.Repository, t reflect.Type) {
	var (
		err  error
		data interface{}
	)
	err = doGetAll(ctx, repo, func(v interface{}, err error) error {
		if err != nil {
			return err
		}
		data = v

		return nil
	})

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	} else if err = json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func doGetAll(ctx context.Context, repo gorepo.Repository, f func(v interface{}, err error) error) error {
	c := make(chan error, 1)
	go func() {
		c <- f(repo.GetAll())
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
	Register("list", &listAction{})
}
