package repository

import (
	"backend/data/entity"
	"context"
)

type ContactRepository interface {
	GetContacts(ctx context.Context) ([]entity.Contact, error)
	InsertContact(ctx context.Context, contact entity.Contact) (entity.Contact, error)
	DeleteContact(ctx context.Context, id int32) (error)
	UpdateContact(ctx context.Context, contact entity.Contact) (entity.Contact, error)
}