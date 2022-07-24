package seeders

import (
	"fmt"
	"mysql_test/controllers/equipo"
	"mysql_test/models"

	"github.com/jaswdr/faker"
	"gorm.io/gorm"
)

func FillEquipos(db *gorm.DB, numEquipos int) ([]models.Equipo, error) {
	var equipos []models.Equipo
	var err error
	equipos, err = GenerateEquipos(db, "", numEquipos)
	if err != nil {
		fmt.Println(err)
	}
	equipos, err = equipo.CreateEquipo(db, equipos)
	if err != nil {
		fmt.Println(err)
	}
	return equipos, err
}

func GenerateEquipos(db *gorm.DB, name string, num int) (u []models.Equipo, err error) {
	item := models.Equipo{}
	var list []models.Equipo
	faker := faker.New()
	for i := 0; i < num; i++ {

		if name != "" {
			item.Nombre = name
		} else {
			item.Nombre = faker.Address().City()
		}

		item.IdEquipos = 0
		list = append(list, item)

	}
	return list, nil
}
