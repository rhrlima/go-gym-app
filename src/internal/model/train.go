package model

import "time"

type Train struct {
	ID        int            `json:"id"`
	Name      string         `json:"name" binding:"required"`
	Sections  []TrainSection `json:"train_sections" binding:"required"`
	Timestamp time.Time      `json:"created_at"`
}

type TrainSection struct {
	ID        int             `json:"id"`
	Name      string          `json:"name" binding:"required"`
	TrainID   int             `json:"train_id"`
	Exercises []TrainExercise `json:"exercises" binding:"required"`
}

type TrainExercise struct {
	ID         int    `json:"id"`
	SectionID  int    `json:"section_id"`
	ExerciseID int    `json:"exercise_id" binding:"required"`
	Sets       []int  `json:"sets" binding:"required"`
	Comment    string `json:"comment"`
}
