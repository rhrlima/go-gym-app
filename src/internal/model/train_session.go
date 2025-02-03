package model

import "time"

type TrainSession struct {
	ID        int       `json:"id"`
	SectionID int       `json:"train_section_id" binding:"required"`
	StartedAt time.Time `json:"started_at" binding:"required"`
	EndedAt   time.Time `json:"ended_at"`
}
