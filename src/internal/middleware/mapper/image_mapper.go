package mapper

import (
	"src/internal/api/response"
	"src/internal/repository/model"
)

func ToImageModel(image response.UploadUpdateImage) model.Image {
	return model.Image{
		Image: image.ImageData,
	}
}

func ToImageResponse(user model.Image) response.ImageResponse {
	return response.ImageResponse{
		ID:       user.ID.String(),
		ImageURL: user.Image,
	}
}
