package service

import (
	dto "github.com/ugurkurekci/GoCategoryAPI/dto/category"
	"github.com/ugurkurekci/GoCategoryAPI/repository"
)

type CategoryService interface {
	CreateCategory(categoryDto dto.CreateCategoryDto) error
	GetCategoryById(id int) (dto.CategoryDto, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) CreateCategory(categoryDto dto.CreateCategoryDto) error {
	return s.repo.AddCategory(categoryDto)
}

func (s *categoryService) GetCategoryById(id int) (dto.CategoryDto, error) {
	return s.repo.GetCategoryById(id)
}
