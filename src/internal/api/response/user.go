package response

type CreateUser struct {
	ClientName    string              `json:"name"`
	ClientSurname string              `json:"surname"`
	Birthday      string              `json:"birthday"`
	Gender        string              `json:"gender"`
	Address       CreateUpdateAddress `json:"address"`
}

type UserResponse struct {
	ID            string `json:"id"`
	ClientName    string `json:"name"`
	ClientSurname string `json:"surname"`
	Birthday      string `json:"birthday"`
	Gender        string `json:"gender"`
	Registration  string `json:"registration_date"`
	Address       string `json:"address"`
}
