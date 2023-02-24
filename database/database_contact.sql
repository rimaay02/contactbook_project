CREATE DATABASE contact_database;
use contact_database;

CREATE TABLE contact (
    id INT NOT NULL AUTO_INCREMENT,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    phone_number VARCHAR(100),
    email VARCHAR(100),
    PRIMARY KEY (id)
);

INSERT INTO contact(id, first_name, last_name, phone_number, email)
VALUES("1", "JANE", "FONDA", "123321", "JANE@EXAMPLE.COM");