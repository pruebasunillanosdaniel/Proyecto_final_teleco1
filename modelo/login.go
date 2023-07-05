package modelo

type Loguin_persona struct {
	Numero_identidad uint     `json:"Num_ide"`
	Tipo_ide         Tipo_ide `json:"Tipo_ide"`
	Password         string   `json:"Clave"`
	ID               string   `json:"ID,omitempty"`
}

type Login_Datos struct {
	Login Loguin_persona `json:"Login"`
	Datos Usuario        `json:"Datos"`
}
