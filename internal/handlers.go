package internal

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"srbolab/internal/driver"
	"srbolab/internal/repository"
	"srbolab/internal/repository/dbRepo"
)

//Repo is the repository
var Repo *DBRepo

// DBRepo is the db repo
type DBRepo struct {
	DB  repository.DatabaseRepo
}

// NewHandlers creates the handlers
func NewHandlers(repo *DBRepo) {
	Repo = repo
}

// NewSqlitelHandlers creates db repo for postgres
func NewDbHandlers(db *driver.DB) *DBRepo {
	return &DBRepo{
		DB:  dbRepo.NewSqliteRepo(db.SQL),
	}
}


func (repo *DBRepo) GetAllExaminationRequests(w http.ResponseWriter, req *http.Request){
	requests, err := repo.DB.GetAllExaminationRequests()
	log.Info("get all examination requests")
	if err != nil {
		log.Error("error retrieving examination requests")
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, requests)
}

//func (repo *DBRepo) GetTemplateByID(w http.ResponseWriter, req *http.Request){
//	vars := mux.Vars(req)
//	templateID, ok := vars["enginenumber"]
//	if !ok {
//		log.Error("error retrieving template param enginenumber")
//		SetErrorResponse(w, errors.New("template not found"))
//		return
//	}
//
//	templ, err := db.GetByID(templateID)
//	if err != nil {
//		log.Error("error retrieving template")
//		SetErrorResponse(w, err)
//		return
//	}
//
//	SetSuccessResponse(w, templ)
//}
//
//func (repo *DBRepo) CreateTemplate(w http.ResponseWriter, req *http.Request){
//	payloadTemplate := internal.ExaminationRequest{}
//	decoder := json.NewDecoder(req.Body)
//	err := decoder.Decode(&payloadTemplate)
//	if err != nil {
//		log.Error("unable to retrieve the just parsed code")
//		SetErrorResponse(w, errors.New("unable to retrieve the just parsed code"))
//		return
//	}
//
//	template, err := db.Create(&payloadTemplate)
//	if err != nil {
//		log.Error("error creating template")
//		SetErrorResponse(w, errors.New("error creating template"))
//		return
//	}
//
//	SetSuccessResponse(w, template)
//}
//
//func (repo *DBRepo) UpdateTemplate(w http.ResponseWriter, req *http.Request){
//	payload := internal.ExaminationRequest{}
//	decoder := json.NewDecoder(req.Body)
//	err := decoder.Decode(&payload)
//	if err != nil {
//		log.Error("unable to retrieve the just parsed code")
//		SetErrorResponse(w, errors.New("unable to retrieve the just parsed code"))
//		return
//	}
//
//	template, err := db.Update(&payload)
//	if err != nil {
//		log.Error("error updating examination request")
//		SetErrorResponse(w, errors.New("error updating examination request"))
//		return
//	}
//
//	SetSuccessResponse(w, template)
//}
//
//func (repo *DBRepo) DeleteTemplate(w http.ResponseWriter, req *http.Request){
//	vars := mux.Vars(req)
//	templateID, ok := vars["vin"]
//	if !ok {
//		log.Error("error retrieving examination request param vin")
//		SetErrorResponse(w, errors.New("examination request not found"))
//		return
//	}
//
//	err := db.DeleteTemplate(templateID)
//	if err != nil {
//		log.Error("error deleting examination request")
//		SetErrorResponse(w, errors.New("erro deleting examination request"))
//		return
//	}
//
//	SetSuccessResponse(w, nil)
//}
