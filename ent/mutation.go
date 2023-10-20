// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ugent-library/projects/ent/predicate"
	"github.com/ugent-library/projects/ent/project"
	"github.com/ugent-library/projects/ent/schema"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeProject = "Project"
)

// ProjectMutation represents an operation that mutates the Project nodes in the graph.
type ProjectMutation struct {
	config
	op                Op
	typ               string
	id                *string
	gismo_id          *string
	identifier        *schema.Identifier
	name              *schema.TranslatedString
	description       *schema.TranslatedString
	founding_date     *string
	dissolution_date  *string
	acronym           *string
	grant_id          *string
	funding_programme *string
	deleted           *bool
	created           *time.Time
	modified          *time.Time
	ts                *string
	clearedFields     map[string]struct{}
	done              bool
	oldValue          func(context.Context) (*Project, error)
	predicates        []predicate.Project
}

var _ ent.Mutation = (*ProjectMutation)(nil)

// projectOption allows management of the mutation configuration using functional options.
type projectOption func(*ProjectMutation)

// newProjectMutation creates new mutation for the Project entity.
func newProjectMutation(c config, op Op, opts ...projectOption) *ProjectMutation {
	m := &ProjectMutation{
		config:        c,
		op:            op,
		typ:           TypeProject,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withProjectID sets the ID field of the mutation.
func withProjectID(id string) projectOption {
	return func(m *ProjectMutation) {
		var (
			err   error
			once  sync.Once
			value *Project
		)
		m.oldValue = func(ctx context.Context) (*Project, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Project.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withProject sets the old Project of the mutation.
func withProject(node *Project) projectOption {
	return func(m *ProjectMutation) {
		m.oldValue = func(context.Context) (*Project, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ProjectMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ProjectMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Project entities.
func (m *ProjectMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ProjectMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ProjectMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Project.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetGismoID sets the "gismo_id" field.
func (m *ProjectMutation) SetGismoID(s string) {
	m.gismo_id = &s
}

// GismoID returns the value of the "gismo_id" field in the mutation.
func (m *ProjectMutation) GismoID() (r string, exists bool) {
	v := m.gismo_id
	if v == nil {
		return
	}
	return *v, true
}

// OldGismoID returns the old "gismo_id" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldGismoID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGismoID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGismoID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGismoID: %w", err)
	}
	return oldValue.GismoID, nil
}

// ResetGismoID resets all changes to the "gismo_id" field.
func (m *ProjectMutation) ResetGismoID() {
	m.gismo_id = nil
}

// SetIdentifier sets the "identifier" field.
func (m *ProjectMutation) SetIdentifier(s schema.Identifier) {
	m.identifier = &s
}

// Identifier returns the value of the "identifier" field in the mutation.
func (m *ProjectMutation) Identifier() (r schema.Identifier, exists bool) {
	v := m.identifier
	if v == nil {
		return
	}
	return *v, true
}

// OldIdentifier returns the old "identifier" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldIdentifier(ctx context.Context) (v schema.Identifier, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIdentifier is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIdentifier requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIdentifier: %w", err)
	}
	return oldValue.Identifier, nil
}

// ResetIdentifier resets all changes to the "identifier" field.
func (m *ProjectMutation) ResetIdentifier() {
	m.identifier = nil
}

// SetName sets the "name" field.
func (m *ProjectMutation) SetName(ss schema.TranslatedString) {
	m.name = &ss
}

// Name returns the value of the "name" field in the mutation.
func (m *ProjectMutation) Name() (r schema.TranslatedString, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldName(ctx context.Context) (v schema.TranslatedString, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *ProjectMutation) ResetName() {
	m.name = nil
}

// SetDescription sets the "description" field.
func (m *ProjectMutation) SetDescription(ss schema.TranslatedString) {
	m.description = &ss
}

// Description returns the value of the "description" field in the mutation.
func (m *ProjectMutation) Description() (r schema.TranslatedString, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldDescription(ctx context.Context) (v schema.TranslatedString, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *ProjectMutation) ResetDescription() {
	m.description = nil
}

// SetFoundingDate sets the "founding_date" field.
func (m *ProjectMutation) SetFoundingDate(s string) {
	m.founding_date = &s
}

// FoundingDate returns the value of the "founding_date" field in the mutation.
func (m *ProjectMutation) FoundingDate() (r string, exists bool) {
	v := m.founding_date
	if v == nil {
		return
	}
	return *v, true
}

// OldFoundingDate returns the old "founding_date" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldFoundingDate(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFoundingDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFoundingDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFoundingDate: %w", err)
	}
	return oldValue.FoundingDate, nil
}

// ClearFoundingDate clears the value of the "founding_date" field.
func (m *ProjectMutation) ClearFoundingDate() {
	m.founding_date = nil
	m.clearedFields[project.FieldFoundingDate] = struct{}{}
}

// FoundingDateCleared returns if the "founding_date" field was cleared in this mutation.
func (m *ProjectMutation) FoundingDateCleared() bool {
	_, ok := m.clearedFields[project.FieldFoundingDate]
	return ok
}

// ResetFoundingDate resets all changes to the "founding_date" field.
func (m *ProjectMutation) ResetFoundingDate() {
	m.founding_date = nil
	delete(m.clearedFields, project.FieldFoundingDate)
}

// SetDissolutionDate sets the "dissolution_date" field.
func (m *ProjectMutation) SetDissolutionDate(s string) {
	m.dissolution_date = &s
}

// DissolutionDate returns the value of the "dissolution_date" field in the mutation.
func (m *ProjectMutation) DissolutionDate() (r string, exists bool) {
	v := m.dissolution_date
	if v == nil {
		return
	}
	return *v, true
}

// OldDissolutionDate returns the old "dissolution_date" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldDissolutionDate(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDissolutionDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDissolutionDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDissolutionDate: %w", err)
	}
	return oldValue.DissolutionDate, nil
}

// ClearDissolutionDate clears the value of the "dissolution_date" field.
func (m *ProjectMutation) ClearDissolutionDate() {
	m.dissolution_date = nil
	m.clearedFields[project.FieldDissolutionDate] = struct{}{}
}

// DissolutionDateCleared returns if the "dissolution_date" field was cleared in this mutation.
func (m *ProjectMutation) DissolutionDateCleared() bool {
	_, ok := m.clearedFields[project.FieldDissolutionDate]
	return ok
}

// ResetDissolutionDate resets all changes to the "dissolution_date" field.
func (m *ProjectMutation) ResetDissolutionDate() {
	m.dissolution_date = nil
	delete(m.clearedFields, project.FieldDissolutionDate)
}

// SetAcronym sets the "acronym" field.
func (m *ProjectMutation) SetAcronym(s string) {
	m.acronym = &s
}

// Acronym returns the value of the "acronym" field in the mutation.
func (m *ProjectMutation) Acronym() (r string, exists bool) {
	v := m.acronym
	if v == nil {
		return
	}
	return *v, true
}

// OldAcronym returns the old "acronym" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldAcronym(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAcronym is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAcronym requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAcronym: %w", err)
	}
	return oldValue.Acronym, nil
}

// ClearAcronym clears the value of the "acronym" field.
func (m *ProjectMutation) ClearAcronym() {
	m.acronym = nil
	m.clearedFields[project.FieldAcronym] = struct{}{}
}

// AcronymCleared returns if the "acronym" field was cleared in this mutation.
func (m *ProjectMutation) AcronymCleared() bool {
	_, ok := m.clearedFields[project.FieldAcronym]
	return ok
}

// ResetAcronym resets all changes to the "acronym" field.
func (m *ProjectMutation) ResetAcronym() {
	m.acronym = nil
	delete(m.clearedFields, project.FieldAcronym)
}

// SetGrantID sets the "grant_id" field.
func (m *ProjectMutation) SetGrantID(s string) {
	m.grant_id = &s
}

// GrantID returns the value of the "grant_id" field in the mutation.
func (m *ProjectMutation) GrantID() (r string, exists bool) {
	v := m.grant_id
	if v == nil {
		return
	}
	return *v, true
}

// OldGrantID returns the old "grant_id" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldGrantID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGrantID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGrantID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGrantID: %w", err)
	}
	return oldValue.GrantID, nil
}

