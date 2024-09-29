package repository

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
	dto "github.com/ugurkurekci/GoCategoryAPI/dto/category"
)

type CategoryRepository interface {
	AddCategory(categoryDto dto.CreateCategoryDto) error
	GetCategoryById(id int) (dto.CategoryDto, error)
}

type categoryRepo struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepo{db: db}
}

func (r *categoryRepo) AddCategory(categoryDto dto.CreateCategoryDto) error {
	query := "INSERT INTO Category (Name) VALUES (?)"
	_, err := r.db.Exec(query, categoryDto.Name)
	return err
}
func (r *categoryRepo) GetCategoryById(id int) (dto.CategoryDto, error) {
	query := "SELECT * FROM Category WHERE Id = ?"
	row := r.db.QueryRow(query, id)
	var categoryDto dto.CategoryDto
	err := row.Scan(&categoryDto.Id, &categoryDto.Name)
	return categoryDto, err
}
