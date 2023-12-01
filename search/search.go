package search

import (
	"encoding/json"
)

type Searcher interface {
	SuggestProjects(q string) (map[string]json.RawMessage, error)
}
