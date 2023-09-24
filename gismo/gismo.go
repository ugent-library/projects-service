package gismo

type Identifier struct {
	Type       string `json:"type"`
	PropertyID string `json:"propertyID"`
	Value      string `json:"value"`
}

type Project struct {
	Type            string        `json:"type"`
	Identifier      []*Identifier `json:"identifier"`
	Name            string        `json:"name,omitempty"`
	Description     string        `json:"description,omitempty"`
	FoundingDate    string        `json:"foundingDate,omitempty"`
	DissolutionDate string        `json:"dissolutionDate,omitempty"`
}
