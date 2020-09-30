package repositories

import (
	"github.com/jinzhu/gorm"

	"vault_dragon_test/models"
)

// KeyRepoInterface : handles basic database operation
type KeyRepoInterface interface {
	AddKey(m *models.Key) RepositoryResult
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

func (u keyRepoInstrumentedImpl) AddKey(key *models.Key) RepositoryResult {
	res := u.db.Save(key)
	return RepositoryResult{
		Value:        res.Value,
		Error:        res.Error,
		RowsAffected: res.RowsAffected,
	}
}
