package seeders

import (
	"mysql_test/models"
	"mysql_test/utils"

	"gorm.io/gorm"
)

func GeneratePartidos(db *gorm.DB, equipoLocal models.Equipo, equipoVisitante models.Equipo, minYear int, maxYear int, fecha string, number int, ramdom1 int64, ramdom2 int64) (u []models.Partido, err error) {
	var list []models.Partido
	var item = models.Partido{}
	for i := 0; i < number; i++ {

		item.IdPartidos = 0
		item.FkEquipoLocal = equipoLocal.IdEquipos
		item.FkEquipoVisitante = equipoVisitante.IdEquipos
		item.GolesLocal = utils.RandomNum(ramdom1)
		item.GolesVisitante = utils.RandomNum(ramdom2)

		if fecha != "" {
			item.FechaPartido = fecha
		} else {
			item.FechaPartido = utils.Randate(minYear, maxYear).Format("2006-01-02")
		}
		list = append(list, item)

	}
	return list, nil
}
