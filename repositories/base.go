package repositories

import (
	"vault_dragon_test/database"
)

var (
	// KeyRepo : Controller will use it to communicate with database
	KeyRepo KeyRepoInterface
)

// InitRepositories : init repos in main.go
func InitRepositories() {
	KeyRepo = NewKeyRepoInterface(database.Client)
}
