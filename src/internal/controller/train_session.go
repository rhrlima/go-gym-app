package controller

import (
	"fmt"
	"gym-app/internal/model"
	"gym-app/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TrainSessionController struct {
	trainSessionUsecase usecase.TrainSessionUsecase
}

func NewTrainSessionController(usecase usecase.TrainSessionUsecase) TrainSessionController {
	return TrainSessionController{
		trainSessionUsecase: usecase,
	}
}

func (tc *TrainSessionController) CreateTrainSession(ctx *gin.Context) {

	var trainSession model.TrainSession
	err := ctx.BindJSON(&trainSession)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedTrainSession, err := tc.trainSessionUsecase.CreateTrainSession(trainSession)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedTrainSession)
}

func (tc *TrainSessionController) GetTrainSessions(ctx *gin.Context) {

	trainSessions, err := tc.trainSessionUsecase.GetTrainSessions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, trainSessions)
}

func (tc *TrainSessionController) GetTrainSessionByID(ctx *gin.Context) {

	id := ctx.Param("trainSessionId")
	if id == "" {
		response := model.Response{
			Message: "trainSessionId is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	trainSessionId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "trainSessionId must be an integer",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	trainSession, err := tc.trainSessionUsecase.GetTrainSessionByID(trainSessionId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, trainSession)
}

func (tc *TrainSessionController) UpdateTrainSession(ctx *gin.Context) {

	var trainSession model.TrainSession
	err := ctx.BindJSON(&trainSession)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	updatedTrainSession, err := tc.trainSessionUsecase.UpdateTrainSession(trainSession)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedTrainSession)
}

func (tc *TrainSessionController) DeleteTrainSession(ctx *gin.Context) {

	id := ctx.Param("trainSessionId")
	if id == "" {
		response := model.Response{
			Message: "trainSessionId is required",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	trainSessionId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "trainSessionId must be an integer",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = tc.trainSessionUsecase.DeleteTrainSession(trainSessionId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	response := model.Response{
		Message: "Train session deleted successfully",
	}
	ctx.JSON(http.StatusOK, response)
}
