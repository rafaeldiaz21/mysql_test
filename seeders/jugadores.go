package seeders

import (
	"fmt"
	"mysql_test/controllers/jugador"
	"mysql_test/models"
	"mysql_test/utils"

	"github.com/jaswdr/faker"
	"gorm.io/gorm"
)

func FillJugadores(db *gorm.DB, equipos []models.Equipo, birthYearInit int, numberPlayers int) ([]models.Jugadores, error) {
	var jugadores []models.Jugadores
	var err error
	var faker = faker.New()
	for _, v := range equipos {
		jugadores, err = GenerateJugadores(db, v, birthYearInit, numberPlayers, faker)
		if err != nil {
			fmt.Println(err)
		}
		jugadores, err = jugador.CreateJugador(db, jugadores)
		if err != nil {
			fmt.Println(err)
		}
	}
	return jugadores, err
}

func GenerateJugadores(db *gorm.DB, equipo models.Equipo, birthYear int, num int, faker faker.Faker) (u []models.Jugadores, err error) {
	var list []models.Jugadores
	item := models.Jugadores{}

	for i := 0; i < num; i++ {
		item.Nombre = faker.Person().Name()
		item.IdJugadores = 0
		item.FechaNacimiento = utils.Randate(birthYear, birthYear+5).Format("2006-01-02")
		item.FkEquipos = equipo.IdEquipos
		list = append(list, item)
	}
	return list, nil
}
