package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	dto "github.com/ugurkurekci/GoCategoryAPI/dto/category"
	"github.com/ugurkurekci/GoCategoryAPI/service"
)

// CategoryController ...
type CategoryController struct {
	service service.CategoryService
}

// NewCategoryController ...
func NewCategoryController(service service.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Add a new category to the database
// @Tags categories
// @Accept json
// @Produce json
// @Param category body dto.CreateCategoryDto true "Category object"
// @Success 201 {object} dto.CreateCategoryDto
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /categories [post]
func (c *CategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var categoryDto dto.CreateCategoryDto

	if err := json.NewDecoder(r.Body).Decode(&categoryDto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := c.service.CreateCategory(categoryDto); err != nil {
		http.Error(w, "Failed to create category: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(categoryDto) // DTO'yu döndür
}

// GetCategory godoc
// @Summary Get category by ID
// @Description Get a category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} dto.CategoryDto
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /categories/{id} [get]
func (c *CategoryController) GetCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/categories/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	categoryDto, err := c.service.GetCategoryById(id)
	if err != nil {
		http.Error(w, "Category not found: "+err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(categoryDto)
}
