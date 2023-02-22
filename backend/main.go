package main

import (
	"backend/data"
	"backend/data/entity"
	"backend/data/repository"
	"context"
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var contactRepository repository.ContactRepository = repository.NewContactRepository(data.GetConnection())

func getContacts(c *gin.Context) {
	contacts, err := contactRepository.GetContacts(context.Background())
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, contacts)
}
func insertContact(c *gin.Context) {
	var newContact entity.Contact
	if err := c.BindJSON(&newContact); err != nil {
		return
	}
	_, err := contactRepository.InsertContact(context.Background(), newContact)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusCreated, newContact)
}

func deleteContact(ctx *gin.Context) {
	var contact entity.Contact
	if err := ctx.BindJSON(&contact); err != nil {
		return
	}
	err := contactRepository.DeleteContact(context.Background(), contact.Id)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	ctx.Status(http.StatusOK)
}

func updateContact(ctx *gin.Context) {
	var updatedContact entity.Contact
	if err := ctx.BindJSON(&updatedContact); err != nil {
		return
	}
	_, err := contactRepository.UpdateContact(context.Background(), updatedContact)
	if err != nil {
		panic(err)
	}
	ctx.IndentedJSON(http.StatusCreated, updatedContact)
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))
	router.GET("/", getContacts)
	router.POST("/contacts", insertContact)
	router.DELETE("/deleteContact", deleteContact)
	router.PUT("/updateContact", updateContact)
	router.Run("0.0.0.0:8080")

}
