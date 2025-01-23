package repository

import (
	"database/sql"
	"fmt"
	"gym-app/model"
)

type ExerciseRepository struct {
	connection *sql.DB
	tagRepository TagRepository
	exerciseTagRepository ExerciseTagRepository
}

func NewExerciseRepository(connection *sql.DB) ExerciseRepository {
	return ExerciseRepository{
		connection: connection,
		tagRepository: NewTagRepository(connection),
		exerciseTagRepository: NewExerciseTagRepository(connection),
	}
}

func (er *ExerciseRepository) CreateExercise(exercise model.Exercise) (int, error) {

	var id int
	tx, err := er.connection.Begin()
	if err != nil {
		return 0, err
	}

	// insert exercise
	query1, err := er.connection.Prepare(
		"INSERT INTO exercises (name) VALUES ($1) RETURNING id")
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return 0, err
	}
	defer query1.Close()

	err = query1.QueryRow(exercise.Name).Scan(&id)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return 0, err
	}

	for _, tag := range exercise.Tags {
		tag, err := er.tagRepository.GetTagByName(tag)
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
			return 0, err
		}

		err = er.exerciseTagRepository.CreateExerciseTag(model.ExerciseTag{
			ExerciseID: id,
			TagID: tag.ID,
		})
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (er *ExerciseRepository) GetExercises() ([]model.Exercise, error) {

	query := "SELECT id, name FROM exercises;"

	rows, err := er.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Exercise{}, err
	}

	var exerciseList []model.Exercise
	var exerciseObject model.Exercise

	for rows.Next() {
		err = rows.Scan(
			&exerciseObject.ID,
			&exerciseObject.Name,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Exercise{}, err
		}

		exerciseList = append(exerciseList, exerciseObject)

		exerciseTags, err := er.exerciseTagRepository.GetExerciseTagsByExerciseId(exerciseObject.ID)
		if err != nil {
			fmt.Println(err)
			return []model.Exercise{}, err
		}

		tag_names := []string{}
		for _, tag := range tags {
			tag, err := er.tagRepository.GetTagByName(tag)
			tag_names = append(tag_names, tag.Name)
		}
		exerciseObject.Tags = tags
	}

	rows.Close()

	return exerciseList, nil
}

func (er *ExerciseRepository) GetExerciseById(id_exercise int) (*model.Exercise, error) {
	query, err := er.connection.Prepare(
		"SELECT * FROM exercises WHERE id = $1",
	)
	if err != nil {
		return nil, err
	}

	var exercise model.Exercise

	err = query.QueryRow(id_exercise).Scan(
		&exercise.ID,
		&exercise.Name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	query.Close()
	return &exercise, nil
}
