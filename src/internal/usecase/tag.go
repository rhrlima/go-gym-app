package usecase

import (
	"gym-app/internal/model"
	"gym-app/internal/repository"
)

type TagUsecase struct {
	repository repository.TagRepository
}

func NewTagUsecase(repository repository.TagRepository) TagUsecase {
	return TagUsecase{
		repository: repository,
	}
}

func (tu *TagUsecase) GetTags() ([]model.Tag, error) {
	return tu.repository.GetTags()
}

func (tu *TagUsecase) CreateTag(tag model.Tag) (model.Tag, error) {
	tagId, err := tu.repository.CreateTag(tag)
	if err != nil {
		return model.Tag{}, err
	}

	tag.ID = tagId

	return tag, nil
}

func (tu *TagUsecase) UpdateTag(tag model.Tag) error {
	return tu.repository.UpdateTag(tag)
}

func (tu *TagUsecase) GetTagByID(tag_id int) (*model.Tag, error) {
	return tu.repository.GetTagByID(tag_id)
}

func (tu *TagUsecase) GetTagByName(name string) (*model.Tag, error) {
	return tu.repository.GetTagByName(name)
}

func (tu *TagUsecase) GetTagsByExerciseID(exerciseId int) ([]model.Tag, error) {
	return tu.repository.GetTagsByExerciseID(exerciseId)
}

func (tu *TagUsecase) DeleteTagByID(tag_id int) error {
	return tu.repository.DeleteTagByID(tag_id)
}
