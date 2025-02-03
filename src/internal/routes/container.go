package routes

import (
	"database/sql"
	"gym-app/internal/controller"
	"gym-app/internal/repository"
	"gym-app/internal/usecase"
)

type Container struct {
	ExerciseController *controller.ExerciseController
	TagController      *controller.TagController
	// ExerciseTagController  *controller.ExerciseTagController
	TrainController        *controller.TrainController
	TrainSessionController *controller.TrainSessionController
}

func NewContainer(db *sql.DB) *Container {
	exerciseRepository := repository.NewExerciseRepository(db)
	tagRepository := repository.NewTagRepository(db)
	exerciseTagRepository := repository.NewExerciseTagRepository(db)
	trainRepository := repository.NewTrainRepository(db)
	trainSessionRepository := repository.NewTrainSessionRepository(db)

	exerciseUsecase := usecase.NewExerciseUsecase(exerciseRepository, tagRepository, exerciseTagRepository)
	tagUsecase := usecase.NewTagUsecase(tagRepository)
	// exerciseTagUsecase := usecase.NewExerciseTagUsecase(exerciseTagRepository)
	trainUsecase := usecase.NewTrainUsecase(trainRepository)
	trainSession := usecase.NewTrainSessionUsecase(trainSessionRepository)

	ExerciseController := controller.NewExerciseController(exerciseUsecase)
	TagController := controller.NewTagController(tagUsecase)
	// exerciseTagController := controller.NewExerciseTagController(exerciseTagUsecase)
	trainController := controller.NewTrainController(trainUsecase)
	trainSessionController := controller.NewTrainSessionController(trainSession)

	return &Container{
		ExerciseController: &ExerciseController,
		TagController:      &TagController,
		// ExerciseTagController:  &exerciseTagController,
		TrainController:        &trainController,
		TrainSessionController: &trainSessionController,
	}
}
