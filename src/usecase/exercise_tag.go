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
	
	err := etc.repository.CreateExerciseTag(exerciseTag)
	if err != nil {
		return model.ExerciseTag{}, err
	}

	return exerciseTag, nil
}

func (etc *ExerciseTagUsecase) GetExerciseTags() ([]model.ExerciseTag, error) {
	return etc.repository.GetExerciseTags()
}

func (etc *ExerciseTagUsecase) GetExerciseTagsByExerciseId(exercise_id int) ([]model.ExerciseTag, error) {
	return etc.repository.GetExerciseTagsByExerciseId(exercise_id)
}
