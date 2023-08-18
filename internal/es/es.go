package es

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/elastic/go-elasticsearch/v7"
)

var Client *elasticsearch.Client

func CreateClient() {
	esURL := os.Getenv("ELASTICSEARCH_URL")
	esUsername := os.Getenv("ELASTICSEARCH_USERNAME")
	esPassword := os.Getenv("ELASTICSEARCH_PASSWORD")
	cfg := elasticsearch.Config{
		Addresses: []string{esURL},
		Username:  esUsername,
		Password:  esPassword,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println("Elasticsearch connection error: ", err)
	}
	fmt.Println("Elasticsearch connection successful.")
	Client = client
}
