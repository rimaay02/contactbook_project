package repository

import (
	"backend/data/entity"
	"context"
	"database/sql"
)

type contactRepositoryImpl struct {
	DB *sql.DB
}

func NewContactRepository(db *sql.DB) ContactRepository {
	return &contactRepositoryImpl{DB: db}
}

func (repository *contactRepositoryImpl) GetContacts(ctx context.Context) ([]entity.Contact, error) {
	script := "SELECT id, first_name, last_name, phone_number, email FROM contact"

	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var contacts []entity.Contact

	for rows.Next() {
		contact := entity.Contact{}
		rows.Scan(&contact.Id, &contact.FirstName, &contact.LastName, &contact.PhoneNumber, &contact.Email)
		contacts = append(contacts, contact)
	}
	return contacts, nil
}
func (repository *contactRepositoryImpl) InsertContact(ctx context.Context, contact entity.Contact) (entity.Contact, error) {
	script := "INSERT INTO contact (first_name, last_name, phone_number, email) VALUES(?,?,?,?)"

	result, err := repository.DB.ExecContext(ctx, script, contact.FirstName, contact.LastName, contact.PhoneNumber, contact.Email)
	if err != nil {
		return contact, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return contact, err
	}
	contact.Id = int32(id)
	return contact, nil

}
func (repository *contactRepositoryImpl) DeleteContact(ctx context.Context, id int32) (error) {
	script := "DELETE FROM contact WHERE Id = ? LIMIT 1"
	_, err := repository.DB.ExecContext(ctx, script, id)
	if err != nil {
		return err
	}
	return nil	
}
