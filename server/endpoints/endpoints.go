package endpoints

import (
	"backup/server/entities"
	"encoding/json"
	"github.com/gocraft/web"
	"log"
	"net/http"
)

func GetForm (c *ServerContext, w web.ResponseWriter, r *web.Request) {
	forms, err := c.Store.GetForms()
	if err != nil{
		log.Printf("error getting forms: %v", err)
		c.ServeResponse(http.StatusBadRequest, "error getting forms")
		return
	}

	c.ServeResponse(http.StatusOK, forms)
}

func SubmitForm (c *ServerContext, w web.ResponseWriter, r *web.Request) {
	form := &entities.Form{}
	err := json.NewDecoder(r.Body).Decode(form)
	if err != nil {
		log.Printf("error decoding form: %v", err)
		c.ServeResponse(http.StatusBadRequest, "error decoding body")
		return
	}

	err = form.Validate()
	if err != nil {
		log.Printf("error validating: %v", err)
		c.ServeResponse(http.StatusBadRequest, err.Error())
		return
	}

	err = c.Store.SaveForm(form)
	if err != nil {
		log.Printf("error saving form: %v", err)
		c.ServeResponse(http.StatusBadRequest, "error saving form")
		return
	}

	c.ServeResponse(200, "ok")
}
