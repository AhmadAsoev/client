package repository

import (
	"client/pkg/domain"
	"github.com/google/uuid"
)

func GetById(id uuid.UUID) (client domain.Client, err error) {
	result := DB.First(&client, id)
	return client, result.Error
}
