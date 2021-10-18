package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Moriartii/bookstore_items-api/domain/items"
	"github.com/Moriartii/bookstore_items-api/domain/queries"
	"github.com/Moriartii/bookstore_items-api/services"
	"github.com/Moriartii/bookstore_items-api/utils/errors"
	"github.com/Moriartii/bookstore_items-api/utils/http_utils"
	"github.com/Moriartii/bookstore_oauth-go/oauth"
	"github.com/gorilla/mux"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.Status)
		if a := json.NewEncoder(w).Encode(err); a != nil {
			fmt.Println("Error json: " + a.Error())
		}
		return
	}
	sellerId := oauth.GetCallerID(r)
	if sellerId == 0 {
		respErr := errors.NewUnauthorizeError("invalid access token")
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
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, getErr := services.ItemsService.Get(itemId)
	if getErr != nil {
		http_utils.RespondError(w, *getErr)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, *apiErr)
		return
	}
	defer r.Body.Close()

	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, *apiErr)
		return
	}

	items, searchErr := services.ItemsService.Search(query)
	if searchErr != nil {
		http_utils.RespondError(w, *searchErr)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, items)

}
