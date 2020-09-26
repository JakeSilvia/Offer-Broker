package endpoints

import (
	"backup/server/store"
	"backup/server/store/sheets"
	"context"
	"encoding/json"
	"github.com/gocraft/web"
	"log"
	"net/http"
)

type ServerContext struct {
	Context context.Context
	w       web.ResponseWriter
	r       *web.Request
	Store   store.Store
}

func (c *ServerContext) InitServerContext(w web.ResponseWriter, r *web.Request, next web.NextMiddlewareFunc) {
	c.w = w
	c.r = r
	c.Context = context.Background()
	service, err := sheets.NewSheet()
	if err != nil {
		panic(err)
	}
	c.Store = service
	next(w, r)
}

func (c *ServerContext) ServeResponse(status int, value interface{}) {
	c.w.Header().Add("Content-Type", "application/json")
	bts, err := json.Marshal(value)
	if err != nil {
		c.w.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.w.WriteHeader(status)
	_, err = c.w.Write(bts)
	if err != nil {
		log.Printf("error writing response: %+v", err)
	}
}
