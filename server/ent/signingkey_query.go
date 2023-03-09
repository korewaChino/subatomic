// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/FyraLabs/subatomic/server/ent/predicate"
	"github.com/FyraLabs/subatomic/server/ent/repo"
	"github.com/FyraLabs/subatomic/server/ent/signingkey"
)

// SigningKeyQuery is the builder for querying SigningKey entities.
type SigningKeyQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.SigningKey
	withRepo   *RepoQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SigningKeyQuery builder.
func (skq *SigningKeyQuery) Where(ps ...predicate.SigningKey) *SigningKeyQuery {
	skq.predicates = append(skq.predicates, ps...)
	return skq
}

// Limit the number of records to be returned by this query.
func (skq *SigningKeyQuery) Limit(limit int) *SigningKeyQuery {
	skq.ctx.Limit = &limit
	return skq
}

// Offset to start from.
func (skq *SigningKeyQuery) Offset(offset int) *SigningKeyQuery {
	skq.ctx.Offset = &offset
	return skq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (skq *SigningKeyQuery) Unique(unique bool) *SigningKeyQuery {
	skq.ctx.Unique = &unique
	return skq
}

// Order specifies how the records should be ordered.
func (skq *SigningKeyQuery) Order(o ...OrderFunc) *SigningKeyQuery {
	skq.order = append(skq.order, o...)
	return skq
}

// QueryRepo chains the current query on the "repo" edge.
func (skq *SigningKeyQuery) QueryRepo() *RepoQuery {
	query := (&RepoClient{config: skq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := skq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := skq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(signingkey.Table, signingkey.FieldID, selector),
			sqlgraph.To(repo.Table, repo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, signingkey.RepoTable, signingkey.RepoColumn),
		)
		fromU = sqlgraph.SetNeighbors(skq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SigningKey entity from the query.
// Returns a *NotFoundError when no SigningKey was found.
func (skq *SigningKeyQuery) First(ctx context.Context) (*SigningKey, error) {
	nodes, err := skq.Limit(1).All(setContextOp(ctx, skq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{signingkey.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (skq *SigningKeyQuery) FirstX(ctx context.Context) *SigningKey {
	node, err := skq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SigningKey ID from the query.
// Returns a *NotFoundError when no SigningKey ID was found.
func (skq *SigningKeyQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = skq.Limit(1).IDs(setContextOp(ctx, skq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{signingkey.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (skq *SigningKeyQuery) FirstIDX(ctx context.Context) string {
	id, err := skq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SigningKey entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SigningKey entity is found.
// Returns a *NotFoundError when no SigningKey entities are found.
func (skq *SigningKeyQuery) Only(ctx context.Context) (*SigningKey, error) {
	nodes, err := skq.Limit(2).All(setContextOp(ctx, skq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{signingkey.Label}
	default:
		return nil, &NotSingularError{signingkey.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (skq *SigningKeyQuery) OnlyX(ctx context.Context) *SigningKey {
	node, err := skq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SigningKey ID in the query.
// Returns a *NotSingularError when more than one SigningKey ID is found.
// Returns a *NotFoundError when no entities are found.
func (skq *SigningKeyQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = skq.Limit(2).IDs(setContextOp(ctx, skq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{signingkey.Label}
	default:
		err = &NotSingularError{signingkey.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (skq *SigningKeyQuery) OnlyIDX(ctx context.Context) string {
	id, err := skq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SigningKeys.
func (skq *SigningKeyQuery) All(ctx context.Context) ([]*SigningKey, error) {
	ctx = setContextOp(ctx, skq.ctx, "All")
	if err := skq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SigningKey, *SigningKeyQuery]()
	return withInterceptors[[]*SigningKey](ctx, skq, qr, skq.inters)
}

// AllX is like All, but panics if an error occurs.
func (skq *SigningKeyQuery) AllX(ctx context.Context) []*SigningKey {
	nodes, err := skq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SigningKey IDs.
func (skq *SigningKeyQuery) IDs(ctx context.Context) (ids []string, err error) {
	if skq.ctx.Unique == nil && skq.path != nil {
		skq.Unique(true)
	}
	ctx = setContextOp(ctx, skq.ctx, "IDs")
	if err = skq.Select(signingkey.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (skq *SigningKeyQuery) IDsX(ctx context.Context) []string {
	ids, err := skq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (skq *SigningKeyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, skq.ctx, "Count")
	if err := skq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, skq, querierCount[*SigningKeyQuery](), skq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (skq *SigningKeyQuery) CountX(ctx context.Context) int {
	count, err := skq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (skq *SigningKeyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, skq.ctx, "Exist")
	switch _, err := skq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (skq *SigningKeyQuery) ExistX(ctx context.Context) bool {
	exist, err := skq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SigningKeyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (skq *SigningKeyQuery) Clone() *SigningKeyQuery {
	if skq == nil {
		return nil
	}
	return &SigningKeyQuery{
		config:     skq.config,
		ctx:        skq.ctx.Clone(),
		order:      append([]OrderFunc{}, skq.order...),
		inters:     append([]Interceptor{}, skq.inters...),
		predicates: append([]predicate.SigningKey{}, skq.predicates...),
		withRepo:   skq.withRepo.Clone(),
		// clone intermediate query.
		sql:  skq.sql.Clone(),
		path: skq.path,
	}
}

// WithRepo tells the query-builder to eager-load the nodes that are connected to
// the "repo" edge. The optional arguments are used to configure the query builder of the edge.
func (skq *SigningKeyQuery) WithRepo(opts ...func(*RepoQuery)) *SigningKeyQuery {
	query := (&RepoClient{config: skq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	skq.withRepo = query
	return skq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PrivateKey string `json:"private_key,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SigningKey.Query().
//		GroupBy(signingkey.FieldPrivateKey).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (skq *SigningKeyQuery) GroupBy(field string, fields ...string) *SigningKeyGroupBy {
	skq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SigningKeyGroupBy{build: skq}
	grbuild.flds = &skq.ctx.Fields
	grbuild.label = signingkey.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PrivateKey string `json:"private_key,omitempty"`
//	}
//
//	client.SigningKey.Query().
//		Select(signingkey.FieldPrivateKey).
//		Scan(ctx, &v)
func (skq *SigningKeyQuery) Select(fields ...string) *SigningKeySelect {
	skq.ctx.Fields = append(skq.ctx.Fields, fields...)
	sbuild := &SigningKeySelect{SigningKeyQuery: skq}
	sbuild.label = signingkey.Label
	sbuild.flds, sbuild.scan = &skq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SigningKeySelect configured with the given aggregations.
func (skq *SigningKeyQuery) Aggregate(fns ...AggregateFunc) *SigningKeySelect {
	return skq.Select().Aggregate(fns...)
}

func (skq *SigningKeyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range skq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, skq); err != nil {
				return err
			}
		}
	}
	for _, f := range skq.ctx.Fields {
		if !signingkey.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if skq.path != nil {
		prev, err := skq.path(ctx)
		if err != nil {
			return err
		}
		skq.sql = prev
	}
	return nil
}

func (skq *SigningKeyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SigningKey, error) {
	var (
		nodes       = []*SigningKey{}
		_spec       = skq.querySpec()
		loadedTypes = [1]bool{
			skq.withRepo != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SigningKey).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SigningKey{config: skq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, skq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := skq.withRepo; query != nil {
		if err := skq.loadRepo(ctx, query, nodes,
			func(n *SigningKey) { n.Edges.Repo = []*Repo{} },
			func(n *SigningKey, e *Repo) { n.Edges.Repo = append(n.Edges.Repo, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (skq *SigningKeyQuery) loadRepo(ctx context.Context, query *RepoQuery, nodes []*SigningKey, init func(*SigningKey), assign func(*SigningKey, *Repo)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*SigningKey)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Repo(func(s *sql.Selector) {
		s.Where(sql.InValues(signingkey.RepoColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.repo_key
		if fk == nil {
			return fmt.Errorf(`foreign-key "repo_key" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "repo_key" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (skq *SigningKeyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := skq.querySpec()
	_spec.Node.Columns = skq.ctx.Fields
	if len(skq.ctx.Fields) > 0 {
		_spec.Unique = skq.ctx.Unique != nil && *skq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, skq.driver, _spec)
}

func (skq *SigningKeyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(signingkey.Table, signingkey.Columns, sqlgraph.NewFieldSpec(signingkey.FieldID, field.TypeString))
	_spec.From = skq.sql
	if unique := skq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if skq.path != nil {
		_spec.Unique = true
	}
	if fields := skq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, signingkey.FieldID)
		for i := range fields {
			if fields[i] != signingkey.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := skq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := skq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := skq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := skq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (skq *SigningKeyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(skq.driver.Dialect())
	t1 := builder.Table(signingkey.Table)
	columns := skq.ctx.Fields
	if len(columns) == 0 {
		columns = signingkey.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if skq.sql != nil {
		selector = skq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if skq.ctx.Unique != nil && *skq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range skq.predicates {
		p(selector)
	}
	for _, p := range skq.order {
		p(selector)
	}
	if offset := skq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := skq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SigningKeyGroupBy is the group-by builder for SigningKey entities.
type SigningKeyGroupBy struct {
	selector
	build *SigningKeyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (skgb *SigningKeyGroupBy) Aggregate(fns ...AggregateFunc) *SigningKeyGroupBy {
	skgb.fns = append(skgb.fns, fns...)
	return skgb
}

// Scan applies the selector query and scans the result into the given value.
func (skgb *SigningKeyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, skgb.build.ctx, "GroupBy")
	if err := skgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SigningKeyQuery, *SigningKeyGroupBy](ctx, skgb.build, skgb, skgb.build.inters, v)
}

func (skgb *SigningKeyGroupBy) sqlScan(ctx context.Context, root *SigningKeyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(skgb.fns))
	for _, fn := range skgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*skgb.flds)+len(skgb.fns))
		for _, f := range *skgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*skgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := skgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SigningKeySelect is the builder for selecting fields of SigningKey entities.
type SigningKeySelect struct {
	*SigningKeyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sks *SigningKeySelect) Aggregate(fns ...AggregateFunc) *SigningKeySelect {
	sks.fns = append(sks.fns, fns...)
	return sks
}

// Scan applies the selector query and scans the result into the given value.
func (sks *SigningKeySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sks.ctx, "Select")
	if err := sks.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SigningKeyQuery, *SigningKeySelect](ctx, sks.SigningKeyQuery, sks, sks.inters, v)
}

func (sks *SigningKeySelect) sqlScan(ctx context.Context, root *SigningKeyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sks.fns))
	for _, fn := range sks.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sks.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sks.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
