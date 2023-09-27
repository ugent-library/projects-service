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

// ClearIdentifier clears the value of the "identifier" field.
func (pu *ProjectUpdate) ClearIdentifier() *ProjectUpdate {
	pu.mutation.ClearIdentifier()
	return pu
}

// SetIsFundedBy sets the "is_funded_by" field.
func (pu *ProjectUpdate) SetIsFundedBy(s schema.Grant) *ProjectUpdate {
	pu.mutation.SetIsFundedBy(s)
	return pu
}

// SetNillableIsFundedBy sets the "is_funded_by" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableIsFundedBy(s *schema.Grant) *ProjectUpdate {
	if s != nil {
		pu.SetIsFundedBy(*s)
	}
	return pu
}

// ClearIsFundedBy clears the value of the "is_funded_by" field.
func (pu *ProjectUpdate) ClearIsFundedBy() *ProjectUpdate {
	pu.mutation.ClearIsFundedBy()
	return pu
}

// SetName sets the "name" field.
func (pu *ProjectUpdate) SetName(s string) *ProjectUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetDescription sets the "description" field.
func (pu *ProjectUpdate) SetDescription(s string) *ProjectUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetFoundingDate sets the "founding_date" field.
func (pu *ProjectUpdate) SetFoundingDate(s string) *ProjectUpdate {
	pu.mutation.SetFoundingDate(s)
	return pu
}

// SetDissolutionDate sets the "dissolution_date" field.
func (pu *ProjectUpdate) SetDissolutionDate(s string) *ProjectUpdate {
	pu.mutation.SetDissolutionDate(s)
	return pu
}

// SetHasAcronym sets the "has_acronym" field.
func (pu *ProjectUpdate) SetHasAcronym(s string) *ProjectUpdate {
	pu.mutation.SetHasAcronym(s)
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
	if pu.mutation.IdentifierCleared() {
		_spec.ClearField(project.FieldIdentifier, field.TypeJSON)
	}
	if value, ok := pu.mutation.IsFundedBy(); ok {
		_spec.SetField(project.FieldIsFundedBy, field.TypeJSON, value)
	}
	if pu.mutation.IsFundedByCleared() {
		_spec.ClearField(project.FieldIsFundedBy, field.TypeJSON)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.FoundingDate(); ok {
		_spec.SetField(project.FieldFoundingDate, field.TypeString, value)
	}
	if value, ok := pu.mutation.DissolutionDate(); ok {
		_spec.SetField(project.FieldDissolutionDate, field.TypeString, value)
	}
	if value, ok := pu.mutation.HasAcronym(); ok {
		_spec.SetField(project.FieldHasAcronym, field.TypeString, value)
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

// ClearIdentifier clears the value of the "identifier" field.
func (puo *ProjectUpdateOne) ClearIdentifier() *ProjectUpdateOne {
	puo.mutation.ClearIdentifier()
	return puo
}

// SetIsFundedBy sets the "is_funded_by" field.
func (puo *ProjectUpdateOne) SetIsFundedBy(s schema.Grant) *ProjectUpdateOne {
	puo.mutation.SetIsFundedBy(s)
	return puo
}

// SetNillableIsFundedBy sets the "is_funded_by" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableIsFundedBy(s *schema.Grant) *ProjectUpdateOne {
	if s != nil {
		puo.SetIsFundedBy(*s)
	}
	return puo
}

// ClearIsFundedBy clears the value of the "is_funded_by" field.
func (puo *ProjectUpdateOne) ClearIsFundedBy() *ProjectUpdateOne {
	puo.mutation.ClearIsFundedBy()
	return puo
}

// SetName sets the "name" field.
func (puo *ProjectUpdateOne) SetName(s string) *ProjectUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetDescription sets the "description" field.
func (puo *ProjectUpdateOne) SetDescription(s string) *ProjectUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetFoundingDate sets the "founding_date" field.
func (puo *ProjectUpdateOne) SetFoundingDate(s string) *ProjectUpdateOne {
	puo.mutation.SetFoundingDate(s)
	return puo
}

// SetDissolutionDate sets the "dissolution_date" field.
func (puo *ProjectUpdateOne) SetDissolutionDate(s string) *ProjectUpdateOne {
	puo.mutation.SetDissolutionDate(s)
	return puo
}

// SetHasAcronym sets the "has_acronym" field.
func (puo *ProjectUpdateOne) SetHasAcronym(s string) *ProjectUpdateOne {
	puo.mutation.SetHasAcronym(s)
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
	if puo.mutation.IdentifierCleared() {
		_spec.ClearField(project.FieldIdentifier, field.TypeJSON)
	}
	if value, ok := puo.mutation.IsFundedBy(); ok {
		_spec.SetField(project.FieldIsFundedBy, field.TypeJSON, value)
	}
	if puo.mutation.IsFundedByCleared() {
		_spec.ClearField(project.FieldIsFundedBy, field.TypeJSON)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.FoundingDate(); ok {
		_spec.SetField(project.FieldFoundingDate, field.TypeString, value)
	}
	if value, ok := puo.mutation.DissolutionDate(); ok {
		_spec.SetField(project.FieldDissolutionDate, field.TypeString, value)
	}
	if value, ok := puo.mutation.HasAcronym(); ok {
		_spec.SetField(project.FieldHasAcronym, field.TypeString, value)
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
