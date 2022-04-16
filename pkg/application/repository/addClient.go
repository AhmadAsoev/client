package repository

import "client/pkg/domain"

func AddClient(client domain.Client) error {
	return DB.Create(client).Error
}
