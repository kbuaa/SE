// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/itclub/app/ent/migrate"

	"github.com/itclub/app/ent/doctor"
	"github.com/itclub/app/ent/drug"
	"github.com/itclub/app/ent/nurse"
	"github.com/itclub/app/ent/patient"
	"github.com/itclub/app/ent/prescription"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Doctor is the client for interacting with the Doctor builders.
	Doctor *DoctorClient
	// Drug is the client for interacting with the Drug builders.
	Drug *DrugClient
	// Nurse is the client for interacting with the Nurse builders.
	Nurse *NurseClient
	// Patient is the client for interacting with the Patient builders.
	Patient *PatientClient
	// Prescription is the client for interacting with the Prescription builders.
	Prescription *PrescriptionClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Doctor = NewDoctorClient(c.config)
	c.Drug = NewDrugClient(c.config)
	c.Nurse = NewNurseClient(c.config)
	c.Patient = NewPatientClient(c.config)
	c.Prescription = NewPrescriptionClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Doctor:       NewDoctorClient(cfg),
		Drug:         NewDrugClient(cfg),
		Nurse:        NewNurseClient(cfg),
		Patient:      NewPatientClient(cfg),
		Prescription: NewPrescriptionClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:       cfg,
		Doctor:       NewDoctorClient(cfg),
		Drug:         NewDrugClient(cfg),
		Nurse:        NewNurseClient(cfg),
		Patient:      NewPatientClient(cfg),
		Prescription: NewPrescriptionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Doctor.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Doctor.Use(hooks...)
	c.Drug.Use(hooks...)
	c.Nurse.Use(hooks...)
	c.Patient.Use(hooks...)
	c.Prescription.Use(hooks...)
}

// DoctorClient is a client for the Doctor schema.
type DoctorClient struct {
	config
}

