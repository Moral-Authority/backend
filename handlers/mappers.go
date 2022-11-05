package handlers

import (
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/models"
	"strconv"
)

func toUserResponse(user models.User) *model.User {
	return &model.User{
		ID:         strconv.Itoa(int(user.ID)),
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Favourites: toFavsResponse(user.Favourites),
	}
}

func toFavsResponse(favs []models.Favourite) []*model.Favourite {
	var response []*model.Favourite
	for _, e := range favs {
		fav := toFavResponse(e)
		response = append(response, fav)
	}
	return response
}

func toFavResponse(fav models.Favourite) *model.Favourite {
	return &model.Favourite{
		ID:      strconv.Itoa(int(fav.ID)),
		Product: toProductResponse(fav.Product),
	}
}

func toProductResponse(product models.Product) *model.Product {
	return &model.Product{
		ID:          strconv.Itoa(int(product.ID)),
		Title:       product.Title,
		URL:         product.Url,
		Description: product.Description,
		UserID:      strconv.Itoa(int(product.UserId)),
		Image:       toImageResponse(product.Image),
	}
}

func toImageResponse(image models.Image) *model.Image {
	return &model.Image{
		ID:       strconv.Itoa(int(image.ID)),
		Location: image.ImageLocation,
	}
}
