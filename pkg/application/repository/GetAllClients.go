package repository

import "client/pkg/domain"

func GetAllClients() (clients []domain.Client, err error) {
	result := DB.Find(&clients)
	return clients, result.Error
}
