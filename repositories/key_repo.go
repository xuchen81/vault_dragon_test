package repositories

import (
	"github.com/jinzhu/gorm"

	"vault_dragon_test/models"
)

// KeyRepoInterface : handles basic database operation
type KeyRepoInterface interface {
	AddKey(m *models.Key) RepositoryResult
	GetKeyByKey(k string) RepositoryResult
}

type keyRepoInstrumentedImpl struct {
	db *gorm.DB
}

// NewKeyRepoInterface : creates a new KeyRepoInterface
func NewKeyRepoInterface(db *gorm.DB) KeyRepoInterface {
	return &keyRepoInstrumentedImpl{
		db: db,
	}
}

func (k keyRepoInstrumentedImpl) AddKey(key *models.Key) RepositoryResult {
	res := k.db.Save(key)
	return RepositoryResult{
		Value:        res.Value,
		Error:        res.Error,
		RowsAffected: res.RowsAffected,
	}
}

func (k keyRepoInstrumentedImpl) GetKeyByKey(kname string) RepositoryResult {
	var key models.Key
	res := k.db.Where("name = ?", kname).Find(&key)
	return RepositoryResult{
		Value:        res.Value,
		Error:        res.Error,
		RowsAffected: res.RowsAffected,
	}
}
