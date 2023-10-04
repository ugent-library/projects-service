// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/projects/ent/project"
	"github.com/ugent-library/projects/ent/schema"
)

// ProjectCreate is the builder for creating a Project entity.
type ProjectCreate struct {
	config
	mutation *ProjectMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetIdentifier sets the "identifier" field.
func (pc *ProjectCreate) SetIdentifier(s []schema.Identifier) *ProjectCreate {
	pc.mutation.SetIdentifier(s)
	return pc
}

// SetName sets the "name" field.
func (pc *ProjectCreate) SetName(s string) *ProjectCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableName(s *string) *ProjectCreate {
	if s != nil {
		pc.SetName(*s)
	}
	return pc
}

// SetDescription sets the "description" field.
func (pc *ProjectCreate) SetDescription(s string) *ProjectCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableDescription(s *string) *ProjectCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetFoundingDate sets the "founding_date" field.
func (pc *ProjectCreate) SetFoundingDate(s string) *ProjectCreate {
	pc.mutation.SetFoundingDate(s)
	return pc
}

// SetNillableFoundingDate sets the "founding_date" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableFoundingDate(s *string) *ProjectCreate {
	if s != nil {
		pc.SetFoundingDate(*s)
	}
	return pc
}

// SetDissolutionDate sets the "dissolution_date" field.
func (pc *ProjectCreate) SetDissolutionDate(s string) *ProjectCreate {
	pc.mutation.SetDissolutionDate(s)
	return pc
}

// SetNillableDissolutionDate sets the "dissolution_date" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableDissolutionDate(s *string) *ProjectCreate {
	if s != nil {
		pc.SetDissolutionDate(*s)
	}
	return pc
}

// SetAcronym sets the "acronym" field.
func (pc *ProjectCreate) SetAcronym(s string) *ProjectCreate {
	pc.mutation.SetAcronym(s)
	return pc
}

// SetNillableAcronym sets the "acronym" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableAcronym(s *string) *ProjectCreate {
	if s != nil {
		pc.SetAcronym(*s)
	}
	return pc
}

// SetGrant sets the "grant" field.
func (pc *ProjectCreate) SetGrant(s string) *ProjectCreate {
	pc.mutation.SetGrant(s)
	return pc
}

// SetNillableGrant sets the "grant" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableGrant(s *string) *ProjectCreate {
	if s != nil {
		pc.SetGrant(*s)
	}
	return pc
}

// SetFundingProgramme sets the "funding_programme" field.
func (pc *ProjectCreate) SetFundingProgramme(s string) *ProjectCreate {
	pc.mutation.SetFundingProgramme(s)
	return pc
}

// SetNillableFundingProgramme sets the "funding_programme" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableFundingProgramme(s *string) *ProjectCreate {
	if s != nil {
		pc.SetFundingProgramme(*s)
	}
	return pc
}

// SetCreated sets the "created" field.
func (pc *ProjectCreate) SetCreated(t time.Time) *ProjectCreate {
	pc.mutation.SetCreated(t)
	return pc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableCreated(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetCreated(*t)
	}
	return pc
}

// SetModified sets the "modified" field.
func (pc *ProjectCreate) SetModified(t time.Time) *ProjectCreate {
	pc.mutation.SetModified(t)
	return pc
}

// SetNillableModified sets the "modified" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableModified(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetModified(*t)
	}
	return pc
}

// SetTs sets the "ts" field.
func (pc *ProjectCreate) SetTs(s string) *ProjectCreate {
	pc.mutation.SetTs(s)
	return pc
}

