package modelo

type Loguin_persona struct {
	Numero_identidad uint     `json:"num_ide"`
	Tipo_ide         Tipo_ide `json:"tipo_ide"`
	Password         string   `json:"clave"`
}
