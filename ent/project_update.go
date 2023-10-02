// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/projects/ent/predicate"
	"github.com/ugent-library/projects/ent/project"
	"github.com/ugent-library/projects/ent/schema"
)

// ProjectUpdate is the builder for updating Project entities.
type ProjectUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectMutation
}

// Where appends a list predicates to the ProjectUpdate builder.
func (pu *ProjectUpdate) Where(ps ...predicate.Project) *ProjectUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetIdentifier sets the "identifier" field.
func (pu *ProjectUpdate) SetIdentifier(s []schema.Identifier) *ProjectUpdate {
	pu.mutation.SetIdentifier(s)
	return pu
}

// AppendIdentifier appends s to the "identifier" field.
func (pu *ProjectUpdate) AppendIdentifier(s []schema.Identifier) *ProjectUpdate {
	pu.mutation.AppendIdentifier(s)
	return pu
}

// SetName sets the "name" field.
func (pu *ProjectUpdate) SetName(s string) *ProjectUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableName(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// ClearName clears the value of the "name" field.
func (pu *ProjectUpdate) ClearName() *ProjectUpdate {
	pu.mutation.ClearName()
	return pu
}

// SetDescription sets the "description" field.
func (pu *ProjectUpdate) SetDescription(s string) *ProjectUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableDescription(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *ProjectUpdate) ClearDescription() *ProjectUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetFoundingDate sets the "founding_date" field.
func (pu *ProjectUpdate) SetFoundingDate(s string) *ProjectUpdate {
	pu.mutation.SetFoundingDate(s)
	return pu
}

// SetNillableFoundingDate sets the "founding_date" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableFoundingDate(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetFoundingDate(*s)
	}
	return pu
}

// ClearFoundingDate clears the value of the "founding_date" field.
func (pu *ProjectUpdate) ClearFoundingDate() *ProjectUpdate {
	pu.mutation.ClearFoundingDate()
	return pu
}

// SetDissolutionDate sets the "dissolution_date" field.
func (pu *ProjectUpdate) SetDissolutionDate(s string) *ProjectUpdate {
	pu.mutation.SetDissolutionDate(s)
	return pu
}

// SetNillableDissolutionDate sets the "dissolution_date" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableDissolutionDate(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetDissolutionDate(*s)
	}
	return pu
}

// ClearDissolutionDate clears the value of the "dissolution_date" field.
func (pu *ProjectUpdate) ClearDissolutionDate() *ProjectUpdate {
	pu.mutation.ClearDissolutionDate()
	return pu
}

// SetAcronym sets the "acronym" field.
func (pu *ProjectUpdate) SetAcronym(s string) *ProjectUpdate {
	pu.mutation.SetAcronym(s)
	return pu
}

// SetNillableAcronym sets the "acronym" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableAcronym(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetAcronym(*s)
	}
	return pu
}

// ClearAcronym clears the value of the "acronym" field.
func (pu *ProjectUpdate) ClearAcronym() *ProjectUpdate {
	pu.mutation.ClearAcronym()
	return pu
}

// SetGrant sets the "grant" field.
func (pu *ProjectUpdate) SetGrant(s string) *ProjectUpdate {
	pu.mutation.SetGrant(s)
	return pu
}

// SetNillableGrant sets the "grant" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableGrant(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetGrant(*s)
	}
	return pu
}

// ClearGrant clears the value of the "grant" field.
func (pu *ProjectUpdate) ClearGrant() *ProjectUpdate {
	pu.mutation.ClearGrant()
	return pu
}

// SetFundingProgramme sets the "funding_programme" field.
func (pu *ProjectUpdate) SetFundingProgramme(s string) *ProjectUpdate {
	pu.mutation.SetFundingProgramme(s)
	return pu
}

// SetNillableFundingProgramme sets the "funding_programme" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableFundingProgramme(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetFundingProgramme(*s)
	}
	return pu
}

// ClearFundingProgramme clears the value of the "funding_programme" field.
func (pu *ProjectUpdate) ClearFundingProgramme() *ProjectUpdate {
	pu.mutation.ClearFundingProgramme()
	return pu
}

// SetModified sets the "modified" field.
func (pu *ProjectUpdate) SetModified(t time.Time) *ProjectUpdate {
	pu.mutation.SetModified(t)
	return pu
}