// SetNillableTs sets the "ts" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableTs(s *string) *ProjectCreate {
	if s != nil {
		pc.SetTs(*s)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProjectCreate) SetID(s string) *ProjectCreate {
	pc.mutation.SetID(s)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableID(s *string) *ProjectCreate {
	if s != nil {
		pc.SetID(*s)
	}
	return pc
}

// Mutation returns the ProjectMutation object of the builder.
func (pc *ProjectCreate) Mutation() *ProjectMutation {
	return pc.mutation
}

// Save creates the Project in the database.
func (pc *ProjectCreate) Save(ctx context.Context) (*Project, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProjectCreate) SaveX(ctx context.Context) *Project {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProjectCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProjectCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProjectCreate) defaults() {
	if _, ok := pc.mutation.Identifier(); !ok {
		v := project.DefaultIdentifier
		pc.mutation.SetIdentifier(v)
	}
	if _, ok := pc.mutation.Created(); !ok {
		v := project.DefaultCreated()
		pc.mutation.SetCreated(v)
	}
	if _, ok := pc.mutation.Modified(); !ok {
		v := project.DefaultModified()
		pc.mutation.SetModified(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := project.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProjectCreate) check() error {
	if _, ok := pc.mutation.Identifier(); !ok {
		return &ValidationError{Name: "identifier", err: errors.New(`ent: missing required field "Project.identifier"`)}
	}
	if _, ok := pc.mutation.Created(); !ok {
		return &ValidationError{Name: "created", err: errors.New(`ent: missing required field "Project.created"`)}
	}
	if _, ok := pc.mutation.Modified(); !ok {
		return &ValidationError{Name: "modified", err: errors.New(`ent: missing required field "Project.modified"`)}
	}
	return nil
}

func (pc *ProjectCreate) sqlSave(ctx context.Context) (*Project, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Project.ID type: %T", _spec.ID.Value)
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProjectCreate) createSpec() (*Project, *sqlgraph.CreateSpec) {
	var (
		_node = &Project{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(project.Table, sqlgraph.NewFieldSpec(project.FieldID, field.TypeString))
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Identifier(); ok {
		_spec.SetField(project.FieldIdentifier, field.TypeJSON, value)
		_node.Identifier = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
		_node.Name = &value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if value, ok := pc.mutation.FoundingDate(); ok {
		_spec.SetField(project.FieldFoundingDate, field.TypeString, value)
		_node.FoundingDate = &value
	}
	if value, ok := pc.mutation.DissolutionDate(); ok {
		_spec.SetField(project.FieldDissolutionDate, field.TypeString, value)
		_node.DissolutionDate = &value
	}
	if value, ok := pc.mutation.Acronym(); ok {
		_spec.SetField(project.FieldAcronym, field.TypeString, value)
		_node.Acronym = &value
	}
	if value, ok := pc.mutation.Grant(); ok {
		_spec.SetField(project.FieldGrant, field.TypeString, value)
		_node.Grant = &value
	}
	if value, ok := pc.mutation.FundingProgramme(); ok {
		_spec.SetField(project.FieldFundingProgramme, field.TypeString, value)
		_node.FundingProgramme = &value
	}
	if value, ok := pc.mutation.Created(); ok {
		_spec.SetField(project.FieldCreated, field.TypeTime, value)
		_node.Created = value
	}
	if value, ok := pc.mutation.Modified(); ok {
		_spec.SetField(project.FieldModified, field.TypeTime, value)
		_node.Modified = value
	}
	if value, ok := pc.mutation.Ts(); ok {
		_spec.SetField(project.FieldTs, field.TypeString, value)
		_node.Ts = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Project.Create().
//		SetIdentifier(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProjectUpsert) {
//			SetIdentifier(v+v).
//		}).
//		Exec(ctx)
func (pc *ProjectCreate) OnConflict(opts ...sql.ConflictOption) *ProjectUpsertOne {
	pc.conflict = opts
	return &ProjectUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Project.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *ProjectCreate) OnConflictColumns(columns ...string) *ProjectUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &ProjectUpsertOne{
		create: pc,
	}
}

type (
	// ProjectUpsertOne is the builder for "upsert"-ing
	//  one Project node.
	ProjectUpsertOne struct {
		create *ProjectCreate
	}

	// ProjectUpsert is the "OnConflict" setter.
	ProjectUpsert struct {
		*sql.UpdateSet
	}
)

// SetIdentifier sets the "identifier" field.
func (u *ProjectUpsert) SetIdentifier(v []schema.Identifier) *ProjectUpsert {
	u.Set(project.FieldIdentifier, v)
	return u
}

// UpdateIdentifier sets the "identifier" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateIdentifier() *ProjectUpsert {
	u.SetExcluded(project.FieldIdentifier)
	return u
}

// SetName sets the "name" field.
func (u *ProjectUpsert) SetName(v string) *ProjectUpsert {
	u.Set(project.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateName() *ProjectUpsert {
	u.SetExcluded(project.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *ProjectUpsert) ClearName() *ProjectUpsert {
	u.SetNull(project.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *ProjectUpsert) SetDescription(v string) *ProjectUpsert {
	u.Set(project.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateDescription() *ProjectUpsert {
	u.SetExcluded(project.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *ProjectUpsert) ClearDescription() *ProjectUpsert {
	u.SetNull(project.FieldDescription)
	return u
}

// SetFoundingDate sets the "founding_date" field.
func (u *ProjectUpsert) SetFoundingDate(v string) *ProjectUpsert {
	u.Set(project.FieldFoundingDate, v)
	return u
}

// UpdateFoundingDate sets the "founding_date" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateFoundingDate() *ProjectUpsert {
	u.SetExcluded(project.FieldFoundingDate)
	return u
}

// ClearFoundingDate clears the value of the "founding_date" field.
func (u *ProjectUpsert) ClearFoundingDate() *ProjectUpsert {
	u.SetNull(project.FieldFoundingDate)
	return u
}

// SetDissolutionDate sets the "dissolution_date" field.
func (u *ProjectUpsert) SetDissolutionDate(v string) *ProjectUpsert {
	u.Set(project.FieldDissolutionDate, v)
	return u
}

// UpdateDissolutionDate sets the "dissolution_date" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateDissolutionDate() *ProjectUpsert {
	u.SetExcluded(project.FieldDissolutionDate)
	return u
}

// ClearDissolutionDate clears the value of the "dissolution_date" field.
func (u *ProjectUpsert) ClearDissolutionDate() *ProjectUpsert {
	u.SetNull(project.FieldDissolutionDate)
	return u
}

// SetAcronym sets the "acronym" field.
func (u *ProjectUpsert) SetAcronym(v string) *ProjectUpsert {
	u.Set(project.FieldAcronym, v)
	return u
}

// UpdateAcronym sets the "acronym" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateAcronym() *ProjectUpsert {
	u.SetExcluded(project.FieldAcronym)
	return u
}

// ClearAcronym clears the value of the "acronym" field.
func (u *ProjectUpsert) ClearAcronym() *ProjectUpsert {
	u.SetNull(project.FieldAcronym)
	return u
}

// SetGrant sets the "grant" field.
func (u *ProjectUpsert) SetGrant(v string) *ProjectUpsert {
	u.Set(project.FieldGrant, v)
	return u
}

// UpdateGrant sets the "grant" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateGrant() *ProjectUpsert {
	u.SetExcluded(project.FieldGrant)
	return u
}

// ClearGrant clears the value of the "grant" field.
func (u *ProjectUpsert) ClearGrant() *ProjectUpsert {
	u.SetNull(project.FieldGrant)
	return u
}

// SetFundingProgramme sets the "funding_programme" field.
func (u *ProjectUpsert) SetFundingProgramme(v string) *ProjectUpsert {
	u.Set(project.FieldFundingProgramme, v)
	return u
}

// UpdateFundingProgramme sets the "funding_programme" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateFundingProgramme() *ProjectUpsert {
	u.SetExcluded(project.FieldFundingProgramme)
	return u
}

// ClearFundingProgramme clears the value of the "funding_programme" field.
func (u *ProjectUpsert) ClearFundingProgramme() *ProjectUpsert {
	u.SetNull(project.FieldFundingProgramme)
	return u
}

// SetModified sets the "modified" field.
func (u *ProjectUpsert) SetModified(v time.Time) *ProjectUpsert {
	u.Set(project.FieldModified, v)
	return u
}

// UpdateModified sets the "modified" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateModified() *ProjectUpsert {
	u.SetExcluded(project.FieldModified)
	return u
}

// SetTs sets the "ts" field.
func (u *ProjectUpsert) SetTs(v string) *ProjectUpsert {
	u.Set(project.FieldTs, v)
	return u
}

// UpdateTs sets the "ts" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateTs() *ProjectUpsert {
	u.SetExcluded(project.FieldTs)
	return u
}

// ClearTs clears the value of the "ts" field.
func (u *ProjectUpsert) ClearTs() *ProjectUpsert {
	u.SetNull(project.FieldTs)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Project.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(project.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ProjectUpsertOne) UpdateNewValues() *ProjectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(project.FieldID)
		}
		if _, exists := u.create.mutation.Created(); exists {
			s.SetIgnore(project.FieldCreated)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Project.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ProjectUpsertOne) Ignore() *ProjectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProjectUpsertOne) DoNothing() *ProjectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProjectCreate.OnConflict
// documentation for more info.
func (u *ProjectUpsertOne) Update(set func(*ProjectUpsert)) *ProjectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProjectUpsert{UpdateSet: update})
	}))
	return u
}

