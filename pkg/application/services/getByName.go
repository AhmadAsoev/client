package services

import (
	"client/pkg/application/repository"
	"client/pkg/domain"
	"net/http"
)

func GetByName(name string) (response domain.Response) {

	clients, err := repository.GetByName(name)
	if err != nil {
		return domain.Response{
			Code:  http.StatusInternalServerError,
			Error: "can't get name from db",
		}
	}
	return domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: clients,
	}
}
