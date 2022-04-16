package repository

import (
	"client/pkg/domain"
	"github.com/google/uuid"
)

func DeleteById(id uuid.UUID) (err error) {

	result := DB.Delete(&domain.Client{}, id)
	return result.Error
}
