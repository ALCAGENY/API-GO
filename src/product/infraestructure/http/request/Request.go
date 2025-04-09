package request

type ProductRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Fecha_Adquisicion string `json:"fecha_adquisicion" validate:"required"`
}