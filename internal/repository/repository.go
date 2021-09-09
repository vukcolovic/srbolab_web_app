package repository

import (
	"srbolab/internal/models"
)

type DatabaseRepo interface {
	GetAllExaminationRequests() (*[]models.ExaminationRequest, error)
	GetExaminationRequestByID(vin string) (*models.ExaminationRequest, error)
	CreateExaminationRequest(req *models.ExaminationRequest) (*models.ExaminationRequest, error)
	UpdateExaminationRequest(req *models.ExaminationRequest) (*models.ExaminationRequest, error)
	DeleteExaminationRequest(engineNumber string) error
}