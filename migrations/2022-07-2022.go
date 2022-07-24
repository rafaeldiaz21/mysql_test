package migrations

import (
	"mysql_test/database"
	"mysql_test/models"
)

var db = database.NewConnection()

func AddTablesJuegos() {
	db.AutoMigrate(&models.Equipo{}, &models.Partido{}, &models.Jugadores{})
}
