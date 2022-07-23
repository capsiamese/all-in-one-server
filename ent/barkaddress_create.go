// Code generated by entc, DO NOT EDIT.

package ent

import (
	"aio/ent/barkaddress"
	"aio/ent/extensionclient"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BarkAddressCreate is the builder for creating a BarkAddress entity.
type BarkAddressCreate struct {
	config
	mutation *BarkAddressMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (bac *BarkAddressCreate) SetName(s string) *BarkAddressCreate {
	bac.mutation.SetName(s)
	return bac
}

// SetTarget sets the "target" field.
func (bac *BarkAddressCreate) SetTarget(s string) *BarkAddressCreate {
	bac.mutation.SetTarget(s)
	return bac
}

// SetIndex sets the "index" field.
func (bac *BarkAddressCreate) SetIndex(i int64) *BarkAddressCreate {
	bac.mutation.SetIndex(i)
	return bac
}

// AddClientIDs adds the "client" edge to the ExtensionClient entity by IDs.
func (bac *BarkAddressCreate) AddClientIDs(ids ...int) *BarkAddressCreate {
	bac.mutation.AddClientIDs(ids...)
	return bac
}

// AddClient adds the "client" edges to the ExtensionClient entity.
func (bac *BarkAddressCreate) AddClient(e ...*ExtensionClient) *BarkAddressCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return bac.AddClientIDs(ids...)
}

// Mutation returns the BarkAddressMutation object of the builder.
func (bac *BarkAddressCreate) Mutation() *BarkAddressMutation {
	return bac.mutation
}

// Save creates the BarkAddress in the database.
func (bac *BarkAddressCreate) Save(ctx context.Context) (*BarkAddress, error) {
	var (
		err  error
		node *BarkAddress
	)
	if len(bac.hooks) == 0 {
		if err = bac.check(); err != nil {
			return nil, err
		}
		node, err = bac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BarkAddressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bac.check(); err != nil {
				return nil, err
			}
			bac.mutation = mutation
			if node, err = bac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bac.hooks) - 1; i >= 0; i-- {
			if bac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bac *BarkAddressCreate) SaveX(ctx context.Context) *BarkAddress {
	v, err := bac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bac *BarkAddressCreate) Exec(ctx context.Context) error {
	_, err := bac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bac *BarkAddressCreate) ExecX(ctx context.Context) {
	if err := bac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bac *BarkAddressCreate) check() error {
	if _, ok := bac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "BarkAddress.name"`)}
	}
	if v, ok := bac.mutation.Name(); ok {
		if err := barkaddress.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "BarkAddress.name": %w`, err)}
		}
	}
	if _, ok := bac.mutation.Target(); !ok {
		return &ValidationError{Name: "target", err: errors.New(`ent: missing required field "BarkAddress.target"`)}
	}
	if v, ok := bac.mutation.Target(); ok {
		if err := barkaddress.TargetValidator(v); err != nil {
			return &ValidationError{Name: "target", err: fmt.Errorf(`ent: validator failed for field "BarkAddress.target": %w`, err)}
		}
	}
	if _, ok := bac.mutation.Index(); !ok {
		return &ValidationError{Name: "index", err: errors.New(`ent: missing required field "BarkAddress.index"`)}
	}
	return nil
}

func (bac *BarkAddressCreate) sqlSave(ctx context.Context) (*BarkAddress, error) {
	_node, _spec := bac.createSpec()
	if err := sqlgraph.CreateNode(ctx, bac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bac *BarkAddressCreate) createSpec() (*BarkAddress, *sqlgraph.CreateSpec) {
	var (
		_node = &BarkAddress{config: bac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: barkaddress.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: barkaddress.FieldID,
			},
		}
	)
	if value, ok := bac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: barkaddress.FieldName,
		})
		_node.Name = value
	}
	if value, ok := bac.mutation.Target(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: barkaddress.FieldTarget,
		})
		_node.Target = value
	}
	if value, ok := bac.mutation.Index(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: barkaddress.FieldIndex,
		})
		_node.Index = value
	}
	if nodes := bac.mutation.ClientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   barkaddress.ClientTable,
			Columns: barkaddress.ClientPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: extensionclient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BarkAddressCreateBulk is the builder for creating many BarkAddress entities in bulk.
type BarkAddressCreateBulk struct {
	config
	builders []*BarkAddressCreate
}

// Save creates the BarkAddress entities in the database.
func (bacb *BarkAddressCreateBulk) Save(ctx context.Context) ([]*BarkAddress, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bacb.builders))
	nodes := make([]*BarkAddress, len(bacb.builders))
	mutators := make([]Mutator, len(bacb.builders))
	for i := range bacb.builders {
		func(i int, root context.Context) {
			builder := bacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BarkAddressMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bacb *BarkAddressCreateBulk) SaveX(ctx context.Context) []*BarkAddress {
	v, err := bacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bacb *BarkAddressCreateBulk) Exec(ctx context.Context) error {
	_, err := bacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bacb *BarkAddressCreateBulk) ExecX(ctx context.Context) {
	if err := bacb.Exec(ctx); err != nil {
		panic(err)
	}
}
