package models

//Todos estos deberían estar cada uno en su archivo.
type PartidosVisitantePorEquipo struct {
	IdEquipos              int    `json:"id_equipos"`
	Equipo                 string `json:"equipo"`
	TotalPartidosVisitante int    `json:"total_partidos_visitante"`
}

type Goles struct {
	IdEquipos      int    `json:"id_equipos"`
	Nombre         string `json:"nombre"`
	GolesLocal     int    `json:"goles_local"`
	GolesVisitante int    `json:"goles_visitante"`
}

type Juegos struct {
	EquipoLocal     string `json:"equipo_local"`
	EquipoVisitante string `json:"equipo_visitante"`
	FechaPartido    string `json:"fecha_partido"`
}

type JugadoresMayoresPorEquipo struct {
	IdJugadores int     `json:"id_jugadores"`
	Nombre      string  `json:"nombre"`
	Edad        float32 `json:"edad"`
	IdEquipos   int     `json:"id_equipos"`
	Equipo      string  `json:"equipo"`
}
