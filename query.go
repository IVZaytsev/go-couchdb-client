package couchdb

import (
	"encoding/json"
)

type ICDBQuery interface {
	SetSelector(selector map[string]interface{}) *cdbQuery
	SetFields(fields ...string) *cdbQuery
	SetSort(sortField string, st SortType) *cdbQuery
	SetLimit(limit int) *cdbQuery
	SetSkip(skip int) *cdbQuery
	SetIndex(index string) *cdbQuery
	ToString() string
	ToBytes() []byte
}

func (query *cdbQuery) SetSelector(selector map[string]interface{}) *cdbQuery {
	query.Selector = selector
	return query
}

func (query *cdbQuery) SetFields(fields ...string) *cdbQuery {
	f := []string{}
	for _, field := range fields {
		if field != "" {
			f = append(f, field)
		}
	}
	query.Fields = f
	return query
}

func (query *cdbQuery) SetSort(sortField string, st SortType) *cdbQuery {
	query.Sort = append(query.Sort, map[string]string{sortField: string(st)})
	return query
}

func (query *cdbQuery) SetLimit(limit int) *cdbQuery {
	query.Limit = limit
	return query
}

func (query *cdbQuery) SetSkip(skip int) *cdbQuery {
	query.Skip = skip
	return query
}

func (query *cdbQuery) SetIndex(index string) *cdbQuery {
	query.UseIndex = index
	return query
}

func (query *cdbQuery) ToString() string {
	queryString, _ := json.Marshal(query)
	return string(queryString)
}

func (query *cdbQuery) ToBytes() []byte {
	queryBytes, _ := json.Marshal(query)
	return queryBytes
}