// ClearGrantID clears the value of the "grant_id" field.
func (m *ProjectMutation) ClearGrantID() {
	m.grant_id = nil
	m.clearedFields[project.FieldGrantID] = struct{}{}
}

// GrantIDCleared returns if the "grant_id" field was cleared in this mutation.
func (m *ProjectMutation) GrantIDCleared() bool {
	_, ok := m.clearedFields[project.FieldGrantID]
	return ok
}

// ResetGrantID resets all changes to the "grant_id" field.
func (m *ProjectMutation) ResetGrantID() {
	m.grant_id = nil
	delete(m.clearedFields, project.FieldGrantID)
}

// SetFundingProgramme sets the "funding_programme" field.
func (m *ProjectMutation) SetFundingProgramme(s string) {
	m.funding_programme = &s
}

// FundingProgramme returns the value of the "funding_programme" field in the mutation.
func (m *ProjectMutation) FundingProgramme() (r string, exists bool) {
	v := m.funding_programme
	if v == nil {
		return
	}
	return *v, true
}

// OldFundingProgramme returns the old "funding_programme" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldFundingProgramme(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFundingProgramme is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFundingProgramme requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFundingProgramme: %w", err)
	}
	return oldValue.FundingProgramme, nil
}

// ClearFundingProgramme clears the value of the "funding_programme" field.
func (m *ProjectMutation) ClearFundingProgramme() {
	m.funding_programme = nil
	m.clearedFields[project.FieldFundingProgramme] = struct{}{}
}

