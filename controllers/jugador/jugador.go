package jugador

import (
	"mysql_test/models"

	"gorm.io/gorm"
)

func CreateJugador(db *gorm.DB, item []models.Jugadores) (u []models.Jugadores, err error) {
	res := db.Create(&item)
	if res.Error != nil {
		return nil, err
	}
	return item, nil
}
