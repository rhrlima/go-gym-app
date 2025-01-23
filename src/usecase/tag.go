package usecase

import (
	"gym-app/model"
	"gym-app/repository"
)

type TagUsecase struct {
	repository repository.TagRepository
}

func NewTagUsecase(repository repository.TagRepository) TagUsecase {
	return TagUsecase{
		repository: repository,
	}
}

func (tu *TagUsecase) CreateTag(tag model.Tag) (model.Tag, error) {
	tagId, err := tu.repository.CreateTag(tag)
	if err != nil {
		return model.Tag{}, err
	}

	tag.ID = tagId

	return tag, nil
}

func (tu *TagUsecase) GetTags() ([]model.Tag, error) {
	return tu.repository.GetTags()
}

func (tu *TagUsecase) GetTagByName(name string) (*model.Tag, error) {
	tag, err := tu.repository.GetTagByName(name)
	if err != nil {
		return nil, err
	}

	return tag, nil
}
