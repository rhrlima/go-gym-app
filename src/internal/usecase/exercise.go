package usecase

import (
	"gym-app/internal/model"
	"gym-app/internal/repository"
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

func (eu *ExerciseUsecase) GetExercises() ([]model.Exercise, error) {

	exercises, err := eu.exerciseRepo.GetExercises()
	if err != nil {
		return nil, err
	}

	for i, exercise := range exercises {
		exerciseTags, err := eu.tagRepo.GetTagsByExerciseID(exercise.ID)
		if err != nil {
			return nil, err
		}

		exercises[i].Tags = exerciseTags
	}

	return exercises, nil
}

func (eu *ExerciseUsecase) CreateExercise(exercise model.Exercise) (model.Exercise, error) {

	exerciseId, err := eu.exerciseRepo.CreateExercise(exercise)
	if err != nil {
		return model.Exercise{}, err
	}

	exercise.ID = exerciseId

	for _, tag := range exercise.Tags {
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

func (eu *ExerciseUsecase) UpdateExercise(exercise model.Exercise) (model.Exercise, error) {

	// updates exercise
	err := eu.exerciseRepo.UpdateExercise(exercise)
	if err != nil {
		return model.Exercise{}, err
	}

	// deletes all tags for the exercise
	err = eu.exerciseTagRepo.DeleteExerciseTags(exercise.ID)
	if err != nil {
		return model.Exercise{}, err
	}

	// adds tags to the exercise again
	for _, tag := range exercise.Tags {
		err = eu.exerciseTagRepo.CreateExerciseTag(model.ExerciseTag{
			ExerciseID: exercise.ID,
			TagID:      tag.ID,
		})
		if err != nil {
			return model.Exercise{}, err
		}
	}

	return exercise, nil
}

func (eu *ExerciseUsecase) GetExerciseByID(exercise_id int) (*model.Exercise, error) {
	exercise, err := eu.exerciseRepo.GetExerciseByID(exercise_id)
	if err != nil {
		return nil, err
	}

	if exercise == nil {
		return nil, nil
	}

	exerciseTags, err := eu.tagRepo.GetTagsByExerciseID(exercise_id)
	if err != nil {
		return nil, err
	}

	exercise.Tags = exerciseTags

	return exercise, nil
}

func (eu *ExerciseUsecase) DeleteExercise(exercise_id int) error {
	return eu.exerciseRepo.DeleteExercise(exercise_id)
}