// SetIdentifier sets the "identifier" field.
func (u *ProjectUpsertOne) SetIdentifier(v []schema.Identifier) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetIdentifier(v)
	})
}

// UpdateIdentifier sets the "identifier" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateIdentifier() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateIdentifier()
	})
}

// SetName sets the "name" field.
func (u *ProjectUpsertOne) SetName(v string) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateName() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *ProjectUpsertOne) ClearName() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearName()
	})
}

// SetDescription sets the "description" field.
func (u *ProjectUpsertOne) SetDescription(v string) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateDescription() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *ProjectUpsertOne) ClearDescription() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearDescription()
	})
}

// SetFoundingDate sets the "founding_date" field.
func (u *ProjectUpsertOne) SetFoundingDate(v string) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetFoundingDate(v)
	})
}

// UpdateFoundingDate sets the "founding_date" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateFoundingDate() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateFoundingDate()
	})
}

// ClearFoundingDate clears the value of the "founding_date" field.
func (u *ProjectUpsertOne) ClearFoundingDate() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearFoundingDate()
	})
}

// SetDissolutionDate sets the "dissolution_date" field.
func (u *ProjectUpsertOne) SetDissolutionDate(v string) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetDissolutionDate(v)
	})
}

// UpdateDissolutionDate sets the "dissolution_date" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateDissolutionDate() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateDissolutionDate()
	})
}

