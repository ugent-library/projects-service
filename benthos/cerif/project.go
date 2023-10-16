package cerif

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/antchfx/xmlquery"
)

var ErrNonCompliantXml = errors.New("non compliant cerif xml")

// type MongoObject struct {
// 	DateCreated string `json:"date_created"`
// 	DateUpdated string `json:"date_updated"`
// 	GismoID     string `json:"_id"`
// }

// type GISMOAction struct {
// 	CreatedDate string        `json:"date_created,omitempty"`
// 	SentDate    time.Time     `json:"sent_date,omitempty"`
// 	Action      string        `json:"action,omitempty"`
// 	Project     *CERIFProject `json:"project,omitempty"`
// }

type CERIFProject struct {
	ID             string                           `json:"id,omitempty"`
	StartDate      string                           `json:"start_date,omitempty"`
	EndDate        string                           `json:"end_date,omitempty"`
	Acronym        string                           `json:"acronym,omitempty"`
	Title          []CERIFTranslatedString          `json:"title,omitempty"`
	Abstract       []CERIFTranslatedString          `json:"abstract,omitempty"`
	Keyword        []CERIFTranslatedString          `json:"keyword,omitempty"`
	Classification []CERIFClassification            `json:"classfication,omitempty"`
	FederatedIDS   []CERIFFederatedIDClassification `json:"federated_ids,omitempty"`
}

type CERIFTranslatedString struct {
	Translation string `json:"translation,omitempty"`
	Lang        string `json:"lang,omitempty"`
	Value       string `json:"value,omitempty"`
}

type CERIFClassification struct {
	URI         string                    `json:"uri,omitempty"`
	Name        []CERIFTranslatedString   `json:"name,omitempty"`
	Description []CERIFTranslatedString   `json:"description,omitempty"`
	Terms       []CERIFClassificationTerm `json:"terms,omitempty"`
}

type CERIFClassificationTerm struct {
	URI       string                  `json:"uri,omitempty"`
	Term      []CERIFTranslatedString `json:"term,omitempty"`
	StartDate time.Time               `json:"start_date,omitempty"`
	EndDate   time.Time               `json:"end_date,omitempty"`
}

type CERIFFederatedIDClassification struct {
	URI         string                  `json:"uri,omitempty"`
	Name        []CERIFTranslatedString `json:"name,omitempty"`
	Description []CERIFTranslatedString `json:"description,omitempty"`
	IDS         []CERIFFederatedID      `json:"ids,omitempty"`
}

type CERIFFederatedID struct {
	ID        string                  `json:"id,omitempty"`
	URI       string                  `json:"uri,omitempty"`
	Term      []CERIFTranslatedString `json:"term,omitempty"`
	StartDate time.Time               `json:"start_date,omitempty"`
	EndDate   time.Time               `json:"end_date,omitempty"`
}

// func ParseGismoAction(buf []byte) (*GISMOAction, error) {
// 	gobj := &MongoObject{}
// 	json.Unmarshal(buf, gobj)

// 	doc, err := xmlquery.Parse(bytes.NewReader(buf))
// 	if err != nil {
// 		return nil, err
// 	}

// 	node := xmlquery.FindOne(doc, "//ns0:cfProj")
// 	if node == nil {
// 		return nil, fmt.Errorf("not a valid cerif document")
// 	}

// 	projects := xmlquery.FindOne(doc, "//ns0:projects")
// 	if projects == nil {
// 		return nil, fmt.Errorf("not a valid cerif document")
// 	}

// 	p := CerifParseProject(node, doc)

// 	a := &GISMOAction{
// 		Project: p,
// 	}

// 	if val, err := parseDateTimeField(projects.SelectAttr("sentDateTime")); err == nil {
// 		a.SentDate = val
// 	}

// 	a.Action = node.SelectAttr("action")
// 	a.CreatedDate = gobj.DateCreated

// 	return a, nil
// }

func CerifParseProject(buf []byte) (*CERIFProject, error) {
	doc, err := xmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	node := xmlquery.FindOne(doc, "//ns0:cfProj")
	if node == nil {
		return nil, fmt.Errorf("not a valid cerif document")
	}

	projects := xmlquery.FindOne(doc, "//ns0:projects")
	if projects == nil {
		return nil, fmt.Errorf("not a valid cerif document")
	}

	p := &CERIFProject{}

	for _, n := range node.SelectElements("*") {
		// log.Println(n.Data)
		switch n.Data {
		case "cfProjId":
			p.ID = n.InnerText()
		case "cfStartDate":
			p.StartDate = n.InnerText()
		case "cfEndDate":
			p.EndDate = n.InnerText()
		case "cfAcro":
			p.Acronym = n.InnerText()
		case "cfTitle":
			p.Title = parseTranslatedField(p.Title, n)
		case "cfAbstr":
			p.Abstract = parseTranslatedField(p.Abstract, n)
		case "cfKeyw":
			p.Keyword = parseTranslatedField(p.Keyword, n)
		case "cfProj_Class":
			p.Classification = parseClassification(p.Classification, n, doc)
		case "cfFedId":
			p.FederatedIDS = parseFederatedIDClassification(p.FederatedIDS, n, doc)
		}
	}

	return p, nil
}

func parseDateTimeField(dt string) (time.Time, error) {
	return time.Parse(time.RFC3339, strings.TrimSpace(dt))
}

