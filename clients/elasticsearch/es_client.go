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
}

type esClient struct {
	client *elastic.Client
}

func Init() {

	client, err := elastic.NewClient(
		elastic.SetBasicAuth("elastic", "changeme"),

		elastic.SetURL("http://192.168.122.104:9200"),
		elastic.SetSniff(false),

		elastic.SetHealthcheckInterval(10*time.Second),
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
