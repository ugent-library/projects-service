// Code generated by ent, DO NOT EDIT.

package project

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ugent-library/projects/ent/schema"
)

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIdentifier holds the string denoting the identifier field in the database.
	FieldIdentifier = "identifier"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldFoundingDate holds the string denoting the founding_date field in the database.
	FieldFoundingDate = "founding_date"
	// FieldDissolutionDate holds the string denoting the dissolution_date field in the database.
	FieldDissolutionDate = "dissolution_date"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldModified holds the string denoting the modified field in the database.
	FieldModified = "modified"
	// Table holds the table name of the project in the database.
	Table = "projects"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldIdentifier,
	FieldName,
	FieldDescription,
	FieldFoundingDate,
	FieldDissolutionDate,
	FieldCreated,
	FieldModified,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIdentifier holds the default value on creation for the "identifier" field.
	DefaultIdentifier []schema.Identifier
	// DefaultCreated holds the default value on creation for the "created" field.
	DefaultCreated func() time.Time
	// DefaultModified holds the default value on creation for the "modified" field.
	DefaultModified func() time.Time
	// UpdateDefaultModified holds the default value on update for the "modified" field.
	UpdateDefaultModified func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the Project queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByFoundingDate orders the results by the founding_date field.
func ByFoundingDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFoundingDate, opts...).ToFunc()
}

// ByDissolutionDate orders the results by the dissolution_date field.
func ByDissolutionDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDissolutionDate, opts...).ToFunc()
}

// ByCreated orders the results by the created field.
func ByCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreated, opts...).ToFunc()
}

// ByModified orders the results by the modified field.
func ByModified(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModified, opts...).ToFunc()
}
