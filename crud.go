package gocrud

import (
	"fmt"
	"net/http"
	"reflect"
	"sort"
	"sync"

	"github.com/vardius/goapi"
	"github.com/vardius/gorepo"
	"golang.org/x/net/context"
)

type Handler interface {
	Handle(context.Context, http.ResponseWriter, *http.Request, *goapi.Context, gorepo.Repository, reflect.Type)
}

var (
	handlersMu sync.RWMutex
	handlers   = make(map[string]Handler)
)

func Register(name string, a Handler) {
	handlersMu.Lock()
	defer handlersMu.Unlock()
	if a == nil {
		panic("handler: Register handler is nil")
	}
	if _, dup := handlers[name]; dup {
		panic("handler: Register called twice for handler " + name)
	}
	handlers[name] = a
}

func Unregister(name string) {
	handlersMu.Lock()
	defer handlersMu.Unlock()
	delete(handlers, name)
}

func Handlers() []string {
	handlersMu.RLock()
	defer handlersMu.RUnlock()
	var list []string
	for name := range handlers {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

func Get(name string) (Handler, error) {
	handlersMu.RLock()
	defer handlersMu.RUnlock()
	handler, ok := handlers[name]
	if !ok {
		return nil, fmt.Errorf("handler: unknown handler %q (forgotten import?)", name)
	}
	return handler, nil
}

func unregisterAll() {
	handlersMu.Lock()
	defer handlersMu.Unlock()
	handlers = make(map[string]Handler)
}
