package repositories

import "database/sql"

// Publications represents a publication repository
type Publications struct {
	db *sql.DB
}

// NewPublicationRepository create a publication repository
func NewPublicationRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

func (repository Publications) Create() {

}
