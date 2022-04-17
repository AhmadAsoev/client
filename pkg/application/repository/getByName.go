package repository

import "client/pkg/domain"

func GetByName(name string) (clients []domain.Client, err error) {
	return clients, DB.Where(&domain.Client{Name: name}).Find(&clients).Error
}
