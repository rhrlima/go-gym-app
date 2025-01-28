package controller

import (
	"gym-app/usecase"
)

type ExerciseTagController struct {
	exerciseTagUsecase usecase.ExerciseTagUsecase
}

func NewExerciseTagController(usecase usecase.ExerciseTagUsecase) ExerciseTagController {
	return ExerciseTagController{
		exerciseTagUsecase: usecase,
	}
}
