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
	contacts,err:=contactRepository.GetContacts(context.Background()) 
	if err!= nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, contacts)
}
func insertContact(c *gin.Context) {
	var newContact entity.Contact

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newContact); err != nil {
			return
	}

	// Add the new album to the slice.
	_, err := contactRepository.InsertContact(context.Background(), newContact)
	if err!= nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusCreated, newContact)
}
func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
        AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
    }))
	router.GET("/",getContacts )
	router.POST("/contacts", insertContact)
	router.Run("0.0.0.0:8080")

}