// ClearDissolutionDate clears the value of the "dissolution_date" field.
func (u *ProjectUpsertOne) ClearDissolutionDate() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearDissolutionDate()
	})
}

// SetAcronym sets the "acronym" field.
func (u *ProjectUpsertOne) SetAcronym(v string) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetAcronym(v)
	})
}

// UpdateAcronym sets the "acronym" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateAcronym() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateAcronym()
	})
}

// ClearAcronym clears the value of the "acronym" field.
func (u *ProjectUpsertOne) ClearAcronym() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearAcronym()
	})
}

// SetGrant sets the "grant" field.
func (u *ProjectUpsertOne) SetGrant(v string) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetGrant(v)
	})
}

// UpdateGrant sets the "grant" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateGrant() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateGrant()
	})
}

// ClearGrant clears the value of the "grant" field.
func (u *ProjectUpsertOne) ClearGrant() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearGrant()
	})
}

// SetFundingProgramme sets the "funding_programme" field.
func (u *ProjectUpsertOne) SetFundingProgramme(v string) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetFundingProgramme(v)
	})
}

// UpdateFundingProgramme sets the "funding_programme" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateFundingProgramme() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateFundingProgramme()
	})
}

// ClearFundingProgramme clears the value of the "funding_programme" field.
func (u *ProjectUpsertOne) ClearFundingProgramme() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearFundingProgramme()
	})
}

// SetModified sets the "modified" field.
func (u *ProjectUpsertOne) SetModified(v time.Time) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetModified(v)
	})
}

// UpdateModified sets the "modified" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateModified() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateModified()
	})
}

// SetTs sets the "ts" field.
func (u *ProjectUpsertOne) SetTs(v string) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetTs(v)
	})
}

// UpdateTs sets the "ts" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateTs() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateTs()
	})
}

// ClearTs clears the value of the "ts" field.
func (u *ProjectUpsertOne) ClearTs() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearTs()
	})
}

// Exec executes the query.
func (u *ProjectUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProjectCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProjectUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ProjectUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ProjectUpsertOne.ID is not supported by MySQL driver. Use ProjectUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProjectUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ProjectCreateBulk is the builder for creating many Project entities in bulk.
type ProjectCreateBulk struct {
	config
	builders []*ProjectCreate
	conflict []sql.ConflictOption
}

// Save creates the Project entities in the database.
func (pcb *ProjectCreateBulk) Save(ctx context.Context) ([]*Project, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Project, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProjectCreateBulk) SaveX(ctx context.Context) []*Project {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProjectCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProjectCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Project.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProjectUpsert) {
//			SetIdentifier(v+v).
//		}).
//		Exec(ctx)
func (pcb *ProjectCreateBulk) OnConflict(opts ...sql.ConflictOption) *ProjectUpsertBulk {
	pcb.conflict = opts
	return &ProjectUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Project.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *ProjectCreateBulk) OnConflictColumns(columns ...string) *ProjectUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &ProjectUpsertBulk{
		create: pcb,
	}
}

// ProjectUpsertBulk is the builder for "upsert"-ing
// a bulk of Project nodes.
type ProjectUpsertBulk struct {
	create *ProjectCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Project.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(project.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ProjectUpsertBulk) UpdateNewValues() *ProjectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(project.FieldID)
			}
			if _, exists := b.mutation.Created(); exists {
				s.SetIgnore(project.FieldCreated)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Project.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ProjectUpsertBulk) Ignore() *ProjectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProjectUpsertBulk) DoNothing() *ProjectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProjectCreateBulk.OnConflict
