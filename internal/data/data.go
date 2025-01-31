package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Extended struct {
	Vendors VendorModel
}

func NewExtended(db *sql.DB) Extended {
	return Extended{
		Vendors: VendorModel{DB: db},
	}
}
