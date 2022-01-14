package couchdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ICDBClient interface {
	GetAllHeaders() []string
}

type cdbClient struct {
	Config *cdbConfig
	DBUrl  string
}

func Init(cfgPath string) (ICDBClient, error) {
	cfg, err := readConfig(cfgPath)
	if err != nil {
		return nil, err
	}

	client := cdbClient{
		Config: cfg,
		DBUrl:  fmt.Sprintf("http://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Hostname, cfg.Port, cfg.Database),
	}

	if err := client.connCheck(); err != nil {
		return nil, err
	}

	return &client, nil
}

func readConfig(cfgPath string) (*cdbConfig, error) {
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

	config := cdbConfig{
		Hostname: "http://127.0.0.1/",
		Port:     5984,
		Username: "admin",
		Password: "adminpw",
		Database: "_users",
	}

	if err = json.Unmarshal(byteValue, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *cdbClient) GetAllHeaders() []string {
	query := cdbQuery{}
	query.SetSelector(map[string]interface{}{"docType": "needsPlanHeader"})
	query.SetFields("")

	rb, err := c.post("_find", query.ToBytes())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(rb))

	return nil
}

func (c *cdbClient) connCheck() error {
	rb, err := c.get("")
	if err != nil {
		return err
	}

	dbInfo := cdbInfo{}
	err = json.Unmarshal(rb, &dbInfo)
	if err != nil {
		return err
	}

	return nil
}

func (c *cdbClient) get(rest string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.DBUrl, rest)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("%d %s: %s", resp.StatusCode, http.StatusText(resp.StatusCode), string(body))
		return nil, err
	}

	return body, nil
}

func (c *cdbClient) post(rest string, body []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.DBUrl, rest)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("%d %s: %s", resp.StatusCode, http.StatusText(resp.StatusCode), string(body))
		return nil, err
	}

	return body, nil
}
