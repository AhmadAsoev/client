package services

import (
	"client/pkg/application/repository"
	"client/pkg/domain"
	"github.com/google/uuid"
	"net/http"
)

func UpdateByID(urlID string, client domain.Client) (response domain.Response) {
	id, err := uuid.Parse(urlID)
	if err != nil {
		return domain.Response{
			Code:  http.StatusBadRequest,
			Error: "not correct",
		}
	}

	if err := client.Validate(); err != nil {
		return domain.Response{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		}
	}

	if err := repository.UpdateByID(id, client); err != nil {
		if err.Error() == "client not found" {
			return domain.Response{
				Code:  http.StatusBadRequest,
				Error: "client not found by ID",
			}
		}
		return domain.Response{
			Code:  http.StatusInternalServerError,
			Error: "Can't update",
		}
	}

	return domain.Response{
		Code:    http.StatusOK,
		Message: "Update is successful",
	}
}
