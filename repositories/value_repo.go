package repositories

import (
	"time"

	"github.com/jinzhu/gorm"

	"vault_dragon_test/models"
)

// ValueRepoInterface : handles basic database operation
type ValueRepoInterface interface {
	AddValue(m *models.Value) RepositoryResult
	GetLatestValueByKeyBefore(key string, t time.Time) RepositoryResult
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

func (v valueRepoInstrumentedImpl) GetLatestValueByKeyBefore(key string, t time.Time) RepositoryResult {

	var value models.Value
	res := v.db.Table("dragon_values").Joins("left join dragon_keys on dragon_values.key_id = dragon_keys.id").Where(
		"dragon_keys.name = ?", key,
	).Where("dragon_values.created_at < ?", t).Last(&value)

	return RepositoryResult{
		Value:        res.Value,
		Error:        res.Error,
		RowsAffected: res.RowsAffected,
	}
}
