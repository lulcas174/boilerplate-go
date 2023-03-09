package models

// importar meu pacote database
import "github.com/lulcas174/app-register/database"

func MigrateAll() {
	database.Db.AutoMigrate(
	// inserir as databases que vão ser migradas
	)
}
