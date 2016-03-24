package controllers

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"testing"
	"github.com/mrtomyum/hr/models"
	"net/url"
)

func TestRunREST(t *testing.T) {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/persons", GetAllPerson),
		//rest.Post("/persons", PostPerson),
		rest.Get("/persons/:id", GetPerson),
		//rest.Delete("/persons/:code", DeletePerson),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func GetAllPerson(w rest.ResponseWriter, r *rest.Request) {
	//lock.RLock()
	db := models.SetupDB("sqlite3", "../models/hr.db")
	persons := []models.Person{}
	db.Debug().Find(&persons)
	//lock.RUnlock()
	w.WriteJson(&persons)
}

func GetPerson(w rest.ResponseWriter, r *rest.Request) {
	escParam := r.PathParam("id")
	decode, err := url.QueryUnescape(escParam)
	if err != nil {
		log.Fatalln(err)
	}
	id := "%" + decode + "%"
	
	db := models.SetupDB("sqlite3", "../models/hr.db")
	persons := []models.Person{}
	if db.Debug().Where("first_name like ? or last_name like ?", id, id).Find(&persons).Error != nil {
	//if db.Debug().Find(&persons, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&persons)
}