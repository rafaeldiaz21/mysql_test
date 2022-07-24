package models

type Equipo struct {
	IdEquipos int `gorm:"primaryKey"`
	Nombre    string
}
