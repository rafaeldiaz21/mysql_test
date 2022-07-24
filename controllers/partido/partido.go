package partido

import (
	"mysql_test/models"

	"gorm.io/gorm"
)

func CreatePartidos(db *gorm.DB, item []models.Partido) (u []models.Partido, err error) {
	res := db.Create(&item)
	if res.Error != nil {
		return nil, err
	}
	return item, nil
}
