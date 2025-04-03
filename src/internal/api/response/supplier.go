package response

type CreateSupplier struct {
	Name    string              `json:"name"`
	Address CreateUpdateAddress `json:"address"`
	Phone   string              `json:"phone_number"`
}

type SupplierResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone_number"`
}
