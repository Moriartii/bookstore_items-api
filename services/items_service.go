package services

import (
	"github.com/Moriartii/bookstore_items-api/domain/items"
	"github.com/Moriartii/bookstore_items-api/utils/errors"
)

var (
	ItemsService itemsService

	//ItemsService itemsServiceInterface = &itemsService{}

)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, errors.RestErr)
	Get(string) (*items.Item, errors.RestErr)
}

type itemsService struct {
}

func (s *itemsService) Create(items.Item) (*items.Item, *errors.RestErr) {
	//err := errors.NewStatusNotImplementedError("implement me, please!")
	return nil, nil // errors.NewStatusNotImplementedError("implement me, please!")
}

func (s *itemsService) Get(string) (*items.Item, *errors.RestErr) {
	//err := errors.NewStatusNotImplementedError("implement me, please!")
	return nil, nil //errors.NewStatusNotImplementedError("implement me, please!")
}
