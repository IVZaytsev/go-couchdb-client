package couchdb

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ICDBClient interface {
}

type cdbClient struct {
	Config *CouchDBConfig
}

func Init(cfgPath string) (ICDBClient, error) {
	cfg, err := readConfig(cfgPath)
	if err != nil {
		return nil, err
	}
	return &cdbClient{Config: cfg}, nil
}

func readConfig(cfgPath string) (*CouchDBConfig, error) {
	if _, err := os.Stat(cfgPath); err != nil {
		return nil, err
	}

	jsonFile, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	config := CouchDBConfig{
		Hostname: "http://127.0.0.1/",
		Port:     5984,
		Username: "admin",
		Password: "adminpw",
	}

	if err = json.Unmarshal(byteValue, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
