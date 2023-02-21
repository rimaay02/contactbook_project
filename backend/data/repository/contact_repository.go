package repository

import (
	"backend/data/entity"
	"context"
)

type ContactRepository interface {
	GetContacts(ctx context.Context) ([]entity.Contact, error)
	InsertContact(ctx context.Context, contact entity.Contact) (entity.Contact, error)
}