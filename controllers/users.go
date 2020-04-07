package controllers

import (
	"fmt"
	"context"
	"net/http"

	"github.com/labstack/echo"

	db "github.com/guimesmo/dextraflix-go/lib"
	"github.com/guimesmo/dextraflix-go/models"
)

func Home(c echo.Context) error {
	dbclient, _ :=  db.GetClient(c)
	collection, _ := db.GetCollection(c, "users")
	err := dbclient.Ping(context.TODO(), nil)
	if err != nil {
		return c.String(http.StatusBadRequest, "Fail")
	}
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	item, err := collection.InsertOne(context.TODO(), user)

	return c.String(http.StatusOK,
		fmt.Sprintf("name: %s, Email: %s, Passwd: %s\n%v,\n\n\n%v", user.Name, user.Email, user.Password, item, err))
}
