package queries

import (
	//"fmt"
	"github.com/olivere/elastic"
)

func (q EsQuery) Build() elastic.Query {
	query := elastic.NewBoolQuery()
	equalsQueries := make([]elastic.Query, 0)
	for _, eq := range q.Equals {

		equalsQueries = append(equalsQueries, elastic.NewMatchQuery(eq.Field, eq.Value))
		//fmt.Println(eq)
	}
	query.Must(equalsQueries...)
	return query
	// elastic.NewMatchQuery()
	//
	// boolQuer
}
