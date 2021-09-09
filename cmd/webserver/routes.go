package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"srbolab/internal"
)

func routes() http.Handler {

	r := mux.NewRouter()
	//r.Path("/login").Methods(http.MethodPost).HandlerFunc(auth.Login)

	//authSubRouter := r.PathPrefix("/auth").Subrouter()
	//authSubRouter.Use(auth.CheckUser)

	apiSubRouter := r.PathPrefix("/api").Subrouter()
	templateSubRouter := apiSubRouter.PathPrefix("/examinrequest").Subrouter()

	//FIXME add pagination
	templateSubRouter.Path("/getall").Methods(http.MethodGet).HandlerFunc(internal.Repo.GetAllExaminationRequests)
	//templateSubRouter.Path("/get/{vin:.+}").Methods(http.MethodGet).HandlerFunc(internal.GetTemplateByID)
	//templateSubRouter.Path("/create").Methods(http.MethodPost).HandlerFunc(internal.CreateTemplate)
	//templateSubRouter.Path("/update").Methods(http.MethodPost).HandlerFunc(internal.UpdateTemplate)
	//templateSubRouter.Path("/delete/{enginenumber:.+}").Methods(http.MethodDelete).HandlerFunc(internal.DeleteTemplate)

	return r;
}