// Mutation returns the ProjectMutation object of the builder.
func (pu *ProjectUpdate) Mutation() *ProjectMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProjectUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProjectUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProjectUpdate) defaults() {
	if _, ok := pu.mutation.Modified(); !ok {
		v := project.UpdateDefaultModified()
		pu.mutation.SetModified(v)
	}
}

func (pu *ProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeString))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Identifier(); ok {
		_spec.SetField(project.FieldIdentifier, field.TypeJSON, value)
	}
	if value, ok := pu.mutation.AppendedIdentifier(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, project.FieldIdentifier, value)
		})
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if pu.mutation.NameCleared() {
		_spec.ClearField(project.FieldName, field.TypeString)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(project.FieldDescription, field.TypeString)
	}
	if value, ok := pu.mutation.FoundingDate(); ok {
		_spec.SetField(project.FieldFoundingDate, field.TypeString, value)
	}
	if pu.mutation.FoundingDateCleared() {
		_spec.ClearField(project.FieldFoundingDate, field.TypeString)
	}
	if value, ok := pu.mutation.DissolutionDate(); ok {
		_spec.SetField(project.FieldDissolutionDate, field.TypeString, value)
	}
	if pu.mutation.DissolutionDateCleared() {
		_spec.ClearField(project.FieldDissolutionDate, field.TypeString)
	}
	if value, ok := pu.mutation.Acronym(); ok {
		_spec.SetField(project.FieldAcronym, field.TypeString, value)
	}
	if pu.mutation.AcronymCleared() {
		_spec.ClearField(project.FieldAcronym, field.TypeString)
	}
	if value, ok := pu.mutation.Grant(); ok {
		_spec.SetField(project.FieldGrant, field.TypeString, value)
	}
	if pu.mutation.GrantCleared() {
		_spec.ClearField(project.FieldGrant, field.TypeString)
	}
	if value, ok := pu.mutation.FundingProgramme(); ok {
		_spec.SetField(project.FieldFundingProgramme, field.TypeString, value)
	}
	if pu.mutation.FundingProgrammeCleared() {
		_spec.ClearField(project.FieldFundingProgramme, field.TypeString)
	}
	if value, ok := pu.mutation.Modified(); ok {
		_spec.SetField(project.FieldModified, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProjectUpdateOne is the builder for updating a single Project entity.
type ProjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectMutation
}

// SetIdentifier sets the "identifier" field.
func (puo *ProjectUpdateOne) SetIdentifier(s []schema.Identifier) *ProjectUpdateOne {
	puo.mutation.SetIdentifier(s)
	return puo
}

// AppendIdentifier appends s to the "identifier" field.
func (puo *ProjectUpdateOne) AppendIdentifier(s []schema.Identifier) *ProjectUpdateOne {
	puo.mutation.AppendIdentifier(s)
	return puo
}

// SetName sets the "name" field.
func (puo *ProjectUpdateOne) SetName(s string) *ProjectUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableName(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// ClearName clears the value of the "name" field.
func (puo *ProjectUpdateOne) ClearName() *ProjectUpdateOne {
	puo.mutation.ClearName()
	return puo
}

// SetDescription sets the "description" field.
func (puo *ProjectUpdateOne) SetDescription(s string) *ProjectUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableDescription(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *ProjectUpdateOne) ClearDescription() *ProjectUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetFoundingDate sets the "founding_date" field.
func (puo *ProjectUpdateOne) SetFoundingDate(s string) *ProjectUpdateOne {
	puo.mutation.SetFoundingDate(s)
	return puo
}

// SetNillableFoundingDate sets the "founding_date" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableFoundingDate(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetFoundingDate(*s)
	}
	return puo
}

// ClearFoundingDate clears the value of the "founding_date" field.
func (puo *ProjectUpdateOne) ClearFoundingDate() *ProjectUpdateOne {
	puo.mutation.ClearFoundingDate()
	return puo
}

// SetDissolutionDate sets the "dissolution_date" field.
func (puo *ProjectUpdateOne) SetDissolutionDate(s string) *ProjectUpdateOne {
	puo.mutation.SetDissolutionDate(s)
	return puo
}

// SetNillableDissolutionDate sets the "dissolution_date" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableDissolutionDate(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetDissolutionDate(*s)
	}
	return puo
}

// ClearDissolutionDate clears the value of the "dissolution_date" field.
func (puo *ProjectUpdateOne) ClearDissolutionDate() *ProjectUpdateOne {
	puo.mutation.ClearDissolutionDate()
	return puo
}

