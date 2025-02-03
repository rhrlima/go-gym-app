package repository

import (
	"database/sql"
	"fmt"
	"gym-app/internal/model"
)

type TrainSessionRepository struct {
	connection *sql.DB
}

func NewTrainSessionRepository(connection *sql.DB) TrainSessionRepository {
	return TrainSessionRepository{
		connection: connection,
	}
}

func (tr *TrainSessionRepository) CreateTrainSession(trainSession model.TrainSession) (int, error) {

	fmt.Println(trainSession)

	var id int
	query := "INSERT INTO train_sessions (train_section_id) VALUES ($1) RETURNING id"
	stmt, err := tr.connection.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(trainSession.SectionID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (tr *TrainSessionRepository) GetTrainSessions() ([]model.TrainSession, error) {

	query := "SELECT * FROM train_sessions"
	rows, err := tr.connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trainSessions []model.TrainSession
	for rows.Next() {
		var trainSession model.TrainSession
		err = rows.Scan(
			&trainSession.ID,
			&trainSession.SectionID,
			&trainSession.StartedAt,
			&trainSession.EndedAt,
		)
		if err != nil {
			return nil, err
		}
		trainSessions = append(trainSessions, trainSession)
	}

	return trainSessions, nil
}

func (tr *TrainSessionRepository) GetTrainSessionByID(train_session_id int) (model.TrainSession, error) {

	var trainSession model.TrainSession
	query := "SELECT * FROM train_sessions WHERE id = $1"
	err := tr.connection.QueryRow(query, train_session_id).Scan(
		&trainSession.ID,
		&trainSession.SectionID,
		&trainSession.StartedAt,
		&trainSession.EndedAt,
	)
	if err != nil {
		return model.TrainSession{}, err
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return model.TrainSession{}, nil
		}
		return model.TrainSession{}, err
	}

	return trainSession, nil
}

func (tr *TrainSessionRepository) UpdateTrainSession(trainSession model.TrainSession) (model.TrainSession, error) {

	query := "UPDATE train_sessions SET train_section_id = $1, started_at = $2, ended_at = $3 WHERE id = $4"
	_, err := tr.connection.Exec(query,
		trainSession.SectionID,
		trainSession.StartedAt,
		trainSession.EndedAt,
		trainSession.ID,
	)
	if err != nil {
		return model.TrainSession{}, err
	}

	return trainSession, nil
}

func (tr *TrainSessionRepository) DeleteTrainSession(train_session_id int) error {

	query := "DELETE FROM train_sessions WHERE id = $1"
	_, err := tr.connection.Exec(query, train_session_id)
	return err
}
