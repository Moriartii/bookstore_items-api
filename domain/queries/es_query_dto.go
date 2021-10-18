package queries

type EsQuery struct {
	Equals []FieldValue
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

//
// QUERY:
// {
// 	"equals": [
// 		{
// 		  "field": "status",
// 		  "value": "pending"
// 		},
//		{
// 		  "field": "seller",
// 		  "value": 1
// 		}
// 	]
// }

// {
// 	"any_equals": [
// 		{
// 		  "field": "status",
// 		  "value": "pending"
// 		},
//		{
// 		  "field": "seller",
// 		  "value": 1
// 		}
// 	]
// }
