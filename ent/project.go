// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ugent-library/projects/ent/project"
	"github.com/ugent-library/projects/ent/schema"
)

// Project is the model entity for the Project schema.
type Project struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Identifier holds the value of the "identifier" field.
	Identifier schema.Identifier `json:"identifier,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// FoundingDate holds the value of the "founding_date" field.
	FoundingDate string `json:"founding_date,omitempty"`
	// DissolutionDate holds the value of the "dissolution_date" field.
	DissolutionDate string `json:"dissolution_date,omitempty"`
	// Acronym holds the value of the "acronym" field.
	Acronym string `json:"acronym,omitempty"`
	// GrantID holds the value of the "grant_id" field.
	GrantID string `json:"grant_id,omitempty"`
	// FundingProgramme holds the value of the "funding_programme" field.
	FundingProgramme string `json:"funding_programme,omitempty"`
	// Created holds the value of the "created" field.
	Created time.Time `json:"created,omitempty"`
	// Modified holds the value of the "modified" field.
	Modified time.Time `json:"modified,omitempty"`
	// Ts holds the value of the "ts" field.
	Ts           string `json:"ts,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Project) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case project.FieldIdentifier:
			values[i] = new([]byte)
		case project.FieldID, project.FieldName, project.FieldDescription, project.FieldFoundingDate, project.FieldDissolutionDate, project.FieldAcronym, project.FieldGrantID, project.FieldFundingProgramme, project.FieldTs:
			values[i] = new(sql.NullString)
		case project.FieldCreated, project.FieldModified:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Project fields.
func (pr *Project) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case project.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pr.ID = value.String
			}
		case project.FieldIdentifier:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field identifier", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Identifier); err != nil {
					return fmt.Errorf("unmarshal field identifier: %w", err)
				}
			}
		case project.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case project.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case project.FieldFoundingDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field founding_date", values[i])
			} else if value.Valid {
				pr.FoundingDate = value.String
			}
		case project.FieldDissolutionDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dissolution_date", values[i])
			} else if value.Valid {
				pr.DissolutionDate = value.String
			}
		case project.FieldAcronym:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field acronym", values[i])
			} else if value.Valid {
				pr.Acronym = value.String
			}
		case project.FieldGrantID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field grant_id", values[i])
			} else if value.Valid {
				pr.GrantID = value.String
			}
		case project.FieldFundingProgramme:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field funding_programme", values[i])
			} else if value.Valid {
				pr.FundingProgramme = value.String
			}
		case project.FieldCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				pr.Created = value.Time
			}
		case project.FieldModified:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified", values[i])
			} else if value.Valid {
				pr.Modified = value.Time
			}
		case project.FieldTs:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ts", values[i])
			} else if value.Valid {
				pr.Ts = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Project.
// This includes values selected through modifiers, order, etc.
func (pr *Project) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// Update returns a builder for updating this Project.
// Note that you need to call Project.Unwrap() before calling this method if this Project
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Project) Update() *ProjectUpdateOne {
	return NewProjectClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Project entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Project) Unwrap() *Project {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Project is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Project) String() string {
	var builder strings.Builder
	builder.WriteString("Project(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("identifier=")
	builder.WriteString(fmt.Sprintf("%v", pr.Identifier))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", ")
	builder.WriteString("founding_date=")
	builder.WriteString(pr.FoundingDate)
	builder.WriteString(", ")
	builder.WriteString("dissolution_date=")
	builder.WriteString(pr.DissolutionDate)
	builder.WriteString(", ")
	builder.WriteString("acronym=")
	builder.WriteString(pr.Acronym)
	builder.WriteString(", ")
	builder.WriteString("grant_id=")
	builder.WriteString(pr.GrantID)
	builder.WriteString(", ")
	builder.WriteString("funding_programme=")
	builder.WriteString(pr.FundingProgramme)
	builder.WriteString(", ")
	builder.WriteString("created=")
	builder.WriteString(pr.Created.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified=")
	builder.WriteString(pr.Modified.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ts=")
	builder.WriteString(pr.Ts)
	builder.WriteByte(')')
	return builder.String()
}

// Projects is a parsable slice of Project.
type Projects []*Project
