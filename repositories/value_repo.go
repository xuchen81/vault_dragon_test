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

func (u valueRepoInstrumentedImpl) AddValue(Value *models.Value) RepositoryResult {
	res := u.db.Save(Value)
	return RepositoryResult{
		Value:        res.Value,
		Error:        res.Error,
		RowsAffected: res.RowsAffected,
	}
}
