package es

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
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
		return
	}
	fmt.Println("Elasticsearch connection successful.")
	Client = client
}

func CreateProjectsIndex(index string) error {
	req := esapi.IndicesExistsRequest{
		Index: []string{index},
	}

	ctx := context.Background() // Create a context

	res, err := req.Do(ctx, Client) // Provide the context and client to the Do method
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		if res.StatusCode == http.StatusNotFound {
			// Index doesn't exist, so create it
			mapping := `{
				"mappings": {
					"properties": {
						"id": { "type": "integer" },
						"name": { "type": "text" },
						"slug": { "type": "text" },
						"description": { "type": "text" },
						"users": {
							"type": "nested",
							"properties": {
								"id": { "type": "integer" },
								"name": { "type": "text" },
								"created_at": { "type": "date" }
							}
						},
						"hashtags": {
							"type": "nested",
							"properties": {
								"id": { "type": "integer" },
								"name": { "type": "text" },
								"created_at": { "type": "date" }
							}
						}
					}
				}
			}`

			createReq := esapi.IndicesCreateRequest{
				Index: index,
				Body:  strings.NewReader(mapping),
			}

			createRes, err := createReq.Do(ctx, Client) // Provide the context and client to the Do method
			if err != nil {
				return err
			}
			defer createRes.Body.Close()

			if createRes.IsError() {
				return fmt.Errorf("failed to create index: %s", createRes.String())
			}

			fmt.Println("Index '" + index + "' created successfully.")
		} else {
			return fmt.Errorf("index check error: %s", res.String())
		}
	} else {
		fmt.Println("Index '" + index + "' already exists.")
	}

	return nil
}
