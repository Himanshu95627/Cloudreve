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
	"github.com/cloudreve/Cloudreve/v4/ent/directlink"
	"github.com/cloudreve/Cloudreve/v4/ent/file"
)

// DirectLinkCreate is the builder for creating a DirectLink entity.
type DirectLinkCreate struct {
	config
	mutation *DirectLinkMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (dlc *DirectLinkCreate) SetCreatedAt(t time.Time) *DirectLinkCreate {
	dlc.mutation.SetCreatedAt(t)
	return dlc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dlc *DirectLinkCreate) SetNillableCreatedAt(t *time.Time) *DirectLinkCreate {
	if t != nil {
		dlc.SetCreatedAt(*t)
	}
	return dlc
}

// SetUpdatedAt sets the "updated_at" field.
func (dlc *DirectLinkCreate) SetUpdatedAt(t time.Time) *DirectLinkCreate {
	dlc.mutation.SetUpdatedAt(t)
	return dlc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dlc *DirectLinkCreate) SetNillableUpdatedAt(t *time.Time) *DirectLinkCreate {
	if t != nil {
		dlc.SetUpdatedAt(*t)
	}
	return dlc
}

// SetDeletedAt sets the "deleted_at" field.
func (dlc *DirectLinkCreate) SetDeletedAt(t time.Time) *DirectLinkCreate {
	dlc.mutation.SetDeletedAt(t)
	return dlc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dlc *DirectLinkCreate) SetNillableDeletedAt(t *time.Time) *DirectLinkCreate {
	if t != nil {
		dlc.SetDeletedAt(*t)
	}
	return dlc
}

// SetName sets the "name" field.
func (dlc *DirectLinkCreate) SetName(s string) *DirectLinkCreate {
	dlc.mutation.SetName(s)
	return dlc
}

// SetDownloads sets the "downloads" field.
func (dlc *DirectLinkCreate) SetDownloads(i int) *DirectLinkCreate {
	dlc.mutation.SetDownloads(i)
	return dlc
}

// SetFileID sets the "file_id" field.
func (dlc *DirectLinkCreate) SetFileID(i int) *DirectLinkCreate {
	dlc.mutation.SetFileID(i)
	return dlc
}

// SetSpeed sets the "speed" field.
func (dlc *DirectLinkCreate) SetSpeed(i int) *DirectLinkCreate {
	dlc.mutation.SetSpeed(i)
	return dlc
}

// SetFile sets the "file" edge to the File entity.
func (dlc *DirectLinkCreate) SetFile(f *File) *DirectLinkCreate {
	return dlc.SetFileID(f.ID)
}

// Mutation returns the DirectLinkMutation object of the builder.
func (dlc *DirectLinkCreate) Mutation() *DirectLinkMutation {
	return dlc.mutation
}

// Save creates the DirectLink in the database.
func (dlc *DirectLinkCreate) Save(ctx context.Context) (*DirectLink, error) {
	if err := dlc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, dlc.sqlSave, dlc.mutation, dlc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dlc *DirectLinkCreate) SaveX(ctx context.Context) *DirectLink {
	v, err := dlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dlc *DirectLinkCreate) Exec(ctx context.Context) error {
	_, err := dlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dlc *DirectLinkCreate) ExecX(ctx context.Context) {
	if err := dlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dlc *DirectLinkCreate) defaults() error {
	if _, ok := dlc.mutation.CreatedAt(); !ok {
		if directlink.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized directlink.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := directlink.DefaultCreatedAt()
		dlc.mutation.SetCreatedAt(v)
	}
	if _, ok := dlc.mutation.UpdatedAt(); !ok {
		if directlink.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized directlink.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := directlink.DefaultUpdatedAt()
		dlc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (dlc *DirectLinkCreate) check() error {
	if _, ok := dlc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DirectLink.created_at"`)}
	}
	if _, ok := dlc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "DirectLink.updated_at"`)}
	}
	if _, ok := dlc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "DirectLink.name"`)}
	}
	if _, ok := dlc.mutation.Downloads(); !ok {
		return &ValidationError{Name: "downloads", err: errors.New(`ent: missing required field "DirectLink.downloads"`)}
	}
	if _, ok := dlc.mutation.FileID(); !ok {
		return &ValidationError{Name: "file_id", err: errors.New(`ent: missing required field "DirectLink.file_id"`)}
	}
	if _, ok := dlc.mutation.Speed(); !ok {
		return &ValidationError{Name: "speed", err: errors.New(`ent: missing required field "DirectLink.speed"`)}
	}
	if _, ok := dlc.mutation.FileID(); !ok {
		return &ValidationError{Name: "file", err: errors.New(`ent: missing required edge "DirectLink.file"`)}
	}
	return nil
}

