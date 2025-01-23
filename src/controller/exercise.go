package controller

import (
	"gym-app/model"
	"gym-app/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type exerciseController struct {
	exerciseUsecase usecase.ExerciseUsecase
}

func NewExerciseController(usecase usecase.ExerciseUsecase) exerciseController {
	return exerciseController{
		exerciseUsecase: usecase,
	}
}

func (ec *exerciseController) CreateExercise(ctx *gin.Context) {

	var exercise model.Exercise
	err := ctx.BindJSON(&exercise)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedExercise, err := ec.exerciseUsecase.CreateExercise(exercise)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedExercise)
}

func (ec *exerciseController) GetExercises(ctx *gin.Context) {

	exercises, err := ec.exerciseUsecase.GetExercises()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, exercises)
}

func (ec *exerciseController) GetExerciseById(ctx *gin.Context) {

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

	exercise, err := ec.exerciseUsecase.GetExerciseById(exerciseId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
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
