package services

import (
	"client/pkg/application/repository"
	"client/pkg/domain"
	"net/http"
)

func GetAllClients() (response domain.Response) {

	clients, err := repository.GetAllClients()
	if err != nil {
		return domain.Response{
			Code:    http.StatusInternalServerError,
			Error:   "Can't get clients from db",
			Message: nil,
		}
	}

	return domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: clients,
	}
}
