package response

type UploadUpdateImage struct {
	ID        string `json:"product_id"`
	ImageData []byte `json:"image"`
}

type ImageResponse struct {
	ID       string `json:"id"`
	ImageURL []byte `json:"image_url"`
}
