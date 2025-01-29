package controller

import (
	"fmt"
	"gym-app/internal/model"
	"gym-app/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrainController struct {
	trainUseCase usecase.TrainUseCase
}

func NewTrainController(usecase usecase.TrainUseCase) TrainController {
	return TrainController{
		trainUseCase: usecase,
	}
}

func (tc *TrainController) CreateTrain(ctx *gin.Context) {

	var train model.Train
	err := ctx.BindJSON(&train)
	if err != nil {
		response := model.Response{
			Message: "Invalid request body",
		}
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println(train)

	insertedTrain, err := tc.trainUseCase.CreateTrain(train)
	if err != nil {
		response := model.Response{
			Message: "Failed to create train",
		}
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusCreated, insertedTrain)
}
