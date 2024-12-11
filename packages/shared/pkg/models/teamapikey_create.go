// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/team"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/teamapikey"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/user"
	"github.com/google/uuid"
)

// TeamAPIKeyCreate is the builder for creating a TeamAPIKey entity.
type TeamAPIKeyCreate struct {
	config
	mutation *TeamAPIKeyMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAPIKey sets the "api_key" field.
func (takc *TeamAPIKeyCreate) SetAPIKey(s string) *TeamAPIKeyCreate {
	takc.mutation.SetAPIKey(s)
	return takc
}

// SetCreatedAt sets the "created_at" field.
func (takc *TeamAPIKeyCreate) SetCreatedAt(t time.Time) *TeamAPIKeyCreate {
	takc.mutation.SetCreatedAt(t)
	return takc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (takc *TeamAPIKeyCreate) SetNillableCreatedAt(t *time.Time) *TeamAPIKeyCreate {
	if t != nil {
		takc.SetCreatedAt(*t)
	}
	return takc
}

// SetUpdatedAt sets the "updated_at" field.
func (takc *TeamAPIKeyCreate) SetUpdatedAt(t time.Time) *TeamAPIKeyCreate {
	takc.mutation.SetUpdatedAt(t)
	return takc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (takc *TeamAPIKeyCreate) SetNillableUpdatedAt(t *time.Time) *TeamAPIKeyCreate {
	if t != nil {
		takc.SetUpdatedAt(*t)
	}
	return takc
}

// SetTeamID sets the "team_id" field.
func (takc *TeamAPIKeyCreate) SetTeamID(u uuid.UUID) *TeamAPIKeyCreate {
	takc.mutation.SetTeamID(u)
	return takc
}

// SetName sets the "name" field.
func (takc *TeamAPIKeyCreate) SetName(s string) *TeamAPIKeyCreate {
	takc.mutation.SetName(s)
	return takc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (takc *TeamAPIKeyCreate) SetNillableName(s *string) *TeamAPIKeyCreate {
	if s != nil {
		takc.SetName(*s)
	}
	return takc
}

// SetCreatedBy sets the "created_by" field.
func (takc *TeamAPIKeyCreate) SetCreatedBy(u uuid.UUID) *TeamAPIKeyCreate {
	takc.mutation.SetCreatedBy(u)
	return takc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (takc *TeamAPIKeyCreate) SetNillableCreatedBy(u *uuid.UUID) *TeamAPIKeyCreate {
	if u != nil {
		takc.SetCreatedBy(*u)
	}
	return takc
}

// SetLastUsed sets the "last_used" field.
func (takc *TeamAPIKeyCreate) SetLastUsed(t time.Time) *TeamAPIKeyCreate {
	takc.mutation.SetLastUsed(t)
	return takc
}

// SetNillableLastUsed sets the "last_used" field if the given value is not nil.
func (takc *TeamAPIKeyCreate) SetNillableLastUsed(t *time.Time) *TeamAPIKeyCreate {
	if t != nil {
		takc.SetLastUsed(*t)
	}
	return takc
}

// SetID sets the "id" field.
func (takc *TeamAPIKeyCreate) SetID(u uuid.UUID) *TeamAPIKeyCreate {
	takc.mutation.SetID(u)
	return takc
}

// SetTeam sets the "team" edge to the Team entity.
func (takc *TeamAPIKeyCreate) SetTeam(t *Team) *TeamAPIKeyCreate {
	return takc.SetTeamID(t.ID)
}

// SetCreatorID sets the "creator" edge to the User entity by ID.
func (takc *TeamAPIKeyCreate) SetCreatorID(id uuid.UUID) *TeamAPIKeyCreate {
	takc.mutation.SetCreatorID(id)
	return takc
}

// SetNillableCreatorID sets the "creator" edge to the User entity by ID if the given value is not nil.
func (takc *TeamAPIKeyCreate) SetNillableCreatorID(id *uuid.UUID) *TeamAPIKeyCreate {
	if id != nil {
		takc = takc.SetCreatorID(*id)
	}
	return takc
}

// SetCreator sets the "creator" edge to the User entity.
func (takc *TeamAPIKeyCreate) SetCreator(u *User) *TeamAPIKeyCreate {
	return takc.SetCreatorID(u.ID)
}

// Mutation returns the TeamAPIKeyMutation object of the builder.
func (takc *TeamAPIKeyCreate) Mutation() *TeamAPIKeyMutation {
	return takc.mutation
}

// Save creates the TeamAPIKey in the database.
func (takc *TeamAPIKeyCreate) Save(ctx context.Context) (*TeamAPIKey, error) {
	takc.defaults()
	return withHooks(ctx, takc.sqlSave, takc.mutation, takc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (takc *TeamAPIKeyCreate) SaveX(ctx context.Context) *TeamAPIKey {
	v, err := takc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (takc *TeamAPIKeyCreate) Exec(ctx context.Context) error {
	_, err := takc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (takc *TeamAPIKeyCreate) ExecX(ctx context.Context) {
	if err := takc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (takc *TeamAPIKeyCreate) defaults() {
	if _, ok := takc.mutation.CreatedAt(); !ok {
		v := teamapikey.DefaultCreatedAt()
		takc.mutation.SetCreatedAt(v)
	}
	if _, ok := takc.mutation.Name(); !ok {
		v := teamapikey.DefaultName
		takc.mutation.SetName(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (takc *TeamAPIKeyCreate) check() error {
	if _, ok := takc.mutation.APIKey(); !ok {
		return &ValidationError{Name: "api_key", err: errors.New(`models: missing required field "TeamAPIKey.api_key"`)}
	}
	if _, ok := takc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`models: missing required field "TeamAPIKey.created_at"`)}
	}
	if _, ok := takc.mutation.TeamID(); !ok {
		return &ValidationError{Name: "team_id", err: errors.New(`models: missing required field "TeamAPIKey.team_id"`)}
	}
	if _, ok := takc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`models: missing required field "TeamAPIKey.name"`)}
	}
	if _, ok := takc.mutation.TeamID(); !ok {
		return &ValidationError{Name: "team", err: errors.New(`models: missing required edge "TeamAPIKey.team"`)}
	}
	return nil
}

func (takc *TeamAPIKeyCreate) sqlSave(ctx context.Context) (*TeamAPIKey, error) {
	if err := takc.check(); err != nil {
		return nil, err
	}
	_node, _spec := takc.createSpec()
	if err := sqlgraph.CreateNode(ctx, takc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	takc.mutation.id = &_node.ID
	takc.mutation.done = true
	return _node, nil
}

func (takc *TeamAPIKeyCreate) createSpec() (*TeamAPIKey, *sqlgraph.CreateSpec) {
	var (
		_node = &TeamAPIKey{config: takc.config}
		_spec = sqlgraph.NewCreateSpec(teamapikey.Table, sqlgraph.NewFieldSpec(teamapikey.FieldID, field.TypeUUID))
	)
	_spec.Schema = takc.schemaConfig.TeamAPIKey
	_spec.OnConflict = takc.conflict
	if id, ok := takc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := takc.mutation.APIKey(); ok {
		_spec.SetField(teamapikey.FieldAPIKey, field.TypeString, value)
		_node.APIKey = value
	}
	if value, ok := takc.mutation.CreatedAt(); ok {
		_spec.SetField(teamapikey.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := takc.mutation.UpdatedAt(); ok {
		_spec.SetField(teamapikey.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := takc.mutation.Name(); ok {
		_spec.SetField(teamapikey.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := takc.mutation.LastUsed(); ok {
		_spec.SetField(teamapikey.FieldLastUsed, field.TypeTime, value)
		_node.LastUsed = &value
	}
	if nodes := takc.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teamapikey.TeamTable,
			Columns: []string{teamapikey.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = takc.schemaConfig.TeamAPIKey
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TeamID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := takc.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   teamapikey.CreatorTable,
			Columns: []string{teamapikey.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = takc.schemaConfig.TeamAPIKey
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CreatedBy = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TeamAPIKey.Create().
//		SetAPIKey(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TeamAPIKeyUpsert) {
//			SetAPIKey(v+v).
//		}).
//		Exec(ctx)
func (takc *TeamAPIKeyCreate) OnConflict(opts ...sql.ConflictOption) *TeamAPIKeyUpsertOne {
	takc.conflict = opts
	return &TeamAPIKeyUpsertOne{
		create: takc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TeamAPIKey.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (takc *TeamAPIKeyCreate) OnConflictColumns(columns ...string) *TeamAPIKeyUpsertOne {
	takc.conflict = append(takc.conflict, sql.ConflictColumns(columns...))
	return &TeamAPIKeyUpsertOne{
		create: takc,
	}
}

type (
	// TeamAPIKeyUpsertOne is the builder for "upsert"-ing
	//  one TeamAPIKey node.
	TeamAPIKeyUpsertOne struct {
		create *TeamAPIKeyCreate
	}

	// TeamAPIKeyUpsert is the "OnConflict" setter.
	TeamAPIKeyUpsert struct {
		*sql.UpdateSet
	}
)

// SetAPIKey sets the "api_key" field.
func (u *TeamAPIKeyUpsert) SetAPIKey(v string) *TeamAPIKeyUpsert {
	u.Set(teamapikey.FieldAPIKey, v)
	return u
}

// UpdateAPIKey sets the "api_key" field to the value that was provided on create.
func (u *TeamAPIKeyUpsert) UpdateAPIKey() *TeamAPIKeyUpsert {
	u.SetExcluded(teamapikey.FieldAPIKey)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TeamAPIKeyUpsert) SetUpdatedAt(v time.Time) *TeamAPIKeyUpsert {
	u.Set(teamapikey.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TeamAPIKeyUpsert) UpdateUpdatedAt() *TeamAPIKeyUpsert {
	u.SetExcluded(teamapikey.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *TeamAPIKeyUpsert) ClearUpdatedAt() *TeamAPIKeyUpsert {
	u.SetNull(teamapikey.FieldUpdatedAt)
	return u
}

// SetTeamID sets the "team_id" field.
func (u *TeamAPIKeyUpsert) SetTeamID(v uuid.UUID) *TeamAPIKeyUpsert {
	u.Set(teamapikey.FieldTeamID, v)
	return u
}

// UpdateTeamID sets the "team_id" field to the value that was provided on create.
func (u *TeamAPIKeyUpsert) UpdateTeamID() *TeamAPIKeyUpsert {
	u.SetExcluded(teamapikey.FieldTeamID)
	return u
}

// SetName sets the "name" field.
func (u *TeamAPIKeyUpsert) SetName(v string) *TeamAPIKeyUpsert {
	u.Set(teamapikey.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TeamAPIKeyUpsert) UpdateName() *TeamAPIKeyUpsert {
	u.SetExcluded(teamapikey.FieldName)
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *TeamAPIKeyUpsert) SetCreatedBy(v uuid.UUID) *TeamAPIKeyUpsert {
	u.Set(teamapikey.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *TeamAPIKeyUpsert) UpdateCreatedBy() *TeamAPIKeyUpsert {
	u.SetExcluded(teamapikey.FieldCreatedBy)
	return u
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *TeamAPIKeyUpsert) ClearCreatedBy() *TeamAPIKeyUpsert {
	u.SetNull(teamapikey.FieldCreatedBy)
	return u
}

// SetLastUsed sets the "last_used" field.
func (u *TeamAPIKeyUpsert) SetLastUsed(v time.Time) *TeamAPIKeyUpsert {
	u.Set(teamapikey.FieldLastUsed, v)
	return u
}

// UpdateLastUsed sets the "last_used" field to the value that was provided on create.
func (u *TeamAPIKeyUpsert) UpdateLastUsed() *TeamAPIKeyUpsert {
	u.SetExcluded(teamapikey.FieldLastUsed)
	return u
}

// ClearLastUsed clears the value of the "last_used" field.
func (u *TeamAPIKeyUpsert) ClearLastUsed() *TeamAPIKeyUpsert {
	u.SetNull(teamapikey.FieldLastUsed)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TeamAPIKey.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(teamapikey.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TeamAPIKeyUpsertOne) UpdateNewValues() *TeamAPIKeyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(teamapikey.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(teamapikey.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TeamAPIKey.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TeamAPIKeyUpsertOne) Ignore() *TeamAPIKeyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TeamAPIKeyUpsertOne) DoNothing() *TeamAPIKeyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TeamAPIKeyCreate.OnConflict
// documentation for more info.
func (u *TeamAPIKeyUpsertOne) Update(set func(*TeamAPIKeyUpsert)) *TeamAPIKeyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TeamAPIKeyUpsert{UpdateSet: update})
	}))
	return u
}

// SetAPIKey sets the "api_key" field.
func (u *TeamAPIKeyUpsertOne) SetAPIKey(v string) *TeamAPIKeyUpsertOne {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.SetAPIKey(v)
	})
}

// UpdateAPIKey sets the "api_key" field to the value that was provided on create.
func (u *TeamAPIKeyUpsertOne) UpdateAPIKey() *TeamAPIKeyUpsertOne {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.UpdateAPIKey()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TeamAPIKeyUpsertOne) SetUpdatedAt(v time.Time) *TeamAPIKeyUpsertOne {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TeamAPIKeyUpsertOne) UpdateUpdatedAt() *TeamAPIKeyUpsertOne {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *TeamAPIKeyUpsertOne) ClearUpdatedAt() *TeamAPIKeyUpsertOne {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetTeamID sets the "team_id" field.
func (u *TeamAPIKeyUpsertOne) SetTeamID(v uuid.UUID) *TeamAPIKeyUpsertOne {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.SetTeamID(v)
	})
}

// UpdateTeamID sets the "team_id" field to the value that was provided on create.
func (u *TeamAPIKeyUpsertOne) UpdateTeamID() *TeamAPIKeyUpsertOne {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.UpdateTeamID()
	})
}

// SetName sets the "name" field.
func (u *TeamAPIKeyUpsertOne) SetName(v string) *TeamAPIKeyUpsertOne {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TeamAPIKeyUpsertOne) UpdateName() *TeamAPIKeyUpsertOne {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.UpdateName()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *TeamAPIKeyUpsertOne) SetCreatedBy(v uuid.UUID) *TeamAPIKeyUpsertOne {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *TeamAPIKeyUpsertOne) UpdateCreatedBy() *TeamAPIKeyUpsertOne {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *TeamAPIKeyUpsertOne) ClearCreatedBy() *TeamAPIKeyUpsertOne {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.ClearCreatedBy()
	})
}

// SetLastUsed sets the "last_used" field.
func (u *TeamAPIKeyUpsertOne) SetLastUsed(v time.Time) *TeamAPIKeyUpsertOne {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.SetLastUsed(v)
	})
}

// UpdateLastUsed sets the "last_used" field to the value that was provided on create.
func (u *TeamAPIKeyUpsertOne) UpdateLastUsed() *TeamAPIKeyUpsertOne {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.UpdateLastUsed()
	})
}

// ClearLastUsed clears the value of the "last_used" field.
func (u *TeamAPIKeyUpsertOne) ClearLastUsed() *TeamAPIKeyUpsertOne {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.ClearLastUsed()
	})
}

// Exec executes the query.
func (u *TeamAPIKeyUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("models: missing options for TeamAPIKeyCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TeamAPIKeyUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TeamAPIKeyUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("models: TeamAPIKeyUpsertOne.ID is not supported by MySQL driver. Use TeamAPIKeyUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TeamAPIKeyUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TeamAPIKeyCreateBulk is the builder for creating many TeamAPIKey entities in bulk.
type TeamAPIKeyCreateBulk struct {
	config
	err      error
	builders []*TeamAPIKeyCreate
	conflict []sql.ConflictOption
}

// Save creates the TeamAPIKey entities in the database.
func (takcb *TeamAPIKeyCreateBulk) Save(ctx context.Context) ([]*TeamAPIKey, error) {
	if takcb.err != nil {
		return nil, takcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(takcb.builders))
	nodes := make([]*TeamAPIKey, len(takcb.builders))
	mutators := make([]Mutator, len(takcb.builders))
	for i := range takcb.builders {
		func(i int, root context.Context) {
			builder := takcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TeamAPIKeyMutation)
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
					_, err = mutators[i+1].Mutate(root, takcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = takcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, takcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, takcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (takcb *TeamAPIKeyCreateBulk) SaveX(ctx context.Context) []*TeamAPIKey {
	v, err := takcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (takcb *TeamAPIKeyCreateBulk) Exec(ctx context.Context) error {
	_, err := takcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (takcb *TeamAPIKeyCreateBulk) ExecX(ctx context.Context) {
	if err := takcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TeamAPIKey.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TeamAPIKeyUpsert) {
//			SetAPIKey(v+v).
//		}).
//		Exec(ctx)
func (takcb *TeamAPIKeyCreateBulk) OnConflict(opts ...sql.ConflictOption) *TeamAPIKeyUpsertBulk {
	takcb.conflict = opts
	return &TeamAPIKeyUpsertBulk{
		create: takcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TeamAPIKey.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (takcb *TeamAPIKeyCreateBulk) OnConflictColumns(columns ...string) *TeamAPIKeyUpsertBulk {
	takcb.conflict = append(takcb.conflict, sql.ConflictColumns(columns...))
	return &TeamAPIKeyUpsertBulk{
		create: takcb,
	}
}

// TeamAPIKeyUpsertBulk is the builder for "upsert"-ing
// a bulk of TeamAPIKey nodes.
type TeamAPIKeyUpsertBulk struct {
	create *TeamAPIKeyCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TeamAPIKey.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(teamapikey.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TeamAPIKeyUpsertBulk) UpdateNewValues() *TeamAPIKeyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(teamapikey.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(teamapikey.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TeamAPIKey.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TeamAPIKeyUpsertBulk) Ignore() *TeamAPIKeyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TeamAPIKeyUpsertBulk) DoNothing() *TeamAPIKeyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TeamAPIKeyCreateBulk.OnConflict
// documentation for more info.
func (u *TeamAPIKeyUpsertBulk) Update(set func(*TeamAPIKeyUpsert)) *TeamAPIKeyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TeamAPIKeyUpsert{UpdateSet: update})
	}))
	return u
}

// SetAPIKey sets the "api_key" field.
func (u *TeamAPIKeyUpsertBulk) SetAPIKey(v string) *TeamAPIKeyUpsertBulk {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.SetAPIKey(v)
	})
}

// UpdateAPIKey sets the "api_key" field to the value that was provided on create.
func (u *TeamAPIKeyUpsertBulk) UpdateAPIKey() *TeamAPIKeyUpsertBulk {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.UpdateAPIKey()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TeamAPIKeyUpsertBulk) SetUpdatedAt(v time.Time) *TeamAPIKeyUpsertBulk {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TeamAPIKeyUpsertBulk) UpdateUpdatedAt() *TeamAPIKeyUpsertBulk {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *TeamAPIKeyUpsertBulk) ClearUpdatedAt() *TeamAPIKeyUpsertBulk {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetTeamID sets the "team_id" field.
func (u *TeamAPIKeyUpsertBulk) SetTeamID(v uuid.UUID) *TeamAPIKeyUpsertBulk {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.SetTeamID(v)
	})
}

// UpdateTeamID sets the "team_id" field to the value that was provided on create.
func (u *TeamAPIKeyUpsertBulk) UpdateTeamID() *TeamAPIKeyUpsertBulk {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.UpdateTeamID()
	})
}

// SetName sets the "name" field.
func (u *TeamAPIKeyUpsertBulk) SetName(v string) *TeamAPIKeyUpsertBulk {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TeamAPIKeyUpsertBulk) UpdateName() *TeamAPIKeyUpsertBulk {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.UpdateName()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *TeamAPIKeyUpsertBulk) SetCreatedBy(v uuid.UUID) *TeamAPIKeyUpsertBulk {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *TeamAPIKeyUpsertBulk) UpdateCreatedBy() *TeamAPIKeyUpsertBulk {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *TeamAPIKeyUpsertBulk) ClearCreatedBy() *TeamAPIKeyUpsertBulk {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.ClearCreatedBy()
	})
}

// SetLastUsed sets the "last_used" field.
func (u *TeamAPIKeyUpsertBulk) SetLastUsed(v time.Time) *TeamAPIKeyUpsertBulk {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.SetLastUsed(v)
	})
}

// UpdateLastUsed sets the "last_used" field to the value that was provided on create.
func (u *TeamAPIKeyUpsertBulk) UpdateLastUsed() *TeamAPIKeyUpsertBulk {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.UpdateLastUsed()
	})
}

// ClearLastUsed clears the value of the "last_used" field.
func (u *TeamAPIKeyUpsertBulk) ClearLastUsed() *TeamAPIKeyUpsertBulk {
	return u.Update(func(s *TeamAPIKeyUpsert) {
		s.ClearLastUsed()
	})
}

// Exec executes the query.
func (u *TeamAPIKeyUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("models: OnConflict was set for builder %d. Set it on the TeamAPIKeyCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("models: missing options for TeamAPIKeyCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TeamAPIKeyUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
