package domain

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

type Client struct {
	Id   uuid.UUID `json:"id" gorm:"column:id"`
	Name string    `json:"name" gorm:"column:name"`
	Cash float64   `json:"cash" gorm:"column:cash"`
	Card string    `json:"card" gorm:"column:card"`
}

func (c Client) TableName() string {
	return "clients"
}

func (c Client) Validate() error {
	if len(c.Name) > 15 {
		return errors.New("name must be less then 15")
	}
	if c.Cash <= 0 {
		return errors.New("name must not be negative or 0")
	}

	panSplit := strings.Split(c.Card, " ")
	for _, pan := range panSplit {
		if len(pan) != 4 {
			return errors.New("error")
		}
	}

	return nil
}