// FundingProgrammeCleared returns if the "funding_programme" field was cleared in this mutation.
func (m *ProjectMutation) FundingProgrammeCleared() bool {
	_, ok := m.clearedFields[project.FieldFundingProgramme]
	return ok
}

// ResetFundingProgramme resets all changes to the "funding_programme" field.
func (m *ProjectMutation) ResetFundingProgramme() {
	m.funding_programme = nil
	delete(m.clearedFields, project.FieldFundingProgramme)
}

// SetDeleted sets the "deleted" field.
func (m *ProjectMutation) SetDeleted(b bool) {
	m.deleted = &b
}

// Deleted returns the value of the "deleted" field in the mutation.
func (m *ProjectMutation) Deleted() (r bool, exists bool) {
	v := m.deleted
	if v == nil {
		return
	}
	return *v, true
}

// OldDeleted returns the old "deleted" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldDeleted(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeleted is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeleted requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeleted: %w", err)
	}
	return oldValue.Deleted, nil
}

// ResetDeleted resets all changes to the "deleted" field.
func (m *ProjectMutation) ResetDeleted() {
	m.deleted = nil
}

// SetCreated sets the "created" field.
func (m *ProjectMutation) SetCreated(t time.Time) {
	m.created = &t
}

// Created returns the value of the "created" field in the mutation.
func (m *ProjectMutation) Created() (r time.Time, exists bool) {
	v := m.created
	if v == nil {
		return
	}
	return *v, true
}

// OldCreated returns the old "created" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldCreated(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreated is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreated requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreated: %w", err)
	}
	return oldValue.Created, nil
}

// ResetCreated resets all changes to the "created" field.
func (m *ProjectMutation) ResetCreated() {
	m.created = nil
}

// SetModified sets the "modified" field.
func (m *ProjectMutation) SetModified(t time.Time) {
	m.modified = &t
}

// Modified returns the value of the "modified" field in the mutation.
func (m *ProjectMutation) Modified() (r time.Time, exists bool) {
	v := m.modified
	if v == nil {
		return
	}
	return *v, true
}

// OldModified returns the old "modified" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldModified(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldModified is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldModified requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldModified: %w", err)
	}
	return oldValue.Modified, nil
}

// ResetModified resets all changes to the "modified" field.
func (m *ProjectMutation) ResetModified() {
	m.modified = nil
}

