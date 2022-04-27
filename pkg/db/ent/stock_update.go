// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/stock-manager/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/stock-manager/pkg/db/ent/stock"
	"github.com/google/uuid"
)

// StockUpdate is the builder for updating Stock entities.
type StockUpdate struct {
	config
	hooks    []Hook
	mutation *StockMutation
}

// Where appends a list predicates to the StockUpdate builder.
func (su *StockUpdate) Where(ps ...predicate.Stock) *StockUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetGoodID sets the "good_id" field.
func (su *StockUpdate) SetGoodID(u uuid.UUID) *StockUpdate {
	su.mutation.SetGoodID(u)
	return su
}

// SetTotal sets the "total" field.
func (su *StockUpdate) SetTotal(i int32) *StockUpdate {
	su.mutation.ResetTotal()
	su.mutation.SetTotal(i)
	return su
}

// AddTotal adds i to the "total" field.
func (su *StockUpdate) AddTotal(i int32) *StockUpdate {
	su.mutation.AddTotal(i)
	return su
}

// SetInService sets the "in_service" field.
func (su *StockUpdate) SetInService(i int32) *StockUpdate {
	su.mutation.ResetInService()
	su.mutation.SetInService(i)
	return su
}

// AddInService adds i to the "in_service" field.
func (su *StockUpdate) AddInService(i int32) *StockUpdate {
	su.mutation.AddInService(i)
	return su
}

// SetSold sets the "sold" field.
func (su *StockUpdate) SetSold(i int32) *StockUpdate {
	su.mutation.ResetSold()
	su.mutation.SetSold(i)
	return su
}

// AddSold adds i to the "sold" field.
func (su *StockUpdate) AddSold(i int32) *StockUpdate {
	su.mutation.AddSold(i)
	return su
}

// Mutation returns the StockMutation object of the builder.
func (su *StockUpdate) Mutation() *StockMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StockUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *StockUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StockUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StockUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *StockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   stock.Table,
			Columns: stock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: stock.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: stock.FieldGoodID,
		})
	}
	if value, ok := su.mutation.Total(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldTotal,
		})
	}
	if value, ok := su.mutation.AddedTotal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldTotal,
		})
	}
	if value, ok := su.mutation.InService(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldInService,
		})
	}
	if value, ok := su.mutation.AddedInService(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldInService,
		})
	}
	if value, ok := su.mutation.Sold(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldSold,
		})
	}
	if value, ok := su.mutation.AddedSold(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldSold,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{stock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// StockUpdateOne is the builder for updating a single Stock entity.
type StockUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StockMutation
}

// SetGoodID sets the "good_id" field.
func (suo *StockUpdateOne) SetGoodID(u uuid.UUID) *StockUpdateOne {
	suo.mutation.SetGoodID(u)
	return suo
}

// SetTotal sets the "total" field.
func (suo *StockUpdateOne) SetTotal(i int32) *StockUpdateOne {
	suo.mutation.ResetTotal()
	suo.mutation.SetTotal(i)
	return suo
}

// AddTotal adds i to the "total" field.
func (suo *StockUpdateOne) AddTotal(i int32) *StockUpdateOne {
	suo.mutation.AddTotal(i)
	return suo
}

// SetInService sets the "in_service" field.
func (suo *StockUpdateOne) SetInService(i int32) *StockUpdateOne {
	suo.mutation.ResetInService()
	suo.mutation.SetInService(i)
	return suo
}

// AddInService adds i to the "in_service" field.
func (suo *StockUpdateOne) AddInService(i int32) *StockUpdateOne {
	suo.mutation.AddInService(i)
	return suo
}

// SetSold sets the "sold" field.
func (suo *StockUpdateOne) SetSold(i int32) *StockUpdateOne {
	suo.mutation.ResetSold()
	suo.mutation.SetSold(i)
	return suo
}

// AddSold adds i to the "sold" field.
func (suo *StockUpdateOne) AddSold(i int32) *StockUpdateOne {
	suo.mutation.AddSold(i)
	return suo
}

// Mutation returns the StockMutation object of the builder.
func (suo *StockUpdateOne) Mutation() *StockMutation {
	return suo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StockUpdateOne) Select(field string, fields ...string) *StockUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Stock entity.
func (suo *StockUpdateOne) Save(ctx context.Context) (*Stock, error) {
	var (
		err  error
		node *Stock
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StockUpdateOne) SaveX(ctx context.Context) *Stock {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StockUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StockUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *StockUpdateOne) sqlSave(ctx context.Context) (_node *Stock, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   stock.Table,
			Columns: stock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: stock.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Stock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, stock.FieldID)
		for _, f := range fields {
			if !stock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != stock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: stock.FieldGoodID,
		})
	}
	if value, ok := suo.mutation.Total(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldTotal,
		})
	}
	if value, ok := suo.mutation.AddedTotal(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldTotal,
		})
	}
	if value, ok := suo.mutation.InService(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldInService,
		})
	}
	if value, ok := suo.mutation.AddedInService(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldInService,
		})
	}
	if value, ok := suo.mutation.Sold(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldSold,
		})
	}
	if value, ok := suo.mutation.AddedSold(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldSold,
		})
	}
	_node = &Stock{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{stock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
