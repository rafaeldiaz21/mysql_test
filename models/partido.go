package models

type Partido struct {
	IdPartidos        int `gorm:"primaryKey"`
	FkEquipoLocal     int
	FkEquipoVisitante int
	GolesLocal        int
	GolesVisitante    int
	FechaPartido      string
	EquipoLocal       Equipo `gorm:"foreignKey:FkEquipoLocal"`
	EquipoVisitante   Equipo `gorm:"foreignKey:FkEquipoVisitante"`
}
