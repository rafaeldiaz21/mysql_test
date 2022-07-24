package models

type Jugadores struct {
	IdJugadores     int `gorm:"primaryKey"`
	FkEquipos       int
	Nombre          string
	FechaNacimiento string
	Equipo          Equipo `gorm:"foreignKey:FkEquipos"`
}
