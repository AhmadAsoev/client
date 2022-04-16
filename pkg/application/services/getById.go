package services

import (
	"client/pkg/application/repository"
	"client/pkg/domain"
	"github.com/google/uuid"
	"net/http"
)

func GetById(urlId string) (response domain.Response) {

	id, err := uuid.Parse(urlId)
	if err != nil {
		return domain.Response{
			Code:    http.StatusBadRequest,
			Error:   "not correct",
			Message: nil,
		}
	}

	client, err := repository.GetById(id)
	if err != nil {
		return domain.Response{
			Code:    http.StatusInternalServerError,
			Error:   "can'get id from db",
			Message: nil,
		}
	}
	return domain.Response{
		Code:    http.StatusOK,
		Error:   "Success get from Id",
		Message: client,
	}
}
