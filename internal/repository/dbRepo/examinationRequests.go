package dbRepo

import (
	"srbolab/internal/models"
)

func (t *sqliteDBRepo) GetAllExaminationRequests() (*[]models.ExaminationRequest, error) {
	templates := []models.ExaminationRequest{}
	err := t.DB.Select(&templates, `SELECT * FROM examination_requests`, nil)
	if err != nil {
		//log.Instance().Error("error getting formats", nil)
		return nil, err
	}

	return &templates, nil
}

func (t *sqliteDBRepo) GetExaminationRequestByID(vin string) (*models.ExaminationRequest, error) {
	requests := []models.ExaminationRequest{}
	err := t.DB.Select(&requests, `SELECT * FROM examination_requests WHERE vin = $1`, vin)
	if err != nil {
		//log.Instance().Error("error getting formats", nil)
		return nil, err
	}

	if len(requests) == 0 {
		return &models.ExaminationRequest{}, nil
	}

	return &requests[0], nil
}

func (t *sqliteDBRepo) CreateExaminationRequest(req *models.ExaminationRequest) (*models.ExaminationRequest, error) {
	_, err := t.DB.Exec(`INSERT INTO examination_requests VALUES ($1,$2,$3)`, req.VIN, req.RequestDate, req.RequestNumber)
	if err != nil {
		//log.Instance().Error("error inserting format", log.AdditionalFields{"Error": err, "DbKey": formatToAdd.DbKey})
		return nil, err
	}

	savedExaminationRequest, err := t.GetExaminationRequestByID(req.VIN)
	if err != nil {
		//log.Instance().Error("error getting created format", log.AdditionalFields{"Error": err})
		return nil, err
	}

	return savedExaminationRequest, nil
}

func (t *sqliteDBRepo) UpdateExaminationRequest(req *models.ExaminationRequest) (*models.ExaminationRequest, error) {
	//_, err := t.DB.Exec(`UPDATE examination_requests SET $1 = req.EngineNumber, $2 = template.ProductionYear, $3 = template.Manufacturer, $4 =template.Color `)
	//if err != nil {
	//	//log.Instance().Error("error inserting format", log.AdditionalFields{"Error": err, "DbKey": formatToAdd.DbKey})
	//	return nil, err
	//}

	updatedxaminationRequest, err := t.GetExaminationRequestByID(req.VIN)
	if err != nil {
		//log.Instance().Error("error getting created format", log.AdditionalFields{"Error": err})
		return nil, err
	}

	return updatedxaminationRequest, nil
}

func (t *sqliteDBRepo) DeleteExaminationRequest(engineNumber string) error {
	//result, err := db.Exec(`DELETE FROM examination_requests WHERE enginenumber = $1`, engineNumber)
	//if err != nil {
	//	//log.Instance().Error("error deleting template")
	//	return err
	//}
	//
	//numRows, err := result.RowsAffected()
	//if err != nil || numRows == 0 {
	//	//log.Instance().Error("error deleting format", log.AdditionalFields{"Error": err, "DbKey": in.DbKey})
	//	return errors.New(fmt.Sprintf("error deleting template with engine number %s", engineNumber))
	//}

	return nil
}
