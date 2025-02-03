package repository

import (
	"database/sql"
	"gym-app/internal/model"
)

type ExerciseRepository struct {
	connection *sql.DB
}

func NewExerciseRepository(connection *sql.DB) ExerciseRepository {
	return ExerciseRepository{
		connection: connection,
	}
}

func (er *ExerciseRepository) GetExercises() ([]model.Exercise, error) {

	query := "SELECT id, name FROM exercises"
	rows, err := er.connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exerciseList []model.Exercise
	for rows.Next() {
		var exerciseObject model.Exercise
		err = rows.Scan(
			&exerciseObject.ID,
			&exerciseObject.Name,
		)

		if err != nil {
			return nil, err
		}

		exerciseList = append(exerciseList, exerciseObject)
	}

	return exerciseList, nil
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

func (er *ExerciseRepository) UpdateExercise(exercise model.Exercise) error {

	query := "UPDATE exercises SET name = $1 WHERE id = $2"
	_, err := er.connection.Exec(query,
		exercise.Name,
		exercise.ID,
	)

	return err
}

func (er *ExerciseRepository) GetExerciseByID(exercise_id int) (*model.Exercise, error) {

	query := "SELECT * FROM exercises WHERE id = $1"
	stmt, err := er.connection.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var exercise model.Exercise
	err = stmt.QueryRow(exercise_id).Scan(
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

func (er *ExerciseRepository) DeleteExercise(exercise_id int) error {
	query := "DELETE FROM exercises WHERE id = $1"
	_, err := er.connection.Exec(query, exercise_id)
	return err
}
