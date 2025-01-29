package repository

import (
	"database/sql"
	"gym-app/internal/model"
)

type ExerciseTagRepository struct {
	connection *sql.DB
}

func NewExerciseTagRepository(connection *sql.DB) ExerciseTagRepository {
	return ExerciseTagRepository{
		connection: connection,
	}
}

func (etr *ExerciseTagRepository) CreateExerciseTag(exerciseTag model.ExerciseTag) error {

	tx, err := etr.connection.Begin()
	if err != nil {
		return err
	}

	query, err := etr.connection.Prepare(
		"INSERT INTO exercise_tags (exercise_id, tag_id) VALUES ($1, $2)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer query.Close()

	_, err = query.Exec(exerciseTag.ExerciseID, exerciseTag.TagID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (etr *ExerciseTagRepository) GetExerciseTags() ([]model.ExerciseTag, error) {

	rows, err := etr.connection.Query("SELECT exercise_id, tag_id FROM exercise_tags")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exerciseTags []model.ExerciseTag
	for rows.Next() {
		var exerciseTag model.ExerciseTag
		err := rows.Scan(&exerciseTag.ExerciseID, &exerciseTag.TagID)
		if err != nil {
			return nil, err
		}

		exerciseTags = append(exerciseTags, exerciseTag)
	}

	return exerciseTags, nil
}

func (etr *ExerciseTagRepository) GetExerciseTagsByExerciseId(exercise_id int) ([]model.ExerciseTag, error) {

	query := "SELECT exercise_id, tag_id FROM exercise_tags WHERE exercise_id = $1"
	rows, err := etr.connection.Query(query, exercise_id)
	if err != nil {
		return []model.ExerciseTag{}, err
	}
	defer rows.Close()

	var exerciseTags []model.ExerciseTag
	for rows.Next() {
		var exerciseTag model.ExerciseTag
		err := rows.Scan(&exerciseTag.ExerciseID, &exerciseTag.TagID)
		if err != nil {
			return nil, err
		}

		exerciseTags = append(exerciseTags, exerciseTag)
	}

	return exerciseTags, nil
}
