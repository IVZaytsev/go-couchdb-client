package couchdb

type CouchDBConfig struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"post"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SortType string

type CouchDBQuery struct {
	Selector map[string]interface{} `json:"selector,omitempty"`
	Fields   interface{}            `json:"fields,omitempty"`
	Sort     []interface{}          `json:"sort,omitempty"`
	Limit    interface{}            `json:"limit,omitempty"`
	Skip     interface{}            `json:"skip,omitempty"`
	UseIndex interface{}            `json:"use_index,omitempty"`
}
