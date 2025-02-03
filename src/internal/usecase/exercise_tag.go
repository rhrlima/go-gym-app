package usecase

import (
	"gym-app/internal/model"
	"gym-app/internal/repository"
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

func (etc *ExerciseTagUsecase) GetExerciseTagByExerciseID(exercise_id int) ([]model.ExerciseTag, error) {
	return etc.repository.GetExerciseTagsByExerciseID(exercise_id)
}
