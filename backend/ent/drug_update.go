// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/itclub/app/ent/drug"
	"github.com/itclub/app/ent/predicate"
	"github.com/itclub/app/ent/prescription"
)

// DrugUpdate is the builder for updating Drug entities.
type DrugUpdate struct {
	config
	hooks      []Hook
	mutation   *DrugMutation
	predicates []predicate.Drug
}

// Where adds a new predicate for the builder.
func (du *DrugUpdate) Where(ps ...predicate.Drug) *DrugUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetDrugName sets the Drug_Name field.
func (du *DrugUpdate) SetDrugName(s string) *DrugUpdate {
	du.mutation.SetDrugName(s)
	return du
}

// AddDrugPrescriptionIDs adds the drug_prescription edge to Prescription by ids.
func (du *DrugUpdate) AddDrugPrescriptionIDs(ids ...int) *DrugUpdate {
	du.mutation.AddDrugPrescriptionIDs(ids...)
	return du
}

// AddDrugPrescription adds the drug_prescription edges to Prescription.
func (du *DrugUpdate) AddDrugPrescription(p ...*Prescription) *DrugUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.AddDrugPrescriptionIDs(ids...)
}

// Mutation returns the DrugMutation object of the builder.
func (du *DrugUpdate) Mutation() *DrugMutation {
	return du.mutation
}

// RemoveDrugPrescriptionIDs removes the drug_prescription edge to Prescription by ids.
func (du *DrugUpdate) RemoveDrugPrescriptionIDs(ids ...int) *DrugUpdate {
	du.mutation.RemoveDrugPrescriptionIDs(ids...)
	return du
}

// RemoveDrugPrescription removes drug_prescription edges to Prescription.
func (du *DrugUpdate) RemoveDrugPrescription(p ...*Prescription) *DrugUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.RemoveDrugPrescriptionIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DrugUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.DrugName(); ok {
		if err := drug.DrugNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "Drug_Name", err: fmt.Errorf("ent: validator failed for field \"Drug_Name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DrugMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DrugUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DrugUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DrugUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DrugUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   drug.Table,
			Columns: drug.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: drug.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.DrugName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: drug.FieldDrugName,
		})
	}
	if nodes := du.mutation.RemovedDrugPrescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   drug.DrugPrescriptionTable,
			Columns: []string{drug.DrugPrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prescription.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DrugPrescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   drug.DrugPrescriptionTable,
			Columns: []string{drug.DrugPrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prescription.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{drug.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DrugUpdateOne is the builder for updating a single Drug entity.
type DrugUpdateOne struct {
	config
	hooks    []Hook
	mutation *DrugMutation
}

// SetDrugName sets the Drug_Name field.
func (duo *DrugUpdateOne) SetDrugName(s string) *DrugUpdateOne {
	duo.mutation.SetDrugName(s)
	return duo
}

// AddDrugPrescriptionIDs adds the drug_prescription edge to Prescription by ids.
func (duo *DrugUpdateOne) AddDrugPrescriptionIDs(ids ...int) *DrugUpdateOne {
	duo.mutation.AddDrugPrescriptionIDs(ids...)
	return duo
}

// AddDrugPrescription adds the drug_prescription edges to Prescription.
func (duo *DrugUpdateOne) AddDrugPrescription(p ...*Prescription) *DrugUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.AddDrugPrescriptionIDs(ids...)
}

// Mutation returns the DrugMutation object of the builder.
func (duo *DrugUpdateOne) Mutation() *DrugMutation {
	return duo.mutation
}

// RemoveDrugPrescriptionIDs removes the drug_prescription edge to Prescription by ids.
func (duo *DrugUpdateOne) RemoveDrugPrescriptionIDs(ids ...int) *DrugUpdateOne {
	duo.mutation.RemoveDrugPrescriptionIDs(ids...)
	return duo
}

// RemoveDrugPrescription removes drug_prescription edges to Prescription.
func (duo *DrugUpdateOne) RemoveDrugPrescription(p ...*Prescription) *DrugUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.RemoveDrugPrescriptionIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DrugUpdateOne) Save(ctx context.Context) (*Drug, error) {
	if v, ok := duo.mutation.DrugName(); ok {
		if err := drug.DrugNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Drug_Name", err: fmt.Errorf("ent: validator failed for field \"Drug_Name\": %w", err)}
		}
	}

	var (
		err  error
		node *Drug
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DrugMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DrugUpdateOne) SaveX(ctx context.Context) *Drug {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DrugUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DrugUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DrugUpdateOne) sqlSave(ctx context.Context) (d *Drug, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   drug.Table,
			Columns: drug.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: drug.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Drug.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.DrugName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: drug.FieldDrugName,
		})
	}
	if nodes := duo.mutation.RemovedDrugPrescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   drug.DrugPrescriptionTable,
			Columns: []string{drug.DrugPrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prescription.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DrugPrescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   drug.DrugPrescriptionTable,
			Columns: []string{drug.DrugPrescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prescription.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	d = &Drug{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{drug.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}
