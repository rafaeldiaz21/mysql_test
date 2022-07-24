package equipo

import (
	"mysql_test/models"

	"gorm.io/gorm"
)

func CreateEquipo(db *gorm.DB, item []models.Equipo) (u []models.Equipo, err error) {
	res := db.Create(&item)
	if res.Error != nil {
		return nil, err
	}
	return item, nil
}