// NewDoctorClient returns a client for the Doctor from the given config.
func NewDoctorClient(c config) *DoctorClient {
	return &DoctorClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `doctor.Hooks(f(g(h())))`.
func (c *DoctorClient) Use(hooks ...Hook) {
	c.hooks.Doctor = append(c.hooks.Doctor, hooks...)
}

// Create returns a create builder for Doctor.
func (c *DoctorClient) Create() *DoctorCreate {
	mutation := newDoctorMutation(c.config, OpCreate)
	return &DoctorCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Doctor.
func (c *DoctorClient) Update() *DoctorUpdate {
	mutation := newDoctorMutation(c.config, OpUpdate)
	return &DoctorUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DoctorClient) UpdateOne(d *Doctor) *DoctorUpdateOne {
	mutation := newDoctorMutation(c.config, OpUpdateOne, withDoctor(d))
	return &DoctorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DoctorClient) UpdateOneID(id int) *DoctorUpdateOne {
	mutation := newDoctorMutation(c.config, OpUpdateOne, withDoctorID(id))
	return &DoctorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Doctor.
func (c *DoctorClient) Delete() *DoctorDelete {
	mutation := newDoctorMutation(c.config, OpDelete)
	return &DoctorDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DoctorClient) DeleteOne(d *Doctor) *DoctorDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DoctorClient) DeleteOneID(id int) *DoctorDeleteOne {
	builder := c.Delete().Where(doctor.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DoctorDeleteOne{builder}
}

// Create returns a query builder for Doctor.
func (c *DoctorClient) Query() *DoctorQuery {
	return &DoctorQuery{config: c.config}
}

// Get returns a Doctor entity by its id.
func (c *DoctorClient) Get(ctx context.Context, id int) (*Doctor, error) {
	return c.Query().Where(doctor.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DoctorClient) GetX(ctx context.Context, id int) *Doctor {
	d, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return d
}

// QueryDoctorPrescription queries the doctor_prescription edge of a Doctor.
func (c *DoctorClient) QueryDoctorPrescription(d *Doctor) *PrescriptionQuery {
	query := &PrescriptionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(doctor.Table, doctor.FieldID, id),
			sqlgraph.To(prescription.Table, prescription.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, doctor.DoctorPrescriptionTable, doctor.DoctorPrescriptionColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DoctorClient) Hooks() []Hook {
	return c.hooks.Doctor
}

// DrugClient is a client for the Drug schema.
type DrugClient struct {
	config
}

// NewDrugClient returns a client for the Drug from the given config.
func NewDrugClient(c config) *DrugClient {
	return &DrugClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `drug.Hooks(f(g(h())))`.
func (c *DrugClient) Use(hooks ...Hook) {
	c.hooks.Drug = append(c.hooks.Drug, hooks...)
}

// Create returns a create builder for Drug.
func (c *DrugClient) Create() *DrugCreate {
	mutation := newDrugMutation(c.config, OpCreate)
	return &DrugCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Drug.
func (c *DrugClient) Update() *DrugUpdate {
	mutation := newDrugMutation(c.config, OpUpdate)
	return &DrugUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DrugClient) UpdateOne(d *Drug) *DrugUpdateOne {
	mutation := newDrugMutation(c.config, OpUpdateOne, withDrug(d))
	return &DrugUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DrugClient) UpdateOneID(id int) *DrugUpdateOne {
	mutation := newDrugMutation(c.config, OpUpdateOne, withDrugID(id))
	return &DrugUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Drug.
func (c *DrugClient) Delete() *DrugDelete {
	mutation := newDrugMutation(c.config, OpDelete)
	return &DrugDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DrugClient) DeleteOne(d *Drug) *DrugDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DrugClient) DeleteOneID(id int) *DrugDeleteOne {
	builder := c.Delete().Where(drug.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DrugDeleteOne{builder}
}

// Create returns a query builder for Drug.
func (c *DrugClient) Query() *DrugQuery {
	return &DrugQuery{config: c.config}
}

// Get returns a Drug entity by its id.
func (c *DrugClient) Get(ctx context.Context, id int) (*Drug, error) {
	return c.Query().Where(drug.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DrugClient) GetX(ctx context.Context, id int) *Drug {
	d, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return d
}

// QueryDrugPrescription queries the drug_prescription edge of a Drug.
func (c *DrugClient) QueryDrugPrescription(d *Drug) *PrescriptionQuery {
	query := &PrescriptionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(drug.Table, drug.FieldID, id),
			sqlgraph.To(prescription.Table, prescription.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, drug.DrugPrescriptionTable, drug.DrugPrescriptionColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DrugClient) Hooks() []Hook {
	return c.hooks.Drug
}

// NurseClient is a client for the Nurse schema.
type NurseClient struct {
	config
}

// NewNurseClient returns a client for the Nurse from the given config.
func NewNurseClient(c config) *NurseClient {
	return &NurseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `nurse.Hooks(f(g(h())))`.
func (c *NurseClient) Use(hooks ...Hook) {
	c.hooks.Nurse = append(c.hooks.Nurse, hooks...)
}

// Create returns a create builder for Nurse.
func (c *NurseClient) Create() *NurseCreate {
	mutation := newNurseMutation(c.config, OpCreate)
	return &NurseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Nurse.
func (c *NurseClient) Update() *NurseUpdate {
	mutation := newNurseMutation(c.config, OpUpdate)
	return &NurseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NurseClient) UpdateOne(n *Nurse) *NurseUpdateOne {
	mutation := newNurseMutation(c.config, OpUpdateOne, withNurse(n))
	return &NurseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NurseClient) UpdateOneID(id int) *NurseUpdateOne {
	mutation := newNurseMutation(c.config, OpUpdateOne, withNurseID(id))
	return &NurseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Nurse.
func (c *NurseClient) Delete() *NurseDelete {
	mutation := newNurseMutation(c.config, OpDelete)
	return &NurseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *NurseClient) DeleteOne(n *Nurse) *NurseDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *NurseClient) DeleteOneID(id int) *NurseDeleteOne {
	builder := c.Delete().Where(nurse.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NurseDeleteOne{builder}
}

// Create returns a query builder for Nurse.
func (c *NurseClient) Query() *NurseQuery {
	return &NurseQuery{config: c.config}
}

// Get returns a Nurse entity by its id.
func (c *NurseClient) Get(ctx context.Context, id int) (*Nurse, error) {
	return c.Query().Where(nurse.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NurseClient) GetX(ctx context.Context, id int) *Nurse {
	n, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return n
}

// QueryNursePrescription queries the nurse_prescription edge of a Nurse.
func (c *NurseClient) QueryNursePrescription(n *Nurse) *PrescriptionQuery {
	query := &PrescriptionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(nurse.Table, nurse.FieldID, id),
			sqlgraph.To(prescription.Table, prescription.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, nurse.NursePrescriptionTable, nurse.NursePrescriptionColumn),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *NurseClient) Hooks() []Hook {
	return c.hooks.Nurse
}

// PatientClient is a client for the Patient schema.
type PatientClient struct {
	config
}

// NewPatientClient returns a client for the Patient from the given config.
func NewPatientClient(c config) *PatientClient {
	return &PatientClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `patient.Hooks(f(g(h())))`.
func (c *PatientClient) Use(hooks ...Hook) {
	c.hooks.Patient = append(c.hooks.Patient, hooks...)
}

// Create returns a create builder for Patient.
func (c *PatientClient) Create() *PatientCreate {
	mutation := newPatientMutation(c.config, OpCreate)
	return &PatientCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Patient.
func (c *PatientClient) Update() *PatientUpdate {
	mutation := newPatientMutation(c.config, OpUpdate)
	return &PatientUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PatientClient) UpdateOne(pa *Patient) *PatientUpdateOne {
	mutation := newPatientMutation(c.config, OpUpdateOne, withPatient(pa))
	return &PatientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PatientClient) UpdateOneID(id int) *PatientUpdateOne {
	mutation := newPatientMutation(c.config, OpUpdateOne, withPatientID(id))
	return &PatientUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Patient.
func (c *PatientClient) Delete() *PatientDelete {
	mutation := newPatientMutation(c.config, OpDelete)
	return &PatientDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PatientClient) DeleteOne(pa *Patient) *PatientDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PatientClient) DeleteOneID(id int) *PatientDeleteOne {
	builder := c.Delete().Where(patient.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PatientDeleteOne{builder}
}

// Create returns a query builder for Patient.
func (c *PatientClient) Query() *PatientQuery {
	return &PatientQuery{config: c.config}
}

// Get returns a Patient entity by its id.
func (c *PatientClient) Get(ctx context.Context, id int) (*Patient, error) {
	return c.Query().Where(patient.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PatientClient) GetX(ctx context.Context, id int) *Patient {
	pa, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pa
}

// QueryPatientPrescription queries the patient_prescription edge of a Patient.
func (c *PatientClient) QueryPatientPrescription(pa *Patient) *PrescriptionQuery {
	query := &PrescriptionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(patient.Table, patient.FieldID, id),
			sqlgraph.To(prescription.Table, prescription.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, patient.PatientPrescriptionTable, patient.PatientPrescriptionColumn),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PatientClient) Hooks() []Hook {
	return c.hooks.Patient
}

// PrescriptionClient is a client for the Prescription schema.
type PrescriptionClient struct {
	config
}

// NewPrescriptionClient returns a client for the Prescription from the given config.
func NewPrescriptionClient(c config) *PrescriptionClient {
	return &PrescriptionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `prescription.Hooks(f(g(h())))`.
func (c *PrescriptionClient) Use(hooks ...Hook) {
	c.hooks.Prescription = append(c.hooks.Prescription, hooks...)
}

// Create returns a create builder for Prescription.
func (c *PrescriptionClient) Create() *PrescriptionCreate {
	mutation := newPrescriptionMutation(c.config, OpCreate)
	return &PrescriptionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Prescription.
func (c *PrescriptionClient) Update() *PrescriptionUpdate {
	mutation := newPrescriptionMutation(c.config, OpUpdate)
	return &PrescriptionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PrescriptionClient) UpdateOne(pr *Prescription) *PrescriptionUpdateOne {
	mutation := newPrescriptionMutation(c.config, OpUpdateOne, withPrescription(pr))
	return &PrescriptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PrescriptionClient) UpdateOneID(id int) *PrescriptionUpdateOne {
	mutation := newPrescriptionMutation(c.config, OpUpdateOne, withPrescriptionID(id))
	return &PrescriptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Prescription.
func (c *PrescriptionClient) Delete() *PrescriptionDelete {
	mutation := newPrescriptionMutation(c.config, OpDelete)
	return &PrescriptionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PrescriptionClient) DeleteOne(pr *Prescription) *PrescriptionDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PrescriptionClient) DeleteOneID(id int) *PrescriptionDeleteOne {
	builder := c.Delete().Where(prescription.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PrescriptionDeleteOne{builder}
}

// Create returns a query builder for Prescription.
func (c *PrescriptionClient) Query() *PrescriptionQuery {
	return &PrescriptionQuery{config: c.config}
}

// Get returns a Prescription entity by its id.
func (c *PrescriptionClient) Get(ctx context.Context, id int) (*Prescription, error) {
	return c.Query().Where(prescription.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PrescriptionClient) GetX(ctx context.Context, id int) *Prescription {
	pr, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pr
}

// QueryDoctor queries the doctor edge of a Prescription.
func (c *PrescriptionClient) QueryDoctor(pr *Prescription) *DoctorQuery {
	query := &DoctorQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(prescription.Table, prescription.FieldID, id),
			sqlgraph.To(doctor.Table, doctor.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, prescription.DoctorTable, prescription.DoctorColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPatient queries the patient edge of a Prescription.
func (c *PrescriptionClient) QueryPatient(pr *Prescription) *PatientQuery {
	query := &PatientQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(prescription.Table, prescription.FieldID, id),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, prescription.PatientTable, prescription.PatientColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryNurse queries the nurse edge of a Prescription.
func (c *PrescriptionClient) QueryNurse(pr *Prescription) *NurseQuery {
	query := &NurseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(prescription.Table, prescription.FieldID, id),
			sqlgraph.To(nurse.Table, nurse.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, prescription.NurseTable, prescription.NurseColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDrug queries the drug edge of a Prescription.
func (c *PrescriptionClient) QueryDrug(pr *Prescription) *DrugQuery {
	query := &DrugQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(prescription.Table, prescription.FieldID, id),
			sqlgraph.To(drug.Table, drug.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, prescription.DrugTable, prescription.DrugColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PrescriptionClient) Hooks() []Hook {
	return c.hooks.Prescription
}