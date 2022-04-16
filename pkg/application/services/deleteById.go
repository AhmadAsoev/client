package services

import (
	"client/pkg/application/repository"
	"client/pkg/domain"
	"github.com/google/uuid"
	"net/http"
)

func DeleteById(urlId string) (response domain.Response) {

	id, err := uuid.Parse(urlId)
	if err != nil {
		return domain.Response{
			Code:    http.StatusBadRequest,
			Error:   "Can't parse to uuid format",
			Message: nil,
		}
	}

	if err := repository.DeleteById(id); err != nil {
		return domain.Response{
			Code:    http.StatusInternalServerError,
			Error:   "Can't delete by id",
			Message: nil,
		}
	}
	return domain.Response{
		Code:    http.StatusOK,
		Error:   "",
		Message: "Delete success",
	}
}
