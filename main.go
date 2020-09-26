package main

import (
	"backup/server/endpoints"
	"backup/server/settings"
	"github.com/gocraft/web"
	"log"
	"net/http"
)

func main () {
	router := web.New(endpoints.ServerContext{})
	router.Middleware((*endpoints.ServerContext).InitServerContext)
	router.Subrouter(endpoints.ServerContext{}, "/api").
		Post("/form", endpoints.SubmitForm).
		Get("/form", endpoints.GetForm)

	fs := http.FileServer(http.Dir("./static"))
	mux := http.NewServeMux()
	mux.Handle("/api/", router)
	mux.Handle("/", fs)
	if err := http.ListenAndServe(":" + settings.PORT, mux); err != nil {
		log.Fatal(err)
	}

}