// SetAcronym sets the "acronym" field.
func (puo *ProjectUpdateOne) SetAcronym(s string) *ProjectUpdateOne {
	puo.mutation.SetAcronym(s)
	return puo
}

// SetNillableAcronym sets the "acronym" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableAcronym(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetAcronym(*s)
	}
	return puo
}

// ClearAcronym clears the value of the "acronym" field.
func (puo *ProjectUpdateOne) ClearAcronym() *ProjectUpdateOne {
	puo.mutation.ClearAcronym()
	return puo
}

// SetGrant sets the "grant" field.
func (puo *ProjectUpdateOne) SetGrant(s string) *ProjectUpdateOne {
	puo.mutation.SetGrant(s)
	return puo
}

// SetNillableGrant sets the "grant" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableGrant(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetGrant(*s)
	}
	return puo
}

// ClearGrant clears the value of the "grant" field.
func (puo *ProjectUpdateOne) ClearGrant() *ProjectUpdateOne {
	puo.mutation.ClearGrant()
	return puo
}

// SetFundingProgramme sets the "funding_programme" field.
func (puo *ProjectUpdateOne) SetFundingProgramme(s string) *ProjectUpdateOne {
	puo.mutation.SetFundingProgramme(s)
	return puo
}

// SetNillableFundingProgramme sets the "funding_programme" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableFundingProgramme(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetFundingProgramme(*s)
	}
	return puo
}

// ClearFundingProgramme clears the value of the "funding_programme" field.
func (puo *ProjectUpdateOne) ClearFundingProgramme() *ProjectUpdateOne {
	puo.mutation.ClearFundingProgramme()
	return puo
}

// SetModified sets the "modified" field.
func (puo *ProjectUpdateOne) SetModified(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetModified(t)
	return puo
}

// Mutation returns the ProjectMutation object of the builder.
func (puo *ProjectUpdateOne) Mutation() *ProjectMutation {
	return puo.mutation
}

// Where appends a list predicates to the ProjectUpdate builder.
func (puo *ProjectUpdateOne) Where(ps ...predicate.Project) *ProjectUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProjectUpdateOne) Select(field string, fields ...string) *ProjectUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Project entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProjectUpdateOne) SaveX(ctx context.Context) *Project {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProjectUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProjectUpdateOne) defaults() {
	if _, ok := puo.mutation.Modified(); !ok {
		v := project.UpdateDefaultModified()
		puo.mutation.SetModified(v)
	}
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (_node *Project, err error) {
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeString))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Project.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, project.FieldID)
		for _, f := range fields {
			if !project.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != project.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Identifier(); ok {
		_spec.SetField(project.FieldIdentifier, field.TypeJSON, value)
	}
	if value, ok := puo.mutation.AppendedIdentifier(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, project.FieldIdentifier, value)
		})
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if puo.mutation.NameCleared() {
		_spec.ClearField(project.FieldName, field.TypeString)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(project.FieldDescription, field.TypeString)
	}
	if value, ok := puo.mutation.FoundingDate(); ok {
		_spec.SetField(project.FieldFoundingDate, field.TypeString, value)
	}
	if puo.mutation.FoundingDateCleared() {
		_spec.ClearField(project.FieldFoundingDate, field.TypeString)
	}
	if value, ok := puo.mutation.DissolutionDate(); ok {
		_spec.SetField(project.FieldDissolutionDate, field.TypeString, value)
	}
	if puo.mutation.DissolutionDateCleared() {
		_spec.ClearField(project.FieldDissolutionDate, field.TypeString)
	}
	if value, ok := puo.mutation.Acronym(); ok {
		_spec.SetField(project.FieldAcronym, field.TypeString, value)
	}
	if puo.mutation.AcronymCleared() {
		_spec.ClearField(project.FieldAcronym, field.TypeString)
	}
	if value, ok := puo.mutation.Grant(); ok {
		_spec.SetField(project.FieldGrant, field.TypeString, value)
	}
	if puo.mutation.GrantCleared() {
		_spec.ClearField(project.FieldGrant, field.TypeString)
	}
	if value, ok := puo.mutation.FundingProgramme(); ok {
		_spec.SetField(project.FieldFundingProgramme, field.TypeString, value)
	}
	if puo.mutation.FundingProgrammeCleared() {
		_spec.ClearField(project.FieldFundingProgramme, field.TypeString)
	}
	if value, ok := puo.mutation.Modified(); ok {
		_spec.SetField(project.FieldModified, field.TypeTime, value)
	}
	_node = &Project{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
