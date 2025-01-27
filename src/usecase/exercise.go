package usecase

import (
	"gym-app/model"
	"gym-app/repository"
)

type ExerciseUsecase struct {
	exerciseRepo    repository.ExerciseRepository
	tagRepo         repository.TagRepository
	exerciseTagRepo repository.ExerciseTagRepository
}

func NewExerciseUsecase(
	exerciseRepo repository.ExerciseRepository,
	tagRepo repository.TagRepository,
	exerciseTagRepo repository.ExerciseTagRepository,
) ExerciseUsecase {
	return ExerciseUsecase{
		exerciseRepo:    exerciseRepo,
		tagRepo:         tagRepo,
		exerciseTagRepo: exerciseTagRepo,
	}
}

func (eu *ExerciseUsecase) CreateExercise(exercise model.Exercise) (model.Exercise, error) {

	exerciseId, err := eu.exerciseRepo.CreateExercise(exercise)
	if err != nil {
		return model.Exercise{}, err
	}

	exercise.ID = exerciseId

	for _, tagName := range exercise.Tags {
		tag, err := eu.tagRepo.GetTagByName(tagName)
		if err != nil {
			return model.Exercise{}, err
		}

		err = eu.exerciseTagRepo.CreateExerciseTag(model.ExerciseTag{
			ExerciseID: exercise.ID,
			TagID:      tag.ID,
		})
		if err != nil {
			// TODO undo create exercise
			return model.Exercise{}, err
		}
	}

	return exercise, nil
}

func (eu *ExerciseUsecase) GetExercises() ([]model.Exercise, error) {

	exercises, err := eu.exerciseRepo.GetExercises()
	if err != nil {
		return nil, err
	}

	for i, exercise := range exercises {
		exerciseTags, err := eu.tagRepo.GetTagsByExerciseId(exercise.ID)
		if err != nil {
			return nil, err
		}

		tagNames := []string{}
		for _, tag := range exerciseTags {
			tagNames = append(tagNames, tag.Name)
		}

		exercises[i].Tags = tagNames
	}

	return exercises, nil
}

func (eu *ExerciseUsecase) GetExerciseById(exercise_id int) (*model.Exercise, error) {
	exercise, err := eu.exerciseRepo.GetExerciseById(exercise_id)
	if err != nil {
		return nil, err
	}

	exerciseTags, err := eu.tagRepo.GetTagsByExerciseId(exercise_id)
	if err != nil {
		return nil, err
	}

	tagNames := []string{}
	for _, tag := range exerciseTags {
		tagNames = append(tagNames, tag.Name)
	}

	exercise.Tags = tagNames

	return exercise, nil
}
