package gismo

type Identifier struct {
	Type       string `json:"type"`
	PropertyID string `json:"property_id"`
	Value      string `json:"value"`
}

type Project struct {
	Type            string        `json:"type"`
	Identifier      []*Identifier `json:"identifier"`
	Name            string        `json:"name,omitempty"`
	Description     string        `json:"description,omitempty"`
	FoundingDate    string        `json:"founding_date,omitempty"`
	DissolutionDate string        `json:"dissolution_date,omitempty"`
}
