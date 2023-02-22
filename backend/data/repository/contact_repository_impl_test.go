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
func TestDeleteContact(t *testing.T) {
	contactRepository := NewContactRepository(data.GetConnection())
	ctx := context.Background()
	contact := entity.Contact {
		Id: 8,
	}
	err := contactRepository.DeleteContact(ctx,contact.Id)
	if err != nil {
		panic(err)
	}	
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

func TestUpdateContacts(t *testing.T) {
	contactRepository := NewContactRepository(data.GetConnection())
	contact := entity.Contact {
		FirstName: "abu",
		LastName: "bakar",
		PhoneNumber: "1212121",
		Email: "abu.bakar@example.com",
		Id: 10,

	}
	result,err := contactRepository.UpdateContact(context.Background(),contact)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}