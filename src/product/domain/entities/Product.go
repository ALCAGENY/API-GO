package entities 

type Product struct {
	ID		  int    `json:"id"`
	Name		  string `json:"name"`
	Fecha_Adquisicion string `json:"fecha_adquisicion"`
}