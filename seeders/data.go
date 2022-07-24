package seeders

import (
	"fmt"
	"mysql_test/controllers/equipo"
	"mysql_test/controllers/partido"
	"mysql_test/database"
	"mysql_test/migrations"
	"mysql_test/models"
)

var db = database.NewConnection()

func InitData(numTemporadas, initYear, numEquipos int) error {
	var err error
	migrations.AddTablesJuegos()

	if err = db.Exec("DELETE FROM jugadores").Error; err != nil {
		return err
	}
	if err = db.Exec("DELETE FROM partidos").Error; err != nil {
		return err
	}
	if err = db.Exec("DELETE FROM equipos").Error; err != nil {
		return err
	}
	if err = CreateData(numTemporadas, initYear, numEquipos); err != nil {
		return err
	}
	return nil

}

// Test data para el ejemplo,
// Este método se debe mejorar y refactorizar. Se hizo de esa manera para forzar los datos
// Con el objeto de responder las preguntas.

func CreateData(temporadas, initYear, numEquipos int) error {
	var db = database.NewConnection()
	var equipos []models.Equipo
	var equipoPrueba []models.Equipo
	var equipoQueNoJuega []models.Equipo
	var partidos []models.Partido
	var juegoLocal []models.Partido
	var juegoVisitante []models.Partido
	var err error
	var fecha1 string = ""
	var fecha2 string = ""
	numEquipos = numEquipos - 1

	/////////////////////////////
	// Creando equipos de prueba
	if equipoPrueba, err = GenerateEquipos(db, "Chacarita", 1); err != nil {
		return err
	}
	if equipos, err = equipo.CreateEquipo(db, equipoPrueba); err != nil {
		return err
	}
	if equipoQueNoJuega, err = GenerateEquipos(db, "Equipo que no Jugó", 1); err != nil {
		return err
	}
	if _, err = equipo.CreateEquipo(db, equipoQueNoJuega); err != nil {
		return err
	}
	if equipos, err = FillEquipos(db, numEquipos); err != nil {
		return err
	}
	equipos = append(equipos, equipoPrueba...)

	//////////////////////////////

	for a := 0; a < temporadas; a++ {

		//Creando Jugadores
		_, err = FillJugadores(db, equipos, 1987, 11)
		if err != nil {
			fmt.Println(err)
		}

		//Creando Partidos
		for i := 0; i < len(equipos)-1; i++ {

			if i == ((len(equipos)/2)-1) && initYear == 2016 {
				fecha1 = "2016-01-01"
			}
			if i == ((len(equipos)/2)-2) && initYear == 2016 {
				fecha2 = "2016-02-12"
			}
			if juegoLocal, err = GeneratePartidos(db, equipos[i], equipos[i+1], initYear, initYear+1, fecha1, 1, 3333333, 9999999); err != nil {
				return err
			}
			if juegoVisitante, err = GeneratePartidos(db, equipos[i+1], equipos[i], initYear, initYear+1, fecha2, 1, 55555, 11111); err != nil {
				return err
			}
			partidos = append(partidos, juegoLocal...)
			partidos = append(partidos, juegoVisitante...)

			if _, err = partido.CreatePartidos(db, partidos); err != nil {
				return err
			}
			partidos = make([]models.Partido, 0)
			fecha1 = ""
			fecha2 = ""
		}
		initYear++
	}
	return nil
}
