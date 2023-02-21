package repository

import (
	"backend/data"
	"backend/data/entity"
	"context"
	"fmt"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestInsertContact(t *testing.T) {
	contactRepository := NewContactRepository(data.GetConnection())
	ctx := context.Background()
	contact := entity.Contact {
		FirstName: "rima",
		LastName: "a",
		PhoneNumber: "092323",
		Email: "rima@example.com",
	}

	result,err := contactRepository.InsertContact(ctx,contact)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestGetContacts(t *testing.T) {
	contactRepository := NewContactRepository(data.GetConnection())
	contacts,err := contactRepository.GetContacts(context.Background())
	if err != nil {
		panic(err)
	}
	for _,contact := range contacts {
		fmt.Println(contact)
	}
}