package repository

import (
	"database/sql"
	"fmt"
	"gym-app/model"
)

type TagRepository struct {
	connection *sql.DB
}

func NewTagRepository(connection *sql.DB) TagRepository {
	return TagRepository{
		connection: connection,
	}
}

func (tr *TagRepository) CreateTag(tag model.Tag) (int, error) {

	var id int

	query := "INSERT INTO tags (name) VALUES ($1) RETURNING id"

	stmt, err := tr.connection.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(tag.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (tr *TagRepository) GetTags() ([]model.Tag, error) {

	query := "SELECT id, name FROM tags;"

	rows, err := tr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Tag{}, err
	}

	var tagList []model.Tag
	var tagObject model.Tag

	for rows.Next() {
		err = rows.Scan(
			&tagObject.ID,
			&tagObject.Name,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Tag{}, err
		}

		tagList = append(tagList, tagObject)
	}

	rows.Close()

	return tagList, nil
}

func (tr *TagRepository) GetTagByName(name string) (*model.Tag, error) {
	query := "SELECT * FROM tags WHERE name=$1"

	stmt, err := tr.connection.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var tag model.Tag
	err = stmt.QueryRow(name).Scan(
		&tag.ID,
		&tag.Name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &tag, nil
}
