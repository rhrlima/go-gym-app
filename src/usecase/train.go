package usecase

import (
	"gym-app/model"
	"gym-app/repository"
)

type TrainUseCase struct {
	trainRepo repository.TrainRepository
}

func NewTrainUsecase(
	trainRepo repository.TrainRepository,
) TrainUseCase {
	return TrainUseCase{
		trainRepo: trainRepo,
	}
}

func (tu *TrainUseCase) CreateTrain(train model.Train) (model.Train, error) {

	trainId, err := tu.trainRepo.CreateTrain(train)
	if err != nil {
		return model.Train{}, err
	}

	train.ID = trainId

	for i := range train.Sections {
		trainSection, err := tu.CreateTrainSection(train.Sections[i], trainId)
		if err != nil {
			return model.Train{}, err
		}

		train.Sections[i] = trainSection
	}

	return train, nil
}

func (tu *TrainUseCase) CreateTrainSection(trainSection model.TrainSection, train_id int) (model.TrainSection, error) {

	trainSection.TrainID = train_id
	trainSectionId, err := tu.trainRepo.CreateTrainSection(trainSection)
	if err != nil {
		return model.TrainSection{}, err
	}

	trainSection.ID = trainSectionId

	for i := range trainSection.Exercises {
		trainExercise, err := tu.CreateTrainExercise(trainSection.Exercises[i], trainSectionId)
		if err != nil {
			return model.TrainSection{}, err
		}

		trainSection.Exercises[i] = trainExercise
	}

	return trainSection, nil
}

func (tu *TrainUseCase) CreateTrainExercise(trainExercise model.TrainExercise, section_id int) (model.TrainExercise, error) {

	trainExercise.SectionID = section_id
	trainExerciseId, err := tu.trainRepo.CreateTrainExercise(trainExercise)
	if err != nil {
		return model.TrainExercise{}, err
	}

	trainExercise.ID = trainExerciseId

	return trainExercise, nil
}
