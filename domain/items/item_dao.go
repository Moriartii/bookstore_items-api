package items

import (
	"encoding/json"
	"fmt"

	"github.com/Moriartii/bookstore_items-api/clients/elasticsearch"
	"github.com/Moriartii/bookstore_items-api/domain/queries"
	"github.com/Moriartii/bookstore_items-api/logger"
	"github.com/Moriartii/bookstore_items-api/utils/errors"
)

const (
	indexItems = "items"
	// typeItem   = "item"
)

func (i *Item) Save() *errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return errors.NewInternalServerError("error when trying to save item in DB (ES)")
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() *errors.RestErr {
	itemId := i.Id

	result, err := elasticsearch.Client.Get(indexItems, i.Id)
	if err != nil {
		return errors.NewNotFoundError(fmt.Sprintf("error when trying to get id %s", i.Id))
	}
	if !result.Found {
		return errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
	}
	logger.Info(string(result.Source))

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return errors.NewInternalServerError("error when trying to parse database response")
	}
	if err := json.Unmarshal(bytes, i); err != nil {
		return errors.NewInternalServerError("error when trying to parse database response")
	}

	i.Id = itemId
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, *errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, errors.NewInternalServerError("error when trying to search doc")
	}
	fmt.Println(result)
	//logger.Info(string(result))

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, errors.NewInternalServerError("error when trying to parse response")
		}
		item.Id = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, errors.NewNotFoundError("no items found given criteria")
	}

	return items, nil
}
