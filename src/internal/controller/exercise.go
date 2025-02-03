package controller

import (
	"gym-app/internal/model"
	"gym-app/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExerciseController struct {
	exerciseUsecase usecase.ExerciseUsecase
}

func NewExerciseController(usecase usecase.ExerciseUsecase) ExerciseController {
	return ExerciseController{
		exerciseUsecase: usecase,
	}
}

func (ec *ExerciseController) GetExercises(ctx *gin.Context) {

	exercises, err := ec.exerciseUsecase.GetExercises()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, exercises)
}

func (ec *ExerciseController) CreateExercise(ctx *gin.Context) {

	var exercise model.Exercise
	err := ctx.BindJSON(&exercise)
	if err != nil {
		response := model.Response{
			Message: "Invalid request body",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	insertedExercise, err := ec.exerciseUsecase.CreateExercise(exercise)
	if err != nil {
		response := model.Response{
			Message: "Failed to create exercise",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusCreated, insertedExercise)
}

func (ec *ExerciseController) UpdateExercise(ctx *gin.Context) {

	var exercise model.Exercise
	err := ctx.BindJSON(&exercise)
	if err != nil {
		response := model.Response{
			Message: "Invalid request body",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	insertedExercise, err := ec.exerciseUsecase.UpdateExercise(exercise)
	if err != nil {
		response := model.Response{
			Message: "Failed to update exercise",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusOK, insertedExercise)
}

func (ec *ExerciseController) GetExerciseByID(ctx *gin.Context) {

	id := ctx.Param("exerciseId")

	if id == "" {
		response := model.Response{
			Message: "Exercise ID cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	exerciseId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Exercise ID must be an integer",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	exercise, err := ec.exerciseUsecase.GetExerciseByID(exerciseId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if exercise == nil {
		response := model.Response{
			Message: "Exercise ID not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}

func (ec *ExerciseController) DeleteExercise(ctx *gin.Context) {

	id := ctx.Param("exerciseId")
	if id == "" {
		response := model.Response{
			Message: "Exercise ID is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	exerciseId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Exercise ID must be an integer",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = ec.exerciseUsecase.DeleteExercise(exerciseId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
