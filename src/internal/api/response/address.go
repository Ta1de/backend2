package response

type CreateUpdateAddress struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
}

type AddressResponse struct {
	Id      string `json:"id"`
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
}
