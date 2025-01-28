package repository

import (
	"database/sql"
	"fmt"
	"gym-app/model"
)

type ExerciseRepository struct {
	connection *sql.DB
}

func NewExerciseRepository(connection *sql.DB) ExerciseRepository {
	return ExerciseRepository{
		connection: connection,
	}
}

func (er *ExerciseRepository) CreateExercise(exercise model.Exercise) (int, error) {

	var id int
	query := "INSERT INTO exercises (name) VALUES ($1) RETURNING id"
	stmt, err := er.connection.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(exercise.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (er *ExerciseRepository) GetExercises() ([]model.Exercise, error) {

	query := "SELECT id, name FROM exercises"

	rows, err := er.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Exercise{}, err
	}
	defer rows.Close()

	var exerciseList []model.Exercise
	var exerciseObject model.Exercise

	for rows.Next() {
		err = rows.Scan(
			&exerciseObject.ID,
			&exerciseObject.Name,
		)

		if err != nil {
			return []model.Exercise{}, err
		}

		exerciseList = append(exerciseList, exerciseObject)
	}

	return exerciseList, nil
}

func (er *ExerciseRepository) GetExerciseById(exercise_id int) (*model.Exercise, error) {
	query, err := er.connection.Prepare(
		"SELECT * FROM exercises WHERE id = $1",
	)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	var exercise model.Exercise

	err = query.QueryRow(exercise_id).Scan(
		&exercise.ID,
		&exercise.Name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &exercise, nil
}
