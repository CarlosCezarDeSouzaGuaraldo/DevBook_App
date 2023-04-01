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

// Create insert a new publication on database
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

// FindPublications get publications from user
func (repository Publications) FindPublications(userID uint64) ([]models.Publication, error) {
	rows, err := repository.db.Query(`
		SELECT DISTINCT p.*, u.nick
		FROM publications p INNER JOIN users u ON u.id = p.author_id
		INNER JOIN followers s ON p.author_id = s.user_id
		WHERE u.id = ? OR s.follower_id = ?
		ORDER BY 1 DESC`,
		userID, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next() {
		var publication models.Publication

		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// FindPublicationByID get a publication
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

func (repository Publications) UpdatePublication(publicationID uint64, publication models.Publication) error {
	statement, err := repository.db.Prepare(`UPDATE publications SET title = ?, content = ? WHERE id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publication.Title, publication.Content, publicationID); err != nil {
		return err
	}

	return nil
}

// DeletePublication remove a publication from database
func (repository Publications) DeletePublication(publicationID uint64) error {
	statement, err := repository.db.Prepare(`DELETE FROM publications WHERE id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}

// FindPublicationByUserID get all publications from a specific user
func (repository Publications) FindPublicationByUserID(userID uint64) ([]models.Publication, error) {
	rows, err := repository.db.Query(
		`SELECT p.*
		FROM publications p INNER JOIN users u ON u.id = p.author_id
		WHERE p.author_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next() {
		var publication models.Publication

		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// LikePublication get all publications from a specific user
func (repository Publications) LikePublication(publicationID uint64) error {
	statement, err := repository.db.Prepare(`UPDATE publications SET likes = likes + 1 WHERE id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}
