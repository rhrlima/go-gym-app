package usecase

import (
	"gym-app/model"
	"gym-app/repository"
)

type ExerciseUsecase struct {
	repository repository.ExerciseRepository
}

func NewExerciseUsecase(repository repository.ExerciseRepository) ExerciseUsecase {
	return ExerciseUsecase{
		repository: repository,
	}
}

func (eu *ExerciseUsecase) CreateExercise(exercise model.Exercise) (model.Exercise, error) {

	exerciseId, err := eu.repository.CreateExercise(exercise)
	if err != nil {
		return model.Exercise{}, err
	}

	exercise.ID = exerciseId

	return exercise, nil
}

func (eu *ExerciseUsecase) GetExercises() ([]model.Exercise, error) {
	return eu.repository.GetExercises()
}

func (eu *ExerciseUsecase) GetExerciseById(id_exercise int) (*model.Exercise, error) {
	exercise, err := eu.repository.GetExerciseById(id_exercise)
	if err != nil {
		return nil, err
	}

	return exercise, nil
}