func (dlc *DirectLinkCreate) sqlSave(ctx context.Context) (*DirectLink, error) {
	if err := dlc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dlc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dlc.mutation.id = &_node.ID
	dlc.mutation.done = true
	return _node, nil
}

func (dlc *DirectLinkCreate) createSpec() (*DirectLink, *sqlgraph.CreateSpec) {
	var (
		_node = &DirectLink{config: dlc.config}
		_spec = sqlgraph.NewCreateSpec(directlink.Table, sqlgraph.NewFieldSpec(directlink.FieldID, field.TypeInt))
	)

	if id, ok := dlc.mutation.ID(); ok {
		_node.ID = id
		id64 := int64(id)
		_spec.ID.Value = id64
	}

	_spec.OnConflict = dlc.conflict
	if value, ok := dlc.mutation.CreatedAt(); ok {
		_spec.SetField(directlink.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dlc.mutation.UpdatedAt(); ok {
		_spec.SetField(directlink.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dlc.mutation.DeletedAt(); ok {
		_spec.SetField(directlink.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := dlc.mutation.Name(); ok {
		_spec.SetField(directlink.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dlc.mutation.Downloads(); ok {
		_spec.SetField(directlink.FieldDownloads, field.TypeInt, value)
		_node.Downloads = value
	}
	if value, ok := dlc.mutation.Speed(); ok {
		_spec.SetField(directlink.FieldSpeed, field.TypeInt, value)
		_node.Speed = value
	}
	if nodes := dlc.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   directlink.FileTable,
			Columns: []string{directlink.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FileID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DirectLink.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DirectLinkUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dlc *DirectLinkCreate) OnConflict(opts ...sql.ConflictOption) *DirectLinkUpsertOne {
	dlc.conflict = opts
	return &DirectLinkUpsertOne{
		create: dlc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DirectLink.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dlc *DirectLinkCreate) OnConflictColumns(columns ...string) *DirectLinkUpsertOne {
	dlc.conflict = append(dlc.conflict, sql.ConflictColumns(columns...))
	return &DirectLinkUpsertOne{
		create: dlc,
	}
}

type (
	// DirectLinkUpsertOne is the builder for "upsert"-ing
	//  one DirectLink node.
	DirectLinkUpsertOne struct {
		create *DirectLinkCreate
	}

	// DirectLinkUpsert is the "OnConflict" setter.
	DirectLinkUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *DirectLinkUpsert) SetUpdatedAt(v time.Time) *DirectLinkUpsert {
	u.Set(directlink.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DirectLinkUpsert) UpdateUpdatedAt() *DirectLinkUpsert {
	u.SetExcluded(directlink.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DirectLinkUpsert) SetDeletedAt(v time.Time) *DirectLinkUpsert {
	u.Set(directlink.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DirectLinkUpsert) UpdateDeletedAt() *DirectLinkUpsert {
	u.SetExcluded(directlink.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *DirectLinkUpsert) ClearDeletedAt() *DirectLinkUpsert {
	u.SetNull(directlink.FieldDeletedAt)
	return u
}

// SetName sets the "name" field.
func (u *DirectLinkUpsert) SetName(v string) *DirectLinkUpsert {
	u.Set(directlink.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DirectLinkUpsert) UpdateName() *DirectLinkUpsert {
	u.SetExcluded(directlink.FieldName)
	return u
}

// SetDownloads sets the "downloads" field.
func (u *DirectLinkUpsert) SetDownloads(v int) *DirectLinkUpsert {
	u.Set(directlink.FieldDownloads, v)
	return u
}

// UpdateDownloads sets the "downloads" field to the value that was provided on create.
func (u *DirectLinkUpsert) UpdateDownloads() *DirectLinkUpsert {
	u.SetExcluded(directlink.FieldDownloads)
	return u
}

// AddDownloads adds v to the "downloads" field.
func (u *DirectLinkUpsert) AddDownloads(v int) *DirectLinkUpsert {
	u.Add(directlink.FieldDownloads, v)
	return u
}

// SetFileID sets the "file_id" field.
func (u *DirectLinkUpsert) SetFileID(v int) *DirectLinkUpsert {
	u.Set(directlink.FieldFileID, v)
	return u
}

// UpdateFileID sets the "file_id" field to the value that was provided on create.
func (u *DirectLinkUpsert) UpdateFileID() *DirectLinkUpsert {
	u.SetExcluded(directlink.FieldFileID)
	return u
}

// SetSpeed sets the "speed" field.
func (u *DirectLinkUpsert) SetSpeed(v int) *DirectLinkUpsert {
	u.Set(directlink.FieldSpeed, v)
	return u
}

// UpdateSpeed sets the "speed" field to the value that was provided on create.
func (u *DirectLinkUpsert) UpdateSpeed() *DirectLinkUpsert {
	u.SetExcluded(directlink.FieldSpeed)
	return u
}

// AddSpeed adds v to the "speed" field.
func (u *DirectLinkUpsert) AddSpeed(v int) *DirectLinkUpsert {
	u.Add(directlink.FieldSpeed, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.DirectLink.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DirectLinkUpsertOne) UpdateNewValues() *DirectLinkUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(directlink.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DirectLink.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DirectLinkUpsertOne) Ignore() *DirectLinkUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DirectLinkUpsertOne) DoNothing() *DirectLinkUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DirectLinkCreate.OnConflict
// documentation for more info.
func (u *DirectLinkUpsertOne) Update(set func(*DirectLinkUpsert)) *DirectLinkUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DirectLinkUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DirectLinkUpsertOne) SetUpdatedAt(v time.Time) *DirectLinkUpsertOne {
	return u.Update(func(s *DirectLinkUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DirectLinkUpsertOne) UpdateUpdatedAt() *DirectLinkUpsertOne {
	return u.Update(func(s *DirectLinkUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DirectLinkUpsertOne) SetDeletedAt(v time.Time) *DirectLinkUpsertOne {
	return u.Update(func(s *DirectLinkUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DirectLinkUpsertOne) UpdateDeletedAt() *DirectLinkUpsertOne {
	return u.Update(func(s *DirectLinkUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *DirectLinkUpsertOne) ClearDeletedAt() *DirectLinkUpsertOne {
	return u.Update(func(s *DirectLinkUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *DirectLinkUpsertOne) SetName(v string) *DirectLinkUpsertOne {
	return u.Update(func(s *DirectLinkUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DirectLinkUpsertOne) UpdateName() *DirectLinkUpsertOne {
	return u.Update(func(s *DirectLinkUpsert) {
		s.UpdateName()
	})
}

// SetDownloads sets the "downloads" field.
func (u *DirectLinkUpsertOne) SetDownloads(v int) *DirectLinkUpsertOne {
	return u.Update(func(s *DirectLinkUpsert) {
		s.SetDownloads(v)
	})
}

// AddDownloads adds v to the "downloads" field.
func (u *DirectLinkUpsertOne) AddDownloads(v int) *DirectLinkUpsertOne {
	return u.Update(func(s *DirectLinkUpsert) {
		s.AddDownloads(v)
	})
}

// UpdateDownloads sets the "downloads" field to the value that was provided on create.
func (u *DirectLinkUpsertOne) UpdateDownloads() *DirectLinkUpsertOne {
	return u.Update(func(s *DirectLinkUpsert) {
		s.UpdateDownloads()
	})
}

// SetFileID sets the "file_id" field.
func (u *DirectLinkUpsertOne) SetFileID(v int) *DirectLinkUpsertOne {
	return u.Update(func(s *DirectLinkUpsert) {
		s.SetFileID(v)
	})
}

// UpdateFileID sets the "file_id" field to the value that was provided on create.
func (u *DirectLinkUpsertOne) UpdateFileID() *DirectLinkUpsertOne {
	return u.Update(func(s *DirectLinkUpsert) {
		s.UpdateFileID()
	})
}

// SetSpeed sets the "speed" field.
func (u *DirectLinkUpsertOne) SetSpeed(v int) *DirectLinkUpsertOne {
	return u.Update(func(s *DirectLinkUpsert) {
		s.SetSpeed(v)
	})
}

// AddSpeed adds v to the "speed" field.
func (u *DirectLinkUpsertOne) AddSpeed(v int) *DirectLinkUpsertOne {
	return u.Update(func(s *DirectLinkUpsert) {
		s.AddSpeed(v)
	})
}

// UpdateSpeed sets the "speed" field to the value that was provided on create.
func (u *DirectLinkUpsertOne) UpdateSpeed() *DirectLinkUpsertOne {
	return u.Update(func(s *DirectLinkUpsert) {
		s.UpdateSpeed()
	})
}

// Exec executes the query.
func (u *DirectLinkUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DirectLinkCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DirectLinkUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DirectLinkUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DirectLinkUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

func (m *DirectLinkCreate) SetRawID(t int) *DirectLinkCreate {
	m.mutation.SetRawID(t)
	return m
}

// DirectLinkCreateBulk is the builder for creating many DirectLink entities in bulk.
type DirectLinkCreateBulk struct {
	config
	err      error
	builders []*DirectLinkCreate
	conflict []sql.ConflictOption
}

// Save creates the DirectLink entities in the database.
func (dlcb *DirectLinkCreateBulk) Save(ctx context.Context) ([]*DirectLink, error) {
	if dlcb.err != nil {
		return nil, dlcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dlcb.builders))
	nodes := make([]*DirectLink, len(dlcb.builders))
	mutators := make([]Mutator, len(dlcb.builders))
	for i := range dlcb.builders {
		func(i int, root context.Context) {
			builder := dlcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DirectLinkMutation)
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
					_, err = mutators[i+1].Mutate(root, dlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dlcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dlcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dlcb *DirectLinkCreateBulk) SaveX(ctx context.Context) []*DirectLink {
	v, err := dlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dlcb *DirectLinkCreateBulk) Exec(ctx context.Context) error {
	_, err := dlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dlcb *DirectLinkCreateBulk) ExecX(ctx context.Context) {
	if err := dlcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DirectLink.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DirectLinkUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dlcb *DirectLinkCreateBulk) OnConflict(opts ...sql.ConflictOption) *DirectLinkUpsertBulk {
	dlcb.conflict = opts
	return &DirectLinkUpsertBulk{
		create: dlcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DirectLink.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dlcb *DirectLinkCreateBulk) OnConflictColumns(columns ...string) *DirectLinkUpsertBulk {
	dlcb.conflict = append(dlcb.conflict, sql.ConflictColumns(columns...))
	return &DirectLinkUpsertBulk{
		create: dlcb,
	}
}

// DirectLinkUpsertBulk is the builder for "upsert"-ing
// a bulk of DirectLink nodes.
type DirectLinkUpsertBulk struct {
	create *DirectLinkCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DirectLink.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DirectLinkUpsertBulk) UpdateNewValues() *DirectLinkUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(directlink.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DirectLink.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DirectLinkUpsertBulk) Ignore() *DirectLinkUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DirectLinkUpsertBulk) DoNothing() *DirectLinkUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DirectLinkCreateBulk.OnConflict
// documentation for more info.
func (u *DirectLinkUpsertBulk) Update(set func(*DirectLinkUpsert)) *DirectLinkUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DirectLinkUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DirectLinkUpsertBulk) SetUpdatedAt(v time.Time) *DirectLinkUpsertBulk {
	return u.Update(func(s *DirectLinkUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DirectLinkUpsertBulk) UpdateUpdatedAt() *DirectLinkUpsertBulk {
	return u.Update(func(s *DirectLinkUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DirectLinkUpsertBulk) SetDeletedAt(v time.Time) *DirectLinkUpsertBulk {
	return u.Update(func(s *DirectLinkUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DirectLinkUpsertBulk) UpdateDeletedAt() *DirectLinkUpsertBulk {
	return u.Update(func(s *DirectLinkUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *DirectLinkUpsertBulk) ClearDeletedAt() *DirectLinkUpsertBulk {
	return u.Update(func(s *DirectLinkUpsert) {
		s.ClearDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *DirectLinkUpsertBulk) SetName(v string) *DirectLinkUpsertBulk {
	return u.Update(func(s *DirectLinkUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DirectLinkUpsertBulk) UpdateName() *DirectLinkUpsertBulk {
	return u.Update(func(s *DirectLinkUpsert) {
		s.UpdateName()
	})
}

// SetDownloads sets the "downloads" field.
func (u *DirectLinkUpsertBulk) SetDownloads(v int) *DirectLinkUpsertBulk {
	return u.Update(func(s *DirectLinkUpsert) {
		s.SetDownloads(v)
	})
}

// AddDownloads adds v to the "downloads" field.
func (u *DirectLinkUpsertBulk) AddDownloads(v int) *DirectLinkUpsertBulk {
	return u.Update(func(s *DirectLinkUpsert) {
		s.AddDownloads(v)
	})
}

// UpdateDownloads sets the "downloads" field to the value that was provided on create.
func (u *DirectLinkUpsertBulk) UpdateDownloads() *DirectLinkUpsertBulk {
	return u.Update(func(s *DirectLinkUpsert) {
		s.UpdateDownloads()
	})
}

// SetFileID sets the "file_id" field.
func (u *DirectLinkUpsertBulk) SetFileID(v int) *DirectLinkUpsertBulk {
	return u.Update(func(s *DirectLinkUpsert) {
		s.SetFileID(v)
	})
}

// UpdateFileID sets the "file_id" field to the value that was provided on create.
func (u *DirectLinkUpsertBulk) UpdateFileID() *DirectLinkUpsertBulk {
	return u.Update(func(s *DirectLinkUpsert) {
		s.UpdateFileID()
	})
}

// SetSpeed sets the "speed" field.
func (u *DirectLinkUpsertBulk) SetSpeed(v int) *DirectLinkUpsertBulk {
	return u.Update(func(s *DirectLinkUpsert) {
		s.SetSpeed(v)
	})
}

// AddSpeed adds v to the "speed" field.
func (u *DirectLinkUpsertBulk) AddSpeed(v int) *DirectLinkUpsertBulk {
	return u.Update(func(s *DirectLinkUpsert) {
		s.AddSpeed(v)
	})
}

// UpdateSpeed sets the "speed" field to the value that was provided on create.
func (u *DirectLinkUpsertBulk) UpdateSpeed() *DirectLinkUpsertBulk {
	return u.Update(func(s *DirectLinkUpsert) {
		s.UpdateSpeed()
	})
}

// Exec executes the query.
func (u *DirectLinkUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DirectLinkCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DirectLinkCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DirectLinkUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
