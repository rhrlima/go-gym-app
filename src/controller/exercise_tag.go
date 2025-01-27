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

// func (etc *ExerciseTagController) CreateExerciseTag(ctx *gin.Context) {

// 	var exerciseTag model.ExerciseTag
// 	err := ctx.BindJSON(&exerciseTag)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	insertedExerciseTag, err := etc.exerciseTagUsecase.CreateExerciseTag(exerciseTag)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, err)
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, insertedExerciseTag)
// }

// func (etc *ExerciseTagController) GetExerciseTags(ctx *gin.Context) {

// 	exerciseTags, err := etc.exerciseTagUsecase.GetExerciseTags()
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, err)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, exerciseTags)
// }

// func (etc *ExerciseTagController) GetExerciseTagsByExerciseId(ctx *gin.Context) {

// 	id := ctx.Param("exerciseId")

// 	if id == "" {
// 		response := model.Response{
// 			Message: "Exercise Tag ID cannot be null",
// 		}
// 		ctx.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	exerciseId, err := strconv.Atoi(id)
// 	if err != nil {
// 		response := model.Response{
// 			Message: "Exercise ID must be an integer",
// 		}
// 		ctx.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	exerciseTag, err := etc.exerciseTagUsecase.GetExerciseTagsByExerciseId(exerciseId)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, err)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, exerciseTag)
// }
