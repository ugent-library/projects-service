// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/projects/ent/project"
	"github.com/ugent-library/projects/ent/projectidentifier"
)

// ProjectIdentifierCreate is the builder for creating a ProjectIdentifier entity.
type ProjectIdentifierCreate struct {
	config
	mutation *ProjectIdentifierMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetExternalID sets the "external_id" field.
func (pic *ProjectIdentifierCreate) SetExternalID(s string) *ProjectIdentifierCreate {
	pic.mutation.SetExternalID(s)
	return pic
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (pic *ProjectIdentifierCreate) AddProjectIDs(ids ...int) *ProjectIdentifierCreate {
	pic.mutation.AddProjectIDs(ids...)
	return pic
}

// AddProjects adds the "projects" edges to the Project entity.
func (pic *ProjectIdentifierCreate) AddProjects(p ...*Project) *ProjectIdentifierCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pic.AddProjectIDs(ids...)
}

// Mutation returns the ProjectIdentifierMutation object of the builder.
func (pic *ProjectIdentifierCreate) Mutation() *ProjectIdentifierMutation {
	return pic.mutation
}

// Save creates the ProjectIdentifier in the database.
func (pic *ProjectIdentifierCreate) Save(ctx context.Context) (*ProjectIdentifier, error) {
	return withHooks(ctx, pic.sqlSave, pic.mutation, pic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pic *ProjectIdentifierCreate) SaveX(ctx context.Context) *ProjectIdentifier {
	v, err := pic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pic *ProjectIdentifierCreate) Exec(ctx context.Context) error {
	_, err := pic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pic *ProjectIdentifierCreate) ExecX(ctx context.Context) {
	if err := pic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pic *ProjectIdentifierCreate) check() error {
	if _, ok := pic.mutation.ExternalID(); !ok {
		return &ValidationError{Name: "external_id", err: errors.New(`ent: missing required field "ProjectIdentifier.external_id"`)}
	}
	return nil
}

func (pic *ProjectIdentifierCreate) sqlSave(ctx context.Context) (*ProjectIdentifier, error) {
	if err := pic.check(); err != nil {
		return nil, err
	}
	_node, _spec := pic.createSpec()
	if err := sqlgraph.CreateNode(ctx, pic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pic.mutation.id = &_node.ID
	pic.mutation.done = true
	return _node, nil
}

func (pic *ProjectIdentifierCreate) createSpec() (*ProjectIdentifier, *sqlgraph.CreateSpec) {
	var (
		_node = &ProjectIdentifier{config: pic.config}
		_spec = sqlgraph.NewCreateSpec(projectidentifier.Table, sqlgraph.NewFieldSpec(projectidentifier.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pic.conflict
	if value, ok := pic.mutation.ExternalID(); ok {
		_spec.SetField(projectidentifier.FieldExternalID, field.TypeString, value)
		_node.ExternalID = value
	}
	if nodes := pic.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectidentifier.ProjectsTable,
			Columns: []string{projectidentifier.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ProjectIdentifier.Create().
//		SetExternalID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProjectIdentifierUpsert) {
//			SetExternalID(v+v).
//		}).
//		Exec(ctx)
func (pic *ProjectIdentifierCreate) OnConflict(opts ...sql.ConflictOption) *ProjectIdentifierUpsertOne {
	pic.conflict = opts
	return &ProjectIdentifierUpsertOne{
		create: pic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ProjectIdentifier.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pic *ProjectIdentifierCreate) OnConflictColumns(columns ...string) *ProjectIdentifierUpsertOne {
	pic.conflict = append(pic.conflict, sql.ConflictColumns(columns...))
	return &ProjectIdentifierUpsertOne{
		create: pic,
	}
}

type (
	// ProjectIdentifierUpsertOne is the builder for "upsert"-ing
	//  one ProjectIdentifier node.
	ProjectIdentifierUpsertOne struct {
		create *ProjectIdentifierCreate
	}

	// ProjectIdentifierUpsert is the "OnConflict" setter.
	ProjectIdentifierUpsert struct {
		*sql.UpdateSet
	}
)

// SetExternalID sets the "external_id" field.
func (u *ProjectIdentifierUpsert) SetExternalID(v string) *ProjectIdentifierUpsert {
	u.Set(projectidentifier.FieldExternalID, v)
	return u
}

// UpdateExternalID sets the "external_id" field to the value that was provided on create.
func (u *ProjectIdentifierUpsert) UpdateExternalID() *ProjectIdentifierUpsert {
	u.SetExcluded(projectidentifier.FieldExternalID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.ProjectIdentifier.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ProjectIdentifierUpsertOne) UpdateNewValues() *ProjectIdentifierUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ProjectIdentifier.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ProjectIdentifierUpsertOne) Ignore() *ProjectIdentifierUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProjectIdentifierUpsertOne) DoNothing() *ProjectIdentifierUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProjectIdentifierCreate.OnConflict
// documentation for more info.
func (u *ProjectIdentifierUpsertOne) Update(set func(*ProjectIdentifierUpsert)) *ProjectIdentifierUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProjectIdentifierUpsert{UpdateSet: update})
	}))
	return u
}

// SetExternalID sets the "external_id" field.
func (u *ProjectIdentifierUpsertOne) SetExternalID(v string) *ProjectIdentifierUpsertOne {
	return u.Update(func(s *ProjectIdentifierUpsert) {
		s.SetExternalID(v)
	})
}

// UpdateExternalID sets the "external_id" field to the value that was provided on create.
func (u *ProjectIdentifierUpsertOne) UpdateExternalID() *ProjectIdentifierUpsertOne {
	return u.Update(func(s *ProjectIdentifierUpsert) {
		s.UpdateExternalID()
	})
}

// Exec executes the query.
func (u *ProjectIdentifierUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProjectIdentifierCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProjectIdentifierUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ProjectIdentifierUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProjectIdentifierUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ProjectIdentifierCreateBulk is the builder for creating many ProjectIdentifier entities in bulk.
type ProjectIdentifierCreateBulk struct {
	config
	builders []*ProjectIdentifierCreate
	conflict []sql.ConflictOption
}

// Save creates the ProjectIdentifier entities in the database.
func (picb *ProjectIdentifierCreateBulk) Save(ctx context.Context) ([]*ProjectIdentifier, error) {
	specs := make([]*sqlgraph.CreateSpec, len(picb.builders))
	nodes := make([]*ProjectIdentifier, len(picb.builders))
	mutators := make([]Mutator, len(picb.builders))
	for i := range picb.builders {
		func(i int, root context.Context) {
			builder := picb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectIdentifierMutation)
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
					_, err = mutators[i+1].Mutate(root, picb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = picb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, picb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, picb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (picb *ProjectIdentifierCreateBulk) SaveX(ctx context.Context) []*ProjectIdentifier {
	v, err := picb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (picb *ProjectIdentifierCreateBulk) Exec(ctx context.Context) error {
	_, err := picb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (picb *ProjectIdentifierCreateBulk) ExecX(ctx context.Context) {
	if err := picb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ProjectIdentifier.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProjectIdentifierUpsert) {
//			SetExternalID(v+v).
//		}).
//		Exec(ctx)
func (picb *ProjectIdentifierCreateBulk) OnConflict(opts ...sql.ConflictOption) *ProjectIdentifierUpsertBulk {
	picb.conflict = opts
	return &ProjectIdentifierUpsertBulk{
		create: picb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ProjectIdentifier.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (picb *ProjectIdentifierCreateBulk) OnConflictColumns(columns ...string) *ProjectIdentifierUpsertBulk {
	picb.conflict = append(picb.conflict, sql.ConflictColumns(columns...))
	return &ProjectIdentifierUpsertBulk{
		create: picb,
	}
}

// ProjectIdentifierUpsertBulk is the builder for "upsert"-ing
// a bulk of ProjectIdentifier nodes.
type ProjectIdentifierUpsertBulk struct {
	create *ProjectIdentifierCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ProjectIdentifier.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ProjectIdentifierUpsertBulk) UpdateNewValues() *ProjectIdentifierUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ProjectIdentifier.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ProjectIdentifierUpsertBulk) Ignore() *ProjectIdentifierUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProjectIdentifierUpsertBulk) DoNothing() *ProjectIdentifierUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProjectIdentifierCreateBulk.OnConflict
// documentation for more info.
func (u *ProjectIdentifierUpsertBulk) Update(set func(*ProjectIdentifierUpsert)) *ProjectIdentifierUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProjectIdentifierUpsert{UpdateSet: update})
	}))
	return u
}

// SetExternalID sets the "external_id" field.
func (u *ProjectIdentifierUpsertBulk) SetExternalID(v string) *ProjectIdentifierUpsertBulk {
	return u.Update(func(s *ProjectIdentifierUpsert) {
		s.SetExternalID(v)
	})
}

// UpdateExternalID sets the "external_id" field to the value that was provided on create.
func (u *ProjectIdentifierUpsertBulk) UpdateExternalID() *ProjectIdentifierUpsertBulk {
	return u.Update(func(s *ProjectIdentifierUpsert) {
		s.UpdateExternalID()
	})
}

// Exec executes the query.
func (u *ProjectIdentifierUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ProjectIdentifierCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProjectIdentifierCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProjectIdentifierUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
