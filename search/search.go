package search

import (
	"encoding/json"
)

type Engine interface {
	SuggestProjects(index, q string) (map[string]json.RawMessage, error)
}

type Searcher struct {
	engine Engine
	index  string
}

func NewSearcher(engine Engine, index string) *Searcher {
	return &Searcher{
		engine: engine,
		index:  index,
	}
}

func (s *Searcher) SuggestProjects(q string) ([]string, error) {
	hits, err := s.engine.SuggestProjects(s.index, q)
	if err != nil {
		return nil, err
	}

	// only return the project ids. fetch project from db.
	res := make([]string, 0, len(hits))
	for k := range hits {
		res = append(res, k)
	}

	return res, nil
}
