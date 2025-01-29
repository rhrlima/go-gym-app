package controller

import (
	"gym-app/internal/usecase"
)

type ExerciseTagController struct {
	exerciseTagUsecase usecase.ExerciseTagUsecase
}

func NewExerciseTagController(usecase usecase.ExerciseTagUsecase) ExerciseTagController {
	return ExerciseTagController{
		exerciseTagUsecase: usecase,
	}
}
