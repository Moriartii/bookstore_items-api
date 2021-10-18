package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/Moriartii/bookstore_items-api/logger"
	"github.com/olivere/elastic"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	logger.Info("Initing Elastic Search")
	client, err := elastic.NewClient(
		elastic.SetBasicAuth("elastic", "changeme"),

		elastic.SetURL("http://192.168.122.104:9200"),
		elastic.SetSniff(false),

		elastic.SetHealthcheckInterval(10*time.Second),

		// elastic.SetErrorLog(logger.Log), // missing method Printf
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().Index(index).BodyJson(doc).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Get(index string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().Index(index).Id(id).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to get id %s", id), err)
		return nil, err
	}
	if !result.Found {
		return nil, nil
	}
	return result, nil
}

func (c *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()

	// if err := c.client.Search(index).Query(query).Validate(); err != nil {
	// 	fmt.Println("result: " + err.Error())
	// 	return nil, nil
	// }

	result, err := c.client.Search(index).Query(query).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to search documents in index %s", index), err)
		return nil, err
	}
	return result, nil
}
