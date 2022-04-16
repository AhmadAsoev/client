package services

import (
	"client/pkg/application/repository"
	"client/pkg/domain"
	"github.com/google/uuid"
	"net/http"
)

func AddClient(clients domain.Client) (response domain.Response) {
	if err := clients.Validate(); err != nil {
		return domain.Response{
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
			Message: "",
		}
	}
	id := uuid.New()
	clients.Id = id

	if err := repository.AddClient(clients); err != nil {
		return domain.Response{
			Code:    http.StatusInternalServerError,
			Error:   "can't add into db",
			Message: "",
		}
	}

	return domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: "success add product",
	}
}
