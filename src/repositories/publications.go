package repositories

import (
	"api/src/models"
	"database/sql"
)

// Publications represents a publication repository
type Publications struct {
	db *sql.DB
}

// NewPublicationRepository create a publication repository
func NewPublicationRepository(db *sql.DB) *Publications {
	return &Publications{db}
}

func (repository Publications) Create(publication models.Publication) (uint64, error) {
	statement, err := repository.db.Prepare(`
	INSERT INTO publications (title, content, author_id) VALUES (?, ?, ?)`,
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Content, publication.AuthorID)
	if err != nil {
		return 0, err
	}

	lastIDInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIDInserted), nil
}

func (repository Publications) FindPublicationByID(publicationID uint64) (models.Publication, error) {
	rows, err := repository.db.Query(`
		SELECT p.*, u.nick 
		FROM publications p INNER JOIN users u ON u.id = p.author_id
		WHERE p.id = ?`,
		publicationID,
	)
	if err != nil {
		return models.Publication{}, err
	}
	defer rows.Close()

	var publication models.Publication

	if rows.Next() {
		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return models.Publication{}, err
		}
	}

	return publication, nil
}
