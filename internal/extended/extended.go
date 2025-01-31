package extended

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Extended struct {
	Contents ContentModel
}

func NewExtended(db *sql.DB) Extended {
	return Extended{
		Contents: ContentModel{DB: db},
	}
}
