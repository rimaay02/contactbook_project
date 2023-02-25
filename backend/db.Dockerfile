FROM mysql:latest

COPY ./database/database_contact.sql /docker-entrypoint-initdb.d/