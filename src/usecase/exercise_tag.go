package usecase

import (
	"gym-app/model"
	"gym-app/repository"
)

type ExerciseTagUsecase struct {
	repository repository.ExerciseTagRepository
}

func NewExerciseTagUsecase(repository repository.ExerciseTagRepository) ExerciseTagUsecase {
	return ExerciseTagUsecase{
		repository: repository,
	}
}

func (etc *ExerciseTagUsecase) CreateExerciseTag(exerciseTag model.ExerciseTag) (model.ExerciseTag, error) {
	return exerciseTag, etc.repository.CreateExerciseTag(exerciseTag)
}
