package repositories

import (
	"github.com/jinzhu/gorm"

	"vault_dragon_test/models"
)

// ValueRepoInterface : handles basic database operation
type ValueRepoInterface interface {
	AddValue(m *models.Value) RepositoryResult
}

type valueRepoInstrumentedImpl struct {
	db *gorm.DB
}

// NewValueRepoInterface : creates a new ValueRepoInterface
func NewValueRepoInterface(db *gorm.DB) ValueRepoInterface {
	return &valueRepoInstrumentedImpl{
		db: db,
	}
}

func (v valueRepoInstrumentedImpl) AddValue(value *models.Value) RepositoryResult {
	res := v.db.Save(value)
	return RepositoryResult{
		Value:        res.Value,
		Error:        res.Error,
		RowsAffected: res.RowsAffected,
	}
}
