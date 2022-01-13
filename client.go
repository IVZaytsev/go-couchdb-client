package couchdb

type ICDBClient interface {
}

type cdbClient struct {
}

func Init(cfgPath string) (ICDBClient, error) {
	

	return &cdbClient{}, nil
}
