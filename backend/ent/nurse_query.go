// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/itclub/app/ent/nurse"
	"github.com/itclub/app/ent/predicate"
	"github.com/itclub/app/ent/prescription"
)

// NurseQuery is the builder for querying Nurse entities.
type NurseQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Nurse
	// eager-loading edges.
	withNursePrescription *PrescriptionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (nq *NurseQuery) Where(ps ...predicate.Nurse) *NurseQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit adds a limit step to the query.
func (nq *NurseQuery) Limit(limit int) *NurseQuery {
	nq.limit = &limit
	return nq
}

// Offset adds an offset step to the query.
func (nq *NurseQuery) Offset(offset int) *NurseQuery {
	nq.offset = &offset
	return nq
}

// Order adds an order step to the query.
func (nq *NurseQuery) Order(o ...OrderFunc) *NurseQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryNursePrescription chains the current query on the nurse_prescription edge.
func (nq *NurseQuery) QueryNursePrescription() *PrescriptionQuery {
	query := &PrescriptionQuery{config: nq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(nurse.Table, nurse.FieldID, nq.sqlQuery()),
			sqlgraph.To(prescription.Table, prescription.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, nurse.NursePrescriptionTable, nurse.NursePrescriptionColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Nurse entity in the query. Returns *NotFoundError when no nurse was found.
func (nq *NurseQuery) First(ctx context.Context) (*Nurse, error) {
	ns, err := nq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ns) == 0 {
		return nil, &NotFoundError{nurse.Label}
	}
	return ns[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NurseQuery) FirstX(ctx context.Context) *Nurse {
	n, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return n
}

// FirstID returns the first Nurse id in the query. Returns *NotFoundError when no id was found.
func (nq *NurseQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{nurse.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (nq *NurseQuery) FirstXID(ctx context.Context) int {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Nurse entity in the query, returns an error if not exactly one entity was returned.
func (nq *NurseQuery) Only(ctx context.Context) (*Nurse, error) {
	ns, err := nq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ns) {
	case 1:
		return ns[0], nil
	case 0:
		return nil, &NotFoundError{nurse.Label}
	default:
		return nil, &NotSingularError{nurse.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NurseQuery) OnlyX(ctx context.Context) *Nurse {
	n, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

// OnlyID returns the only Nurse id in the query, returns an error if not exactly one id was returned.
func (nq *NurseQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{nurse.Label}
	default:
		err = &NotSingularError{nurse.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NurseQuery) OnlyIDX(ctx context.Context) int {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Nurses.
func (nq *NurseQuery) All(ctx context.Context) ([]*Nurse, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return nq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nq *NurseQuery) AllX(ctx context.Context) []*Nurse {
	ns, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ns
}

// IDs executes the query and returns a list of Nurse ids.
func (nq *NurseQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := nq.Select(nurse.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NurseQuery) IDsX(ctx context.Context) []int {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NurseQuery) Count(ctx context.Context) (int, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return nq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NurseQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NurseQuery) Exist(ctx context.Context) (bool, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return nq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NurseQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NurseQuery) Clone() *NurseQuery {
	return &NurseQuery{
		config:     nq.config,
		limit:      nq.limit,
		offset:     nq.offset,
		order:      append([]OrderFunc{}, nq.order...),
		unique:     append([]string{}, nq.unique...),
		predicates: append([]predicate.Nurse{}, nq.predicates...),
		// clone intermediate query.
		sql:  nq.sql.Clone(),
		path: nq.path,
	}
}

//  WithNursePrescription tells the query-builder to eager-loads the nodes that are connected to
// the "nurse_prescription" edge. The optional arguments used to configure the query builder of the edge.
func (nq *NurseQuery) WithNursePrescription(opts ...func(*PrescriptionQuery)) *NurseQuery {
	query := &PrescriptionQuery{config: nq.config}
	for _, opt := range opts {
		opt(query)
	}
	nq.withNursePrescription = query
	return nq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		NurseEmail string `json:"Nurse_Email,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Nurse.Query().
//		GroupBy(nurse.FieldNurseEmail).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (nq *NurseQuery) GroupBy(field string, fields ...string) *NurseGroupBy {
	group := &NurseGroupBy{config: nq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return nq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		NurseEmail string `json:"Nurse_Email,omitempty"`
//	}
//
//	client.Nurse.Query().
//		Select(nurse.FieldNurseEmail).
//		Scan(ctx, &v)
//
func (nq *NurseQuery) Select(field string, fields ...string) *NurseSelect {
	selector := &NurseSelect{config: nq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return nq.sqlQuery(), nil
	}
	return selector
}

func (nq *NurseQuery) prepareQuery(ctx context.Context) error {
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NurseQuery) sqlAll(ctx context.Context) ([]*Nurse, error) {
	var (
		nodes       = []*Nurse{}
		_spec       = nq.querySpec()
		loadedTypes = [1]bool{
			nq.withNursePrescription != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Nurse{config: nq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := nq.withNursePrescription; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Nurse)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Prescription(func(s *sql.Selector) {
			s.Where(sql.InValues(nurse.NursePrescriptionColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.nurse_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "nurse_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "nurse_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.NursePrescription = append(node.Edges.NursePrescription, n)
		}
	}

	return nodes, nil
}

func (nq *NurseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NurseQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := nq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (nq *NurseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nurse.Table,
			Columns: nurse.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nurse.FieldID,
			},
		},
		From:   nq.sql,
		Unique: true,
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NurseQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(nurse.Table)
	selector := builder.Select(t1.Columns(nurse.Columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(nurse.Columns...)...)
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NurseGroupBy is the builder for group-by Nurse entities.
type NurseGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NurseGroupBy) Aggregate(fns ...AggregateFunc) *NurseGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the group-by query and scan the result into the given value.
func (ngb *NurseGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ngb.path(ctx)
	if err != nil {
		return err
	}
	ngb.sql = query
	return ngb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ngb *NurseGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ngb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ngb *NurseGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ngb.fields) > 1 {
		return nil, errors.New("ent: NurseGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ngb *NurseGroupBy) StringsX(ctx context.Context) []string {
	v, err := ngb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (ngb *NurseGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ngb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{nurse.Label}
	default:
		err = fmt.Errorf("ent: NurseGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ngb *NurseGroupBy) StringX(ctx context.Context) string {
	v, err := ngb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ngb *NurseGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ngb.fields) > 1 {
		return nil, errors.New("ent: NurseGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ngb *NurseGroupBy) IntsX(ctx context.Context) []int {
	v, err := ngb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (ngb *NurseGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ngb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{nurse.Label}
	default:
		err = fmt.Errorf("ent: NurseGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ngb *NurseGroupBy) IntX(ctx context.Context) int {
	v, err := ngb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ngb *NurseGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ngb.fields) > 1 {
		return nil, errors.New("ent: NurseGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ngb *NurseGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ngb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (ngb *NurseGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ngb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{nurse.Label}
	default:
		err = fmt.Errorf("ent: NurseGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ngb *NurseGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ngb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ngb *NurseGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ngb.fields) > 1 {
		return nil, errors.New("ent: NurseGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ngb *NurseGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ngb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (ngb *NurseGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ngb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{nurse.Label}
	default:
		err = fmt.Errorf("ent: NurseGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ngb *NurseGroupBy) BoolX(ctx context.Context) bool {
	v, err := ngb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ngb *NurseGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ngb.sqlQuery().Query()
	if err := ngb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ngb *NurseGroupBy) sqlQuery() *sql.Selector {
	selector := ngb.sql
	columns := make([]string, 0, len(ngb.fields)+len(ngb.fns))
	columns = append(columns, ngb.fields...)
	for _, fn := range ngb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(ngb.fields...)
}

// NurseSelect is the builder for select fields of Nurse entities.
type NurseSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ns *NurseSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ns.path(ctx)
	if err != nil {
		return err
	}
	ns.sql = query
	return ns.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ns *NurseSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ns.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ns *NurseSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ns.fields) > 1 {
		return nil, errors.New("ent: NurseSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ns *NurseSelect) StringsX(ctx context.Context) []string {
	v, err := ns.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ns *NurseSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ns.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{nurse.Label}
	default:
		err = fmt.Errorf("ent: NurseSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ns *NurseSelect) StringX(ctx context.Context) string {
	v, err := ns.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ns *NurseSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ns.fields) > 1 {
		return nil, errors.New("ent: NurseSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ns *NurseSelect) IntsX(ctx context.Context) []int {
	v, err := ns.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ns *NurseSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ns.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{nurse.Label}
	default:
		err = fmt.Errorf("ent: NurseSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ns *NurseSelect) IntX(ctx context.Context) int {
	v, err := ns.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ns *NurseSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ns.fields) > 1 {
		return nil, errors.New("ent: NurseSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ns *NurseSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ns.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ns *NurseSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ns.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{nurse.Label}
	default:
		err = fmt.Errorf("ent: NurseSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ns *NurseSelect) Float64X(ctx context.Context) float64 {
	v, err := ns.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ns *NurseSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ns.fields) > 1 {
		return nil, errors.New("ent: NurseSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ns *NurseSelect) BoolsX(ctx context.Context) []bool {
	v, err := ns.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ns *NurseSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ns.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{nurse.Label}
	default:
		err = fmt.Errorf("ent: NurseSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ns *NurseSelect) BoolX(ctx context.Context) bool {
	v, err := ns.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ns *NurseSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ns.sqlQuery().Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ns *NurseSelect) sqlQuery() sql.Querier {
	selector := ns.sql
	selector.Select(selector.Columns(ns.fields...)...)
	return selector
}
