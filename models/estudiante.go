package models

// esto datos se tomaron teniendo encuenta la solicitud de una contancia para el caso de la UdeA
type Estudiante struct {
	Id        int    `db:"id" json:"id"`
	Usuario   string `db:"usuario" json:"usuario"`
	Nombre    string `db:"nombre" json:"nombre"`
	Identidad string `db:"identidad" json:"identidad"`
	Programa  string `db:"programa" json:"programa"`
	Semestre  int    `db:"semestre" json:"semestre"`
	Situacion string `db:"situacion" json:"situacion"`
	Creditos  string `db:"creditos" json:"creditos"`
	Nivel     string `db:"nivel" json:"nivel"`
}
