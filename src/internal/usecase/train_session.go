package usecase

import (
	"gym-app/internal/model"
	"gym-app/internal/repository"
)

type TrainSessionUsecase struct {
	trainSessionRepo repository.TrainSessionRepository
}

func NewTrainSessionUsecase(repo repository.TrainSessionRepository) TrainSessionUsecase {
	return TrainSessionUsecase{
		trainSessionRepo: repo,
	}
}

func (tu *TrainSessionUsecase) CreateTrainSession(trainSession model.TrainSession) (model.TrainSession, error) {
	trainSessionId, err := tu.trainSessionRepo.CreateTrainSession(trainSession)
	if err != nil {
		return model.TrainSession{}, err
	}

	trainSession.ID = trainSessionId

	return trainSession, nil
}

func (tu *TrainSessionUsecase) GetTrainSessions() ([]model.TrainSession, error) {
	return tu.trainSessionRepo.GetTrainSessions()
}

func (tu *TrainSessionUsecase) GetTrainSessionByID(train_session_id int) (model.TrainSession, error) {
	return tu.trainSessionRepo.GetTrainSessionByID(train_session_id)
}

func (tu *TrainSessionUsecase) UpdateTrainSession(trainSession model.TrainSession) (model.TrainSession, error) {
	return tu.trainSessionRepo.UpdateTrainSession(trainSession)
}

func (tu *TrainSessionUsecase) DeleteTrainSession(train_session_id int) error {
	return tu.trainSessionRepo.DeleteTrainSession(train_session_id)
}
