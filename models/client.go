package models

import "time"

type Client struct {
	ClienteID          int       `json:"Cliente_ID"`
	NombreUsuario      string    `json:"Nombre_Usuario"`
	Contrasena         string    `json:"Contrasena"`
	Nombre             string    `json:"Nombre"`
	Apellidos          string    `json:"Apellidos"`
	CorreoElectronico  string    `json:"Correo_Electronico"`
	Edad               int       `json:"Edad"`
	Estatura           float64   `json:"Estatura"`
	Peso               float64   `json:"Peso"`
	IMC                float64   `json:"IMC"`
	GEB                float64   `json:"GEB"`
	ETA                float64   `json:"ETA"`
	FechaCreacion      time.Time `json:"Fecha_Creacion"`
	FechaActualizacion time.Time `json:"Fecha_Actualizacion"`
}
