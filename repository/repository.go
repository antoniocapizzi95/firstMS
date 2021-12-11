package repository

import (
	"context"
	"firstMS/repository/models"
)

type AddressBookRepo interface {
	GetAddressBook(ctx context.Context) (*models.AddressBook, error)
	StoreOnePerson(ctx context.Context, person models.Person) error
}
