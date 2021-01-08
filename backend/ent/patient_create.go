// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/itclub/app/ent/patient"
	"github.com/itclub/app/ent/prescription"
)

// PatientCreate is the builder for creating a Patient entity.
type PatientCreate struct {
	config
	mutation *PatientMutation
	hooks    []Hook
}

// SetPatientName sets the Patient_Name field.
func (pc *PatientCreate) SetPatientName(s string) *PatientCreate {
	pc.mutation.SetPatientName(s)
	return pc
}

// AddPatientPrescriptionIDs adds the patient_prescription edge to Prescription by ids.
func (pc *PatientCreate) AddPatientPrescriptionIDs(ids ...int) *PatientCreate {
	pc.mutation.AddPatientPrescriptionIDs(ids...)
	return pc
}

// AddPatientPrescription adds the patient_prescription edges to Prescription.
func (pc *PatientCreate) AddPatientPrescription(p ...*Prescription) *PatientCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPatientPrescriptionIDs(ids...)
}

// Mutation returns the PatientMutation object of the builder.
func (pc *PatientCreate) Mutation() *PatientMutation {
	return pc.mutation
}

// Save creates the Patient in the database.
func (pc *PatientCreate) Save(ctx context.Context) (*Patient, error) {
	if _, ok := pc.mutation.PatientName(); !ok {
		return nil, &ValidationError{Name: "Patient_Name", err: errors.New("ent: missing required field \"Patient_Name\"")}
	}
	if v, ok := pc.mutation.PatientName(); ok {
		if err := patient.PatientNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Patient_Name", err: fmt.Errorf("ent: validator failed for field \"Patient_Name\": %w", err)}
		}
	}
	var (
		err  error
		node *Patient
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PatientCreate) SaveX(ctx context.Context) *Patient {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PatientCreate) sqlSave(ctx context.Context) (*Patient, error) {
	pa, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pa.ID = int(id)
	return pa, nil
}

func (pc *PatientCreate) createSpec() (*Patient, *sqlgraph.CreateSpec) {
	var (
		pa    = &Patient{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: patient.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.PatientName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientName,
		})
		pa.PatientName = value
	}
	if nodes := pc.mutation.PatientPrescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PatientPrescriptionTable,
			Columns: []string{patient.PatientPrescriptionColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pa, _spec
}
