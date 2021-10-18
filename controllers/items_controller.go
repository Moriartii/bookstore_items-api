package controllers

import (
	"encoding/json"
	"github.com/Moriartii/bookstore_items-api/domain/items"
	"github.com/Moriartii/bookstore_items-api/services"
	"github.com/Moriartii/bookstore_items-api/utils/errors"
	"github.com/Moriartii/bookstore_items-api/utils/http_utils"
	"github.com/Moriartii/bookstore_oauth-go/oauth"
	"io/ioutil"
	"net/http"
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
		//http_utils.RespondError(w, *err)
		//http_utils.RespondJson(w, err.Status, err)
		return
	}
	sellerId := oauth.GetCallerID(r)
	if sellerId == 0 {
		respErr := errors.NewUnauthorizeError("invalid request body")
		http_utils.RespondError(w, *respErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, *respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := errors.NewBadRequestError("invalid item json body")
		http_utils.RespondError(w, *respErr)
		return
	}

	itemRequest.Seller = sellerId

	// //TODO: Unmarshal request into the item struct
	// item := items.Item{
	// 	Seller: oauth.GetCallerID(r),
	// }

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, *createErr)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
