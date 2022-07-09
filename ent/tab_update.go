// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"notification/ent/group"
	"notification/ent/predicate"
	"notification/ent/tab"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TabUpdate is the builder for updating Tab entities.
type TabUpdate struct {
	config
	hooks    []Hook
	mutation *TabMutation
}

// Where appends a list predicates to the TabUpdate builder.
func (tu *TabUpdate) Where(ps ...predicate.Tab) *TabUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TabUpdate) SetName(s string) *TabUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetURL sets the "url" field.
func (tu *TabUpdate) SetURL(s string) *TabUpdate {
	tu.mutation.SetURL(s)
	return tu
}

// SetSeq sets the "seq" field.
func (tu *TabUpdate) SetSeq(i int32) *TabUpdate {
	tu.mutation.ResetSeq()
	tu.mutation.SetSeq(i)
	return tu
}

// AddSeq adds i to the "seq" field.
func (tu *TabUpdate) AddSeq(i int32) *TabUpdate {
	tu.mutation.AddSeq(i)
	return tu
}

// SetFavicon sets the "favicon" field.
func (tu *TabUpdate) SetFavicon(s string) *TabUpdate {
	tu.mutation.SetFavicon(s)
	return tu
}

// SetNillableFavicon sets the "favicon" field if the given value is not nil.
func (tu *TabUpdate) SetNillableFavicon(s *string) *TabUpdate {
	if s != nil {
		tu.SetFavicon(*s)
	}
	return tu
}

// ClearFavicon clears the value of the "favicon" field.
func (tu *TabUpdate) ClearFavicon() *TabUpdate {
	tu.mutation.ClearFavicon()
	return tu
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (tu *TabUpdate) SetGroupID(id int) *TabUpdate {
	tu.mutation.SetGroupID(id)
	return tu
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (tu *TabUpdate) SetNillableGroupID(id *int) *TabUpdate {
	if id != nil {
		tu = tu.SetGroupID(*id)
	}
	return tu
}

// SetGroup sets the "group" edge to the Group entity.
func (tu *TabUpdate) SetGroup(g *Group) *TabUpdate {
	return tu.SetGroupID(g.ID)
}

// Mutation returns the TabMutation object of the builder.
func (tu *TabUpdate) Mutation() *TabMutation {
	return tu.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (tu *TabUpdate) ClearGroup() *TabUpdate {
	tu.mutation.ClearGroup()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TabUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TabMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TabUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TabUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TabUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TabUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tab.Table,
			Columns: tab.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tab.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tab.FieldName,
		})
	}
	if value, ok := tu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tab.FieldURL,
		})
	}
	if value, ok := tu.mutation.Seq(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: tab.FieldSeq,
		})
	}
	if value, ok := tu.mutation.AddedSeq(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: tab.FieldSeq,
		})
	}
	if value, ok := tu.mutation.Favicon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tab.FieldFavicon,
		})
	}
	if tu.mutation.FaviconCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tab.FieldFavicon,
		})
	}
	if tu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tab.GroupTable,
			Columns: []string{tab.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tab.GroupTable,
			Columns: []string{tab.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tab.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TabUpdateOne is the builder for updating a single Tab entity.
type TabUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TabMutation
}

// SetName sets the "name" field.
func (tuo *TabUpdateOne) SetName(s string) *TabUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetURL sets the "url" field.
func (tuo *TabUpdateOne) SetURL(s string) *TabUpdateOne {
	tuo.mutation.SetURL(s)
	return tuo
}

// SetSeq sets the "seq" field.
func (tuo *TabUpdateOne) SetSeq(i int32) *TabUpdateOne {
	tuo.mutation.ResetSeq()
	tuo.mutation.SetSeq(i)
	return tuo
}

// AddSeq adds i to the "seq" field.
func (tuo *TabUpdateOne) AddSeq(i int32) *TabUpdateOne {
	tuo.mutation.AddSeq(i)
	return tuo
}

// SetFavicon sets the "favicon" field.
func (tuo *TabUpdateOne) SetFavicon(s string) *TabUpdateOne {
	tuo.mutation.SetFavicon(s)
	return tuo
}

// SetNillableFavicon sets the "favicon" field if the given value is not nil.
func (tuo *TabUpdateOne) SetNillableFavicon(s *string) *TabUpdateOne {
	if s != nil {
		tuo.SetFavicon(*s)
	}
	return tuo
}

// ClearFavicon clears the value of the "favicon" field.
func (tuo *TabUpdateOne) ClearFavicon() *TabUpdateOne {
	tuo.mutation.ClearFavicon()
	return tuo
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (tuo *TabUpdateOne) SetGroupID(id int) *TabUpdateOne {
	tuo.mutation.SetGroupID(id)
	return tuo
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (tuo *TabUpdateOne) SetNillableGroupID(id *int) *TabUpdateOne {
	if id != nil {
		tuo = tuo.SetGroupID(*id)
	}
	return tuo
}

// SetGroup sets the "group" edge to the Group entity.
func (tuo *TabUpdateOne) SetGroup(g *Group) *TabUpdateOne {
	return tuo.SetGroupID(g.ID)
}

// Mutation returns the TabMutation object of the builder.
func (tuo *TabUpdateOne) Mutation() *TabMutation {
	return tuo.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (tuo *TabUpdateOne) ClearGroup() *TabUpdateOne {
	tuo.mutation.ClearGroup()
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TabUpdateOne) Select(field string, fields ...string) *TabUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tab entity.
func (tuo *TabUpdateOne) Save(ctx context.Context) (*Tab, error) {
	var (
		err  error
		node *Tab
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TabMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TabUpdateOne) SaveX(ctx context.Context) *Tab {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TabUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TabUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TabUpdateOne) sqlSave(ctx context.Context) (_node *Tab, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tab.Table,
			Columns: tab.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tab.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tab.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tab.FieldID)
		for _, f := range fields {
			if !tab.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tab.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tab.FieldName,
		})
	}
	if value, ok := tuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tab.FieldURL,
		})
	}
	if value, ok := tuo.mutation.Seq(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: tab.FieldSeq,
		})
	}
	if value, ok := tuo.mutation.AddedSeq(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: tab.FieldSeq,
		})
	}
	if value, ok := tuo.mutation.Favicon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tab.FieldFavicon,
		})
	}
	if tuo.mutation.FaviconCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tab.FieldFavicon,
		})
	}
	if tuo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tab.GroupTable,
			Columns: []string{tab.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tab.GroupTable,
			Columns: []string{tab.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tab{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tab.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
