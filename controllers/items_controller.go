package controllers

import (
	"fmt"
	"net/http"

	"github.com/Moriartii/bookstore_items-api/domain/items"
	"github.com/Moriartii/bookstore_items-api/services"
	"github.com/Moriartii/bookstore_oauth-go/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO: Return error json to the user
		return
	}
	item := items.Item{
		Seller: oauth.GetCallerID(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		//TODO: Return error json to user
		return
	}

	fmt.Println(result)
	//TODO: Return created item as json with HTTP status 201 Created
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
