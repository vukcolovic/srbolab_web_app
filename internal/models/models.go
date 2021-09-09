package models

import "time"

type ExaminationRequest struct {
	Id            string    `json:"Id" db:"id"`
	VIN           string    `json:"VIN" db:"vin"`
	RequestDate  time.Time `json:"Date" db:"request_date"`
	RequestNumber string    `json:"RequestNumber" db:"request_number"`
	CreatedAt     time.Time `json:"Created" db:"created_at"`
	UpdatedAt     time.Time `json:"Created" db:"updated_at"`
}
