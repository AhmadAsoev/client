package repository

import (
	"client/pkg/domain"
	"errors"
	"github.com/google/uuid"
)

func UpdateByID(id uuid.UUID, client domain.Client) (err error) {
	if DB.First(&domain.Client{}, id).Error != nil {
		return errors.New("client not found")
	}
	return DB.Where(&domain.Client{Id: id}).Updates(domain.Client{
		Name: client.Name,
		Cash: client.Cash,
		Card: client.Card,
	}).Error
}
