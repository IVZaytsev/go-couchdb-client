package couchdb

type cdbConfig struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"post"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type SortType string

type cdbQuery struct {
	Selector       interface{}   `json:"selector,omitempty"`
	Fields         interface{}   `json:"fields,omitempty"`
	Sort           []interface{} `json:"sort,omitempty"`
	Limit          interface{}   `json:"limit,omitempty"`
	Skip           interface{}   `json:"skip,omitempty"`
	UseIndex       interface{}   `json:"use_index,omitempty"`
	ExecutionStats bool          `json:"execution_stats,omitempty"`
}

type cdbInfo struct {
	DBName   string `json:"db_name"`
	DocCount int    `json:"doc_count"`
}
