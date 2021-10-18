package items

import (
	"github.com/Moriartii/bookstore_items-api/clients/elasticsearch"
	"github.com/Moriartii/bookstore_items-api/utils/errors"
)

const (
	IndexItems = "items"
)

func (i *Item) Save() *errors.RestErr {
	result, err := elasticsearch.Client.Index(IndexItems, i)
	if err != nil {
		return errors.NewInternalServerError("error when trying to save item in DB (ES)")
	}
	i.Id = result.Id
	return nil
}
