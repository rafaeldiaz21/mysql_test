package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mysql_test/database"
	"mysql_test/models"
	"mysql_test/queries"
	"mysql_test/seeders"
)

//go:generate go run ./generators/queries/main.go
func main() {

	var db = database.NewConnection()

	defer (func() {
		sqlDB, err := db.DB()
		sqlDB.Close()

		if err != nil {
			log.Printf("non-clean exit: %v", err)
		}
	})()

	//Creando Datos de Prueba
	seeders.InitData(3, 2016, 12)

	var jugadoresMayores []models.JugadoresMayoresPorEquipo
	var juegos []models.Juegos
	var goles models.Goles
	var partidosVisitante []models.PartidosVisitantePorEquipo
	var requestQuestion3 []string = []string{"2016-01-01", "2016-02-12"}
	var requestQuestion4 string = "Chacarita"

	fmt.Println("1. ¿Cuál es el jugador más viejo de cada equipo?")
	db.Raw(queries.JugadorMayorEdadPorEquipo).Scan(&jugadoresMayores)
	a, err := json.MarshalIndent(jugadoresMayores, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(a))

	fmt.Println("2. ¿Cuántos partidos jugó de visitante cada equipo? (nota: hay equipos no jugaron ningún partido)?")
	db.Raw(queries.PartidosVisitantePorEquipo).Scan(&partidosVisitante)
	b, err := json.MarshalIndent(partidosVisitante, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	fmt.Println("3. ¿Qué equipos jugaron el 01/01/2016 y el 12/02/2016?")
	db.Raw(queries.JuegosPorFechas, requestQuestion3).Scan(&juegos)
	c, err := json.MarshalIndent(juegos, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(c))

	fmt.Println("4. Diga el total de goles que hizo el equipo “Chacarita” en su historia (como local o visitante)")
	db.Raw(queries.TotalesGolesPorEquipo, requestQuestion4, requestQuestion4, requestQuestion4).Scan(&goles)
	d, err := json.MarshalIndent(goles, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(d))

}
