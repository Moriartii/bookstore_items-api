package services

import (
	"github.com/Moriartii/bookstore_items-api/domain/items"
	"github.com/Moriartii/bookstore_items-api/domain/queries"
	"github.com/Moriartii/bookstore_items-api/utils/errors"
	//"net/http"
)

var (
	//ItemsService itemsService

	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *errors.RestErr)
	Get(string) (*items.Item, *errors.RestErr)
	Search(queries.EsQuery) ([]items.Item, *errors.RestErr)
}

type itemsService struct {
}

func (s *itemsService) Create(item items.Item) (*items.Item, *errors.RestErr) {
	//err := errors.NewStatusNotImplementedError("implement me, please!")
	// return nil, &errors.RestErr{
	// 	Status:  http.StatusNotImplemented,
	// 	Message: "implement me!",
	// 	Error:   "not_implemented",
	// } // errors.NewStatusNotImplementedError("implement me, please!")

	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil

}

func (s *itemsService) Get(id string) (*items.Item, *errors.RestErr) {
	//err := errors.NewStatusNotImplementedError("implement me, please!")
	//return nil, nil //errors.NewStatusNotImplementedError("implement me, please!")

	item := items.Item{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, *errors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)

}