func parseTranslatedField(f []CERIFTranslatedString, n *xmlquery.Node) []CERIFTranslatedString {
	s := CERIFTranslatedString{
		Translation: n.SelectAttr("cfTrans"),
		Lang:        n.SelectAttr("cfLangCode"),
		Value:       n.InnerText(),
	}

	if f == nil {
		f = []CERIFTranslatedString{s}
	} else {
		f = append(f, s)
	}

	return f
}

func parseClassification(field []CERIFClassification, term *xmlquery.Node, doc *xmlquery.Node) []CERIFClassification {
	// Resolve the classification for the term
	classSchemeId := xmlquery.FindOne(term, "/ns1:cfClassSchemeId").InnerText()
	path := fmt.Sprintf("//ns0:cfClassScheme/ns1:cfClassSchemeId[.='%s']/..", classSchemeId)
	xmlClass := xmlquery.FindOne(doc, path)

	tmp := CERIFClassification{}
	for _, v := range xmlClass.SelectElements("*") {
		switch v.Data {
		case "cfURI":
			tmp.URI = v.InnerText()
		case "cfDescr":
			tmp.Description = parseTranslatedField(tmp.Description, v)
		case "cfName":
			tmp.Name = parseTranslatedField(tmp.Name, v)
		}
	}

	// Fetch if classification already was added to project
	var class *CERIFClassification
	for k, v := range field {
		if v.URI == tmp.URI {
			class = &field[k]
			break
		}
	}

	if class == nil {
		field = append(field, tmp)
		class = &tmp
	}

	// Transform term to a CERIFClassificationTerm
	t := CERIFClassificationTerm{}

	// Resolve the term for the term id
	tid := xmlquery.FindOne(term, "/ns1:cfClassId").InnerText()
	path = fmt.Sprintf("//ns0:cfClass/ns1:cfClassId[.='%s']/..", tid)
	xmlTerm := xmlquery.FindOne(doc, path)

	for _, v := range term.SelectElements("*") {
		switch v.Data {
		case "cfStartDate":
			if val, err := parseDateTimeField(v.InnerText()); err == nil {
				t.StartDate = val
			}
		case "cfEndDate":
			if val, err := parseDateTimeField(v.InnerText()); err == nil {
				t.EndDate = val
			}
		}
	}

	for _, v := range xmlTerm.SelectElements("*") {
		switch v.Data {
		case "cfURI":
			t.URI = v.InnerText()
		case "cfTerm":
			t.Term = parseTranslatedField(t.Term, v)
		}
	}

	// Add or replace CERIFClassificationTerm to / in CERIFClassification
	found := false
	for k, v := range class.Terms {
		if v.URI == t.URI {
			found = true
			class.Terms[k] = t
			break
		}
	}

	if !found {
		class.Terms = append(class.Terms, t)
	}

	// Update the field & return the entire value
	for k, cc := range field {
		if cc.URI == class.URI {
			field[k] = *class
		}
	}

	return field
}

func parseFederatedIDClassification(field []CERIFFederatedIDClassification, term *xmlquery.Node, doc *xmlquery.Node) []CERIFFederatedIDClassification {
	// Resolve the classification for the term
	classSchemeId := xmlquery.FindOne(term, "/ns1:cfClassSchemeId").InnerText()
	path := fmt.Sprintf("//ns0:cfClassScheme/ns1:cfClassSchemeId[.='%s']/..", classSchemeId)
	xmlClass := xmlquery.FindOne(doc, path)

	tmp := CERIFFederatedIDClassification{}
	for _, v := range xmlClass.SelectElements("*") {
		//log.Println(v.Data)
		switch v.Data {
		case "cfURI":
			tmp.URI = v.InnerText()
		case "cfDescr":
			tmp.Description = parseTranslatedField(tmp.Description, v)
		case "cfName":
			tmp.Name = parseTranslatedField(tmp.Name, v)
		}
	}

	// Fetch if classification already was added to project
	var class *CERIFFederatedIDClassification
	for k, v := range field {
		if v.URI == tmp.URI {
			class = &field[k]
			break
		}
	}

	if class == nil {
		field = append(field, tmp)
		class = &tmp
	}

	// Transform term to a CERIFFederatedID
	c := CERIFFederatedID{}

	// Resolve the term for the term id
	fedId := xmlquery.FindOne(term, "/ns1:cfClassId").InnerText()
	path = fmt.Sprintf("//ns0:cfClass/ns1:cfClassId[.='%s']/..", fedId)
	xmlTerm := xmlquery.FindOne(doc, path)

	for _, v := range term.SelectElements("*") {
		// log.Printf("** %s", v.Data)
		switch v.Data {
		case "cfFedId":
			c.ID = v.InnerText()
		case "cfStartDate":
			if val, err := parseDateTimeField(v.InnerText()); err == nil {
				c.StartDate = val
			}
		case "cfEndDate":
			if val, err := parseDateTimeField(v.InnerText()); err == nil {
				c.EndDate = val
			}
		}
	}

	for _, v := range xmlTerm.SelectElements("*") {
		// log.Printf("-- %s", v.Data)
		switch v.Data {
		case "cfURI":
			c.URI = v.InnerText()
		case "cfTerm":
			c.Term = parseTranslatedField(c.Term, v)
		}
	}

	// Add or replace CERIFFederatedID to / in CERIFFederatedIDClassification
	found := false
	for k, v := range class.IDS {
		if v.ID == c.ID && v.URI == c.URI {
			found = true
			class.IDS[k] = c
			break
		}
	}

	if !found {
		class.IDS = append(class.IDS, c)
	}

	// Update the field & return the entire value
	for k, cc := range field {
		if cc.URI == class.URI {
			field[k] = *class
		}
	}

	return field
}
