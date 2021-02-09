package repository

import "unit_test_exercise/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
