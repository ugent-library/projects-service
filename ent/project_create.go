// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/projects/ent/project"
	"github.com/ugent-library/projects/ent/projectidentifier"
	"github.com/ugent-library/projects/ent/schema"
)

// ProjectCreate is the builder for creating a Project entity.
type ProjectCreate struct {
	config
	mutation *ProjectMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetProjectIdentifierID sets the "project_identifier_id" field.
func (pc *ProjectCreate) SetProjectIdentifierID(i int) *ProjectCreate {
	pc.mutation.SetProjectIdentifierID(i)
	return pc
}

// SetIdentifier sets the "identifier" field.
func (pc *ProjectCreate) SetIdentifier(s schema.Identifier) *ProjectCreate {
	pc.mutation.SetIdentifier(s)
	return pc
}

// SetNillableIdentifier sets the "identifier" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableIdentifier(s *schema.Identifier) *ProjectCreate {
	if s != nil {
		pc.SetIdentifier(*s)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *ProjectCreate) SetName(ss schema.TranslatedString) *ProjectCreate {
	pc.mutation.SetName(ss)
	return pc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableName(ss *schema.TranslatedString) *ProjectCreate {
	if ss != nil {
		pc.SetName(*ss)
	}
	return pc
}

// SetDescription sets the "description" field.
func (pc *ProjectCreate) SetDescription(ss schema.TranslatedString) *ProjectCreate {
	pc.mutation.SetDescription(ss)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableDescription(ss *schema.TranslatedString) *ProjectCreate {
	if ss != nil {
		pc.SetDescription(*ss)
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

// SetGrantID sets the "grant_id" field.
func (pc *ProjectCreate) SetGrantID(s string) *ProjectCreate {
	pc.mutation.SetGrantID(s)
	return pc
}

// SetNillableGrantID sets the "grant_id" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableGrantID(s *string) *ProjectCreate {
	if s != nil {
		pc.SetGrantID(*s)
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

// SetDeleted sets the "deleted" field.
func (pc *ProjectCreate) SetDeleted(b bool) *ProjectCreate {
	pc.mutation.SetDeleted(b)
	return pc
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableDeleted(b *bool) *ProjectCreate {
	if b != nil {
		pc.SetDeleted(*b)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProjectCreate) SetCreatedAt(t time.Time) *ProjectCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableCreatedAt(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ProjectCreate) SetUpdatedAt(t time.Time) *ProjectCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableUpdatedAt(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetIdentifiedByID sets the "identifiedBy" edge to the ProjectIdentifier entity by ID.
func (pc *ProjectCreate) SetIdentifiedByID(id int) *ProjectCreate {
	pc.mutation.SetIdentifiedByID(id)
	return pc
}

// SetIdentifiedBy sets the "identifiedBy" edge to the ProjectIdentifier entity.
func (pc *ProjectCreate) SetIdentifiedBy(p *ProjectIdentifier) *ProjectCreate {
	return pc.SetIdentifiedByID(p.ID)
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
	if _, ok := pc.mutation.Name(); !ok {
		v := project.DefaultName
		pc.mutation.SetName(v)
	}
	if _, ok := pc.mutation.Description(); !ok {
		v := project.DefaultDescription
		pc.mutation.SetDescription(v)
	}
	if _, ok := pc.mutation.Deleted(); !ok {
		v := project.DefaultDeleted
		pc.mutation.SetDeleted(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := project.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := project.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProjectCreate) check() error {
	if _, ok := pc.mutation.ProjectIdentifierID(); !ok {
		return &ValidationError{Name: "project_identifier_id", err: errors.New(`ent: missing required field "Project.project_identifier_id"`)}
	}
	if _, ok := pc.mutation.Identifier(); !ok {
		return &ValidationError{Name: "identifier", err: errors.New(`ent: missing required field "Project.identifier"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Project.name"`)}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Project.description"`)}
	}
	if _, ok := pc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "Project.deleted"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Project.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Project.updated_at"`)}
	}
	if _, ok := pc.mutation.IdentifiedByID(); !ok {
		return &ValidationError{Name: "identifiedBy", err: errors.New(`ent: missing required edge "Project.identifiedBy"`)}
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
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProjectCreate) createSpec() (*Project, *sqlgraph.CreateSpec) {
	var (
		_node = &Project{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(project.Table, sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pc.conflict
	if value, ok := pc.mutation.Identifier(); ok {
		_spec.SetField(project.FieldIdentifier, field.TypeJSON, value)
		_node.Identifier = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeJSON, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeJSON, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.FoundingDate(); ok {
		_spec.SetField(project.FieldFoundingDate, field.TypeString, value)
		_node.FoundingDate = value
	}
	if value, ok := pc.mutation.DissolutionDate(); ok {
		_spec.SetField(project.FieldDissolutionDate, field.TypeString, value)
		_node.DissolutionDate = value
	}
	if value, ok := pc.mutation.Acronym(); ok {
		_spec.SetField(project.FieldAcronym, field.TypeString, value)
		_node.Acronym = value
	}
	if value, ok := pc.mutation.GrantID(); ok {
		_spec.SetField(project.FieldGrantID, field.TypeString, value)
		_node.GrantID = value
	}
	if value, ok := pc.mutation.FundingProgramme(); ok {
		_spec.SetField(project.FieldFundingProgramme, field.TypeString, value)
		_node.FundingProgramme = value
	}
	if value, ok := pc.mutation.Deleted(); ok {
		_spec.SetField(project.FieldDeleted, field.TypeBool, value)
		_node.Deleted = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(project.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(project.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := pc.mutation.IdentifiedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.IdentifiedByTable,
			Columns: []string{project.IdentifiedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projectidentifier.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProjectIdentifierID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Project.Create().
//		SetProjectIdentifierID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProjectUpsert) {
//			SetProjectIdentifierID(v+v).
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

// SetProjectIdentifierID sets the "project_identifier_id" field.
func (u *ProjectUpsert) SetProjectIdentifierID(v int) *ProjectUpsert {
	u.Set(project.FieldProjectIdentifierID, v)
	return u
}

// UpdateProjectIdentifierID sets the "project_identifier_id" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateProjectIdentifierID() *ProjectUpsert {
	u.SetExcluded(project.FieldProjectIdentifierID)
	return u
}

// SetIdentifier sets the "identifier" field.
func (u *ProjectUpsert) SetIdentifier(v schema.Identifier) *ProjectUpsert {
	u.Set(project.FieldIdentifier, v)
	return u
}

// UpdateIdentifier sets the "identifier" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateIdentifier() *ProjectUpsert {
	u.SetExcluded(project.FieldIdentifier)
	return u
}

// SetName sets the "name" field.
func (u *ProjectUpsert) SetName(v schema.TranslatedString) *ProjectUpsert {
	u.Set(project.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateName() *ProjectUpsert {
	u.SetExcluded(project.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *ProjectUpsert) SetDescription(v schema.TranslatedString) *ProjectUpsert {
	u.Set(project.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateDescription() *ProjectUpsert {
	u.SetExcluded(project.FieldDescription)
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

// SetGrantID sets the "grant_id" field.
func (u *ProjectUpsert) SetGrantID(v string) *ProjectUpsert {
	u.Set(project.FieldGrantID, v)
	return u
}

// UpdateGrantID sets the "grant_id" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateGrantID() *ProjectUpsert {
	u.SetExcluded(project.FieldGrantID)
	return u
}

// ClearGrantID clears the value of the "grant_id" field.
func (u *ProjectUpsert) ClearGrantID() *ProjectUpsert {
	u.SetNull(project.FieldGrantID)
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

// SetDeleted sets the "deleted" field.
func (u *ProjectUpsert) SetDeleted(v bool) *ProjectUpsert {
	u.Set(project.FieldDeleted, v)
	return u
}

// UpdateDeleted sets the "deleted" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateDeleted() *ProjectUpsert {
	u.SetExcluded(project.FieldDeleted)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProjectUpsert) SetUpdatedAt(v time.Time) *ProjectUpsert {
	u.Set(project.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateUpdatedAt() *ProjectUpsert {
	u.SetExcluded(project.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Project.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ProjectUpsertOne) UpdateNewValues() *ProjectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(project.FieldCreatedAt)
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

// SetProjectIdentifierID sets the "project_identifier_id" field.
func (u *ProjectUpsertOne) SetProjectIdentifierID(v int) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetProjectIdentifierID(v)
	})
}

// UpdateProjectIdentifierID sets the "project_identifier_id" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateProjectIdentifierID() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateProjectIdentifierID()
	})
}

// SetIdentifier sets the "identifier" field.
func (u *ProjectUpsertOne) SetIdentifier(v schema.Identifier) *ProjectUpsertOne {
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
func (u *ProjectUpsertOne) SetName(v schema.TranslatedString) *ProjectUpsertOne {
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

// SetDescription sets the "description" field.
func (u *ProjectUpsertOne) SetDescription(v schema.TranslatedString) *ProjectUpsertOne {
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

// SetGrantID sets the "grant_id" field.
func (u *ProjectUpsertOne) SetGrantID(v string) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetGrantID(v)
	})
}

// UpdateGrantID sets the "grant_id" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateGrantID() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateGrantID()
	})
}

// ClearGrantID clears the value of the "grant_id" field.
func (u *ProjectUpsertOne) ClearGrantID() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearGrantID()
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

// SetDeleted sets the "deleted" field.
func (u *ProjectUpsertOne) SetDeleted(v bool) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetDeleted(v)
	})
}

// UpdateDeleted sets the "deleted" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateDeleted() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateDeleted()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProjectUpsertOne) SetUpdatedAt(v time.Time) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateUpdatedAt() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateUpdatedAt()
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
func (u *ProjectUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProjectUpsertOne) IDX(ctx context.Context) int {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
//			SetProjectIdentifierID(v+v).
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
//		).
//		Exec(ctx)
func (u *ProjectUpsertBulk) UpdateNewValues() *ProjectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(project.FieldCreatedAt)
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

// SetProjectIdentifierID sets the "project_identifier_id" field.
func (u *ProjectUpsertBulk) SetProjectIdentifierID(v int) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetProjectIdentifierID(v)
	})
}

// UpdateProjectIdentifierID sets the "project_identifier_id" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateProjectIdentifierID() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateProjectIdentifierID()
	})
}

// SetIdentifier sets the "identifier" field.
func (u *ProjectUpsertBulk) SetIdentifier(v schema.Identifier) *ProjectUpsertBulk {
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
func (u *ProjectUpsertBulk) SetName(v schema.TranslatedString) *ProjectUpsertBulk {
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

// SetDescription sets the "description" field.
func (u *ProjectUpsertBulk) SetDescription(v schema.TranslatedString) *ProjectUpsertBulk {
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

// SetGrantID sets the "grant_id" field.
func (u *ProjectUpsertBulk) SetGrantID(v string) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetGrantID(v)
	})
}

// UpdateGrantID sets the "grant_id" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateGrantID() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateGrantID()
	})
}

// ClearGrantID clears the value of the "grant_id" field.
func (u *ProjectUpsertBulk) ClearGrantID() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearGrantID()
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

// SetDeleted sets the "deleted" field.
func (u *ProjectUpsertBulk) SetDeleted(v bool) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetDeleted(v)
	})
}

// UpdateDeleted sets the "deleted" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateDeleted() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateDeleted()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProjectUpsertBulk) SetUpdatedAt(v time.Time) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateUpdatedAt() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateUpdatedAt()
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