// documentation for more info.
func (u *ProjectUpsertBulk) Update(set func(*ProjectUpsert)) *ProjectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProjectUpsert{UpdateSet: update})
	}))
	return u
}

// SetIdentifier sets the "identifier" field.
func (u *ProjectUpsertBulk) SetIdentifier(v []schema.Identifier) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetIdentifier(v)
	})
}

// UpdateIdentifier sets the "identifier" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateIdentifier() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateIdentifier()
	})
}

// SetName sets the "name" field.
func (u *ProjectUpsertBulk) SetName(v string) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateName() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *ProjectUpsertBulk) ClearName() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearName()
	})
}

// SetDescription sets the "description" field.
func (u *ProjectUpsertBulk) SetDescription(v string) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateDescription() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *ProjectUpsertBulk) ClearDescription() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearDescription()
	})
}

// SetFoundingDate sets the "founding_date" field.
func (u *ProjectUpsertBulk) SetFoundingDate(v string) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetFoundingDate(v)
	})
}

// UpdateFoundingDate sets the "founding_date" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateFoundingDate() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateFoundingDate()
	})
}

// ClearFoundingDate clears the value of the "founding_date" field.
func (u *ProjectUpsertBulk) ClearFoundingDate() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearFoundingDate()
	})
}

// SetDissolutionDate sets the "dissolution_date" field.
func (u *ProjectUpsertBulk) SetDissolutionDate(v string) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetDissolutionDate(v)
	})
}

// UpdateDissolutionDate sets the "dissolution_date" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateDissolutionDate() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateDissolutionDate()
	})
}

// ClearDissolutionDate clears the value of the "dissolution_date" field.
func (u *ProjectUpsertBulk) ClearDissolutionDate() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearDissolutionDate()
	})
}

// SetAcronym sets the "acronym" field.
func (u *ProjectUpsertBulk) SetAcronym(v string) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetAcronym(v)
	})
}

// UpdateAcronym sets the "acronym" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateAcronym() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateAcronym()
	})
}

// ClearAcronym clears the value of the "acronym" field.
func (u *ProjectUpsertBulk) ClearAcronym() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearAcronym()
	})
}

// SetGrant sets the "grant" field.
func (u *ProjectUpsertBulk) SetGrant(v string) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetGrant(v)
	})
}

// UpdateGrant sets the "grant" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateGrant() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateGrant()
	})
}

// ClearGrant clears the value of the "grant" field.
func (u *ProjectUpsertBulk) ClearGrant() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearGrant()
	})
}

// SetFundingProgramme sets the "funding_programme" field.
func (u *ProjectUpsertBulk) SetFundingProgramme(v string) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetFundingProgramme(v)
	})
}

// UpdateFundingProgramme sets the "funding_programme" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateFundingProgramme() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateFundingProgramme()
	})
}

// ClearFundingProgramme clears the value of the "funding_programme" field.
func (u *ProjectUpsertBulk) ClearFundingProgramme() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearFundingProgramme()
	})
}

// SetModified sets the "modified" field.
func (u *ProjectUpsertBulk) SetModified(v time.Time) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetModified(v)
	})
}

// UpdateModified sets the "modified" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateModified() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateModified()
	})
}

// SetTs sets the "ts" field.
func (u *ProjectUpsertBulk) SetTs(v string) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetTs(v)
	})
}

// UpdateTs sets the "ts" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateTs() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateTs()
	})
}

// ClearTs clears the value of the "ts" field.
func (u *ProjectUpsertBulk) ClearTs() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearTs()
	})
}

// Exec executes the query.
func (u *ProjectUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ProjectCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProjectCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProjectUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
