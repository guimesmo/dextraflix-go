package controllers

import (
	"net/http"
	"context"

	"github.com/labstack/echo"
	db "github.com/guimesmo/dextraflix-go/lib"
	"github.com/guimesmo/dextraflix-go/models"
)

func AuthorCreate(c echo.Context) error {
	collection, _ := db.GetCollection(c, "authors")
	author := new(models.Profile) 

	if err := c.Bind(author); err != nil {
		return err
	}
	item, err := collection.InsertOne(context.TODO(), author)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, item)
}