// SetTs sets the "ts" field.
func (m *ProjectMutation) SetTs(s string) {
	m.ts = &s
}

// Ts returns the value of the "ts" field in the mutation.
func (m *ProjectMutation) Ts() (r string, exists bool) {
	v := m.ts
	if v == nil {
		return
	}
	return *v, true
}

// OldTs returns the old "ts" field's value of the Project entity.
// If the Project object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProjectMutation) OldTs(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTs is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTs requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTs: %w", err)
	}
	return oldValue.Ts, nil
}

// ClearTs clears the value of the "ts" field.
func (m *ProjectMutation) ClearTs() {
	m.ts = nil
	m.clearedFields[project.FieldTs] = struct{}{}
}

// TsCleared returns if the "ts" field was cleared in this mutation.
func (m *ProjectMutation) TsCleared() bool {
	_, ok := m.clearedFields[project.FieldTs]
	return ok
}

// ResetTs resets all changes to the "ts" field.
func (m *ProjectMutation) ResetTs() {
	m.ts = nil
	delete(m.clearedFields, project.FieldTs)
}

// Where appends a list predicates to the ProjectMutation builder.
func (m *ProjectMutation) Where(ps ...predicate.Project) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ProjectMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ProjectMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Project, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ProjectMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ProjectMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Project).
func (m *ProjectMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ProjectMutation) Fields() []string {
	fields := make([]string, 0, 13)
	if m.gismo_id != nil {
		fields = append(fields, project.FieldGismoID)
	}
	if m.identifier != nil {
		fields = append(fields, project.FieldIdentifier)
	}
	if m.name != nil {
		fields = append(fields, project.FieldName)
	}
	if m.description != nil {
		fields = append(fields, project.FieldDescription)
	}
	if m.founding_date != nil {
		fields = append(fields, project.FieldFoundingDate)
	}
	if m.dissolution_date != nil {
		fields = append(fields, project.FieldDissolutionDate)
	}
	if m.acronym != nil {
		fields = append(fields, project.FieldAcronym)
	}
	if m.grant_id != nil {
		fields = append(fields, project.FieldGrantID)
	}
	if m.funding_programme != nil {
		fields = append(fields, project.FieldFundingProgramme)
	}
	if m.deleted != nil {
		fields = append(fields, project.FieldDeleted)
	}
	if m.created != nil {
		fields = append(fields, project.FieldCreated)
	}
	if m.modified != nil {
		fields = append(fields, project.FieldModified)
	}
	if m.ts != nil {
		fields = append(fields, project.FieldTs)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ProjectMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case project.FieldGismoID:
		return m.GismoID()
	case project.FieldIdentifier:
		return m.Identifier()
	case project.FieldName:
		return m.Name()
	case project.FieldDescription:
		return m.Description()
	case project.FieldFoundingDate:
		return m.FoundingDate()
	case project.FieldDissolutionDate:
		return m.DissolutionDate()
	case project.FieldAcronym:
		return m.Acronym()
	case project.FieldGrantID:
		return m.GrantID()
	case project.FieldFundingProgramme:
		return m.FundingProgramme()
	case project.FieldDeleted:
		return m.Deleted()
	case project.FieldCreated:
		return m.Created()
	case project.FieldModified:
		return m.Modified()
	case project.FieldTs:
		return m.Ts()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ProjectMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case project.FieldGismoID:
		return m.OldGismoID(ctx)
	case project.FieldIdentifier:
		return m.OldIdentifier(ctx)
	case project.FieldName:
		return m.OldName(ctx)
	case project.FieldDescription:
		return m.OldDescription(ctx)
	case project.FieldFoundingDate:
		return m.OldFoundingDate(ctx)
	case project.FieldDissolutionDate:
		return m.OldDissolutionDate(ctx)
	case project.FieldAcronym:
		return m.OldAcronym(ctx)
	case project.FieldGrantID:
		return m.OldGrantID(ctx)
	case project.FieldFundingProgramme:
		return m.OldFundingProgramme(ctx)
	case project.FieldDeleted:
		return m.OldDeleted(ctx)
	case project.FieldCreated:
		return m.OldCreated(ctx)
	case project.FieldModified:
		return m.OldModified(ctx)
	case project.FieldTs:
		return m.OldTs(ctx)
	}
	return nil, fmt.Errorf("unknown Project field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProjectMutation) SetField(name string, value ent.Value) error {
	switch name {
	case project.FieldGismoID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGismoID(v)
		return nil
	case project.FieldIdentifier:
		v, ok := value.(schema.Identifier)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIdentifier(v)
		return nil
	case project.FieldName:
		v, ok := value.(schema.TranslatedString)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case project.FieldDescription:
		v, ok := value.(schema.TranslatedString)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case project.FieldFoundingDate:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFoundingDate(v)
		return nil
	case project.FieldDissolutionDate:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDissolutionDate(v)
		return nil
	case project.FieldAcronym:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAcronym(v)
		return nil
	case project.FieldGrantID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGrantID(v)
		return nil
	case project.FieldFundingProgramme:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFundingProgramme(v)
		return nil
	case project.FieldDeleted:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeleted(v)
		return nil
	case project.FieldCreated:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreated(v)
		return nil
	case project.FieldModified:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetModified(v)
		return nil
	case project.FieldTs:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTs(v)
		return nil
	}
	return fmt.Errorf("unknown Project field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ProjectMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ProjectMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProjectMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Project numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ProjectMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(project.FieldFoundingDate) {
		fields = append(fields, project.FieldFoundingDate)
	}
	if m.FieldCleared(project.FieldDissolutionDate) {
		fields = append(fields, project.FieldDissolutionDate)
	}
	if m.FieldCleared(project.FieldAcronym) {
		fields = append(fields, project.FieldAcronym)
	}
	if m.FieldCleared(project.FieldGrantID) {
		fields = append(fields, project.FieldGrantID)
	}
	if m.FieldCleared(project.FieldFundingProgramme) {
		fields = append(fields, project.FieldFundingProgramme)
	}
	if m.FieldCleared(project.FieldTs) {
		fields = append(fields, project.FieldTs)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ProjectMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ProjectMutation) ClearField(name string) error {
	switch name {
	case project.FieldFoundingDate:
		m.ClearFoundingDate()
		return nil
	case project.FieldDissolutionDate:
		m.ClearDissolutionDate()
		return nil
	case project.FieldAcronym:
		m.ClearAcronym()
		return nil
	case project.FieldGrantID:
		m.ClearGrantID()
		return nil
	case project.FieldFundingProgramme:
		m.ClearFundingProgramme()
		return nil
	case project.FieldTs:
		m.ClearTs()
		return nil
	}
	return fmt.Errorf("unknown Project nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ProjectMutation) ResetField(name string) error {
	switch name {
	case project.FieldGismoID:
		m.ResetGismoID()
		return nil
	case project.FieldIdentifier:
		m.ResetIdentifier()
		return nil
	case project.FieldName:
		m.ResetName()
		return nil
	case project.FieldDescription:
		m.ResetDescription()
		return nil
	case project.FieldFoundingDate:
		m.ResetFoundingDate()
		return nil
	case project.FieldDissolutionDate:
		m.ResetDissolutionDate()
		return nil
	case project.FieldAcronym:
		m.ResetAcronym()
		return nil
	case project.FieldGrantID:
		m.ResetGrantID()
		return nil
	case project.FieldFundingProgramme:
		m.ResetFundingProgramme()
		return nil
	case project.FieldDeleted:
		m.ResetDeleted()
		return nil
	case project.FieldCreated:
		m.ResetCreated()
		return nil
	case project.FieldModified:
		m.ResetModified()
		return nil
	case project.FieldTs:
		m.ResetTs()
		return nil
	}
	return fmt.Errorf("unknown Project field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ProjectMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ProjectMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ProjectMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ProjectMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ProjectMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ProjectMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ProjectMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Project unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ProjectMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Project edge %s", name)
}
