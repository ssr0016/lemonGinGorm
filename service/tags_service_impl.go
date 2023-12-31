package service

import (
	"github.com/go-playground/validator"

	"golang/data/request"
	"golang/data/response"
	"golang/helper"
	"golang/model"
	"golang/repository"

)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagsServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}

//Struct methods

// Create implements TagsService
func (t *TagsServiceImpl) Create(tags request.CreateTagsRequest) {
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}

	t.TagsRepository.Save(tagModel)
}

// Delete implements TagsService
func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

// FindAll implements TagsService
func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()
	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}

// FindById implements TagsService
func (t *TagsServiceImpl) FindById(tagsId int) response.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagsId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}

	return tagResponse
}

// Update implements TagsService
func (t *TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
	t.TagsRepository.FindById(tags.Id)
}
