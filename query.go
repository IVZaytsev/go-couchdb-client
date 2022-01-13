package couchdb

import "encoding/json"

type ICDBQuery interface {
	SetSelector(selector map[string]interface{}) *CouchDBQuery
	AddSelector(selector map[string]interface{}) *CouchDBQuery
	SetFields(fields []string) *CouchDBQuery
	SetSort(sortField string, st SortType) *CouchDBQuery
	SetLimit(limit int) *CouchDBQuery
	SetSkip(skip int) *CouchDBQuery
	SetIndex(index string) *CouchDBQuery
	ToString() string
}

func (query *CouchDBQuery) SetSelector(selector map[string]interface{}) *CouchDBQuery {
	query.Selector = selector
	return query
}

func (query *CouchDBQuery) AddSelector(selector map[string]interface{}) *CouchDBQuery {
	for key, value := range selector {
		query.Selector[key] = value
	}
	return query
}

func (query *CouchDBQuery) SetFields(fields []string) *CouchDBQuery {
	query.Fields = fields
	return query
}

func (query *CouchDBQuery) SetSort(sortField string, st SortType) *CouchDBQuery {
	query.Sort = append(query.Sort, map[string]string{sortField: string(st)})
	return query
}

func (query *CouchDBQuery) SetLimit(limit int) *CouchDBQuery {
	query.Limit = limit
	return query
}

func (query *CouchDBQuery) SetSkip(skip int) *CouchDBQuery {
	query.Skip = skip
	return query
}

func (query *CouchDBQuery) SetIndex(index string) *CouchDBQuery {
	query.UseIndex = index
	return query
}

func (query *CouchDBQuery) ToString() string {
	queryString, _ := json.Marshal(query)
	return string(queryString)
}
