package response

type CreateProduct struct {
	Name           string  `json:"name"`
	Category       string  `json:"category"`
	Price          float64 `json:"price"`
	AvailableStock int     `json:"available_stock"`
	SupplierID     string  `json:"supplierID"`
}

type ProductResponse struct {
	ID             string  `json:"ID"`
	Name           string  `json:"name"`
	Category       string  `json:"category"`
	Price          float64 `json:"price"`
	AvailableStock int     `json:"available_stock"`
	LastUpdateDate string  `json:"lastUpdateDate"`
	SupplierID     string  `json:"supplierID"`
	ImageID        string  `json:"imageID"`
}
