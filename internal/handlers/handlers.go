package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"search/internal/es"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/gorilla/mux"
)

func UsersSearch(w http.ResponseWriter, r *http.Request) {
	queryParam := mux.Vars(r)["query"]

	// Normalizing the query parameter to lowercase and split into individual names
	names := strings.Split(strings.ToLower(queryParam), " ")

	// Create the query JSON structure with the normalized names
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"nested": map[string]interface{}{
				"path": "users",
				"query": map[string]interface{}{
					"bool": map[string]interface{}{
						"should": createTermQueries("users.name", names),
					},
				},
			},
		},
	}

	// Convert the query to JSON bytes
	queryBytes, err := json.Marshal(query)
	if err != nil {
		http.Error(w, "Error marshaling query to JSON", http.StatusInternalServerError)
		return
	}

	// Create the Elasticsearch request
	req := esapi.SearchRequest{
		Index: []string{"projects"},
		Body:  strings.NewReader(string(queryBytes)),
	}

	// Perform the Elasticsearch search
	res, err := req.Do(r.Context(), es.Client)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Elasticsearch search error", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Copy the Elasticsearch response to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	_, _ = io.Copy(w, res.Body)
}

func HashtagsSearch(w http.ResponseWriter, r *http.Request) {
	queryParam := mux.Vars(r)["query"]

	// Normalizing the query parameter to lowercase and split into individual hashtags
	hashtags := strings.Split(strings.ToLower(queryParam), " ")

	// Create the query JSON structure with the normalized hashtags
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"nested": map[string]interface{}{
				"path": "hashtags",
				"query": map[string]interface{}{
					"bool": map[string]interface{}{
						"should": createTermQueries("hashtags.name", hashtags),
					},
				},
			},
		},
	}

	// Convert the query to JSON bytes
	queryBytes, err := json.Marshal(query)
	if err != nil {
		http.Error(w, "Error marshaling query to JSON", http.StatusInternalServerError)
		return
	}

	// Create the Elasticsearch request
	req := esapi.SearchRequest{
		Index: []string{"projects"},
		Body:  strings.NewReader(string(queryBytes)),
	}

	// Perform the Elasticsearch search
	res, err := req.Do(r.Context(), es.Client)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Elasticsearch search error", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Copy the Elasticsearch response to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	_, _ = io.Copy(w, res.Body)
}

func FuzzySearch(w http.ResponseWriter, r *http.Request) {
	queryParam := mux.Vars(r)["query"]

	// Create the fuzzy query JSON structure
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":     queryParam,
				"fields":    []string{"slug", "description"},
				"fuzziness": "AUTO",
			},
		},
	}

	// Convert the query to JSON bytes
	queryBytes, err := json.Marshal(query)
	if err != nil {
		http.Error(w, "Error marshaling query to JSON", http.StatusInternalServerError)
		return
	}

	// Create the Elasticsearch request
	req := esapi.SearchRequest{
		Index: []string{"projects"},
		Body:  strings.NewReader(string(queryBytes)),
	}

	// Perform the Elasticsearch search
	res, err := req.Do(r.Context(), es.Client)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Elasticsearch search error", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Copy the Elasticsearch response to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	_, _ = io.Copy(w, res.Body)
}

// Helper function to create an array of term queries for the given field and values
func createTermQueries(field string, values []string) []interface{} {
	queries := make([]interface{}, len(values))
	for i, value := range values {
		queries[i] = map[string]interface{}{
			"term": map[string]interface{}{
				field: value,
			},
		}
	}
	return queries
}
