package routes

import (
	"database/sql"
	"gym-app/controller"
	"gym-app/repository"
	"gym-app/usecase"
)

type Container struct {
	ExerciseController    *controller.ExerciseController
	TagController         *controller.TagController
	ExerciseTagController *controller.ExerciseTagController
	TrainController       *controller.TrainController
}

func NewContainer(db *sql.DB) *Container {
	exerciseRepository := repository.NewExerciseRepository(db)
	tagRepository := repository.NewTagRepository(db)
	exerciseTagRepository := repository.NewExerciseTagRepository(db)
	trainRepository := repository.NewTrainRepository(db)

	exerciseUsecase := usecase.NewExerciseUsecase(exerciseRepository, tagRepository, exerciseTagRepository)
	tagUsecase := usecase.NewTagUsecase(tagRepository)
	exerciseTagUsecase := usecase.NewExerciseTagUsecase(exerciseTagRepository)
	trainUsecase := usecase.NewTrainUsecase(trainRepository)

	ExerciseController := controller.NewExerciseController(exerciseUsecase)
	TagController := controller.NewTagController(tagUsecase)
	exerciseTagController := controller.NewExerciseTagController(exerciseTagUsecase)
	trainController := controller.NewTrainController(trainUsecase)

	return &Container{
		ExerciseController:    &ExerciseController,
		TagController:         &TagController,
		ExerciseTagController: &exerciseTagController,
		TrainController:       &trainController,
	}
}
