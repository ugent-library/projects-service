package es6

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
)

type M map[string]any

type responseBody struct {
	Hits struct {
		Total int `json:"total"`
		Hits  []struct {
			ID     string          `json:"_id"`
			Source json.RawMessage `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type Config struct {
	Conn  []string
	Index string
}

type search struct {
	client *elasticsearch.Client
	index  string
}

func NewSearcher(c Config) (*search, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: c.Conn,
	})

	if err != nil {
		return nil, fmt.Errorf("elastic search: can't connect with server: %w", err)
	}

	return &search{
		client: client,
		index:  c.Index,
	}, nil
}

func (s *search) search(requestBody M, responseBody any) error {
	sorts := make([]string, 0)

	if v, exists := requestBody["sort"]; exists {
		sorts = append(sorts, v.([]string)...)
		delete(requestBody, "sort")
	}

	opts := []func(*esapi.SearchRequest){
		s.client.Search.WithContext(context.Background()),
		s.client.Search.WithIndex(s.index),
		s.client.Search.WithTrackTotalHits(true),
		s.client.Search.WithSort(sorts...),
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(requestBody); err != nil {
		return err
	}
	opts = append(opts, s.client.Search.WithBody(&buf))

	res, err := s.client.Search(opts...)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		buf := &bytes.Buffer{}
		if _, err := io.Copy(buf, res.Body); err != nil {
			return err
		}
		return errors.New("elastic search: error response: " + buf.String())
	}

	if err := json.NewDecoder(res.Body).Decode(responseBody); err != nil {
		return fmt.Errorf("elastic search: error parsing the response body: %w", err)
	}

	return nil
}

func (s *search) SuggestProjects(q string) (map[string]json.RawMessage, error) {
	var boosts = map[string]string{
		"iweto_id":               "100",
		"gismo_id":               "80",
		"eu_call_id":             "80",
		"cordis_id":              "80",
		"eu_framework_programme": "80",
		"acronym":                "70",
		"all":                    "0.1",
		"phrase_ngram":           "0.05",
		"ngram":                  "0.01",
	}

	limit := 20

	var query M = M{
		"match_all": M{},
	}

	q = strings.TrimSpace(q)

	if q != "" {
		dismaxQueries := make([]M, 0, len(boosts))
		for field, boost := range boosts {
			dismaxQuery := M{
				"match": M{
					field: M{
						"query":    q,
						"operator": "AND",
						"boost":    boost,
					},
				},
			}
			dismaxQueries = append(dismaxQueries, dismaxQuery)
		}
		query = M{
			"dis_max": M{
				"queries": dismaxQueries,
			},
		}
	}

	requestBody := M{
		"query": query,
		"size":  limit,
		"sort":  []string{"_score:desc"},
	}

	responseBody := &responseBody{}

	if err := s.search(requestBody, responseBody); err != nil {
		return nil, err
	}

	res := make(map[string]json.RawMessage)
	for _, hit := range responseBody.Hits.Hits {
		res[hit.ID] = hit.Source
	}

	return res, nil
}
