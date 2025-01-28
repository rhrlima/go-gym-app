package repository

import (
	"database/sql"
	"fmt"
	"gym-app/model"
	"strings"
)

type TrainRepository struct {
	connection *sql.DB
}

func NewTrainRepository(connection *sql.DB) TrainRepository {
	return TrainRepository{
		connection: connection,
	}
}

func (tr *TrainRepository) CreateTrain(train model.Train) (int, error) {

	var id int
	query := "INSERT INTO train (name) VALUES ($1) RETURNING id"
	stmt, err := tr.connection.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(train.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (tr *TrainRepository) CreateTrainSection(trainSection model.TrainSection) (int, error) {

	var id int
	query := "INSERT INTO train_sections (name, train_id) VALUES ($1, $2) RETURNING id"
	stmt, err := tr.connection.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(trainSection.Name, trainSection.TrainID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (tr *TrainRepository) CreateTrainExercise(trainExercise model.TrainExercise) (int, error) {

	var id int
	query := "INSERT INTO train_exercises (section_id, exercise_id, sets, comment) VALUES ($1, $2, $3, $4) RETURNING id"
	stmt, err := tr.connection.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		trainExercise.SectionID,
		trainExercise.ExerciseID,
		strings.Trim(fmt.Sprint(trainExercise.Sets), "[]"),
		trainExercise.Comment,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
