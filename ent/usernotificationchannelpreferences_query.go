// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"wallit/ent/predicate"
	"wallit/ent/user"
	"wallit/ent/usernotificationchannelpreferences"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserNotificationChannelPreferencesQuery is the builder for querying UserNotificationChannelPreferences entities.
type UserNotificationChannelPreferencesQuery struct {
	config
	ctx             *QueryContext
	order           []OrderFunc
	inters          []Interceptor
	predicates      []predicate.UserNotificationChannelPreferences
	withChanelUsers *UserQuery
	withFKs         bool
	modifiers       []func(*sql.Selector)
	loadTotal       []func(context.Context, []*UserNotificationChannelPreferences) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserNotificationChannelPreferencesQuery builder.
func (uncpq *UserNotificationChannelPreferencesQuery) Where(ps ...predicate.UserNotificationChannelPreferences) *UserNotificationChannelPreferencesQuery {
	uncpq.predicates = append(uncpq.predicates, ps...)
	return uncpq
}

// Limit the number of records to be returned by this query.
func (uncpq *UserNotificationChannelPreferencesQuery) Limit(limit int) *UserNotificationChannelPreferencesQuery {
	uncpq.ctx.Limit = &limit
	return uncpq
}

// Offset to start from.
func (uncpq *UserNotificationChannelPreferencesQuery) Offset(offset int) *UserNotificationChannelPreferencesQuery {
	uncpq.ctx.Offset = &offset
	return uncpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uncpq *UserNotificationChannelPreferencesQuery) Unique(unique bool) *UserNotificationChannelPreferencesQuery {
	uncpq.ctx.Unique = &unique
	return uncpq
}

// Order specifies how the records should be ordered.
func (uncpq *UserNotificationChannelPreferencesQuery) Order(o ...OrderFunc) *UserNotificationChannelPreferencesQuery {
	uncpq.order = append(uncpq.order, o...)
	return uncpq
}

// QueryChanelUsers chains the current query on the "chanel_users" edge.
func (uncpq *UserNotificationChannelPreferencesQuery) QueryChanelUsers() *UserQuery {
	query := (&UserClient{config: uncpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uncpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uncpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usernotificationchannelpreferences.Table, usernotificationchannelpreferences.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, usernotificationchannelpreferences.ChanelUsersTable, usernotificationchannelpreferences.ChanelUsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(uncpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserNotificationChannelPreferences entity from the query.
// Returns a *NotFoundError when no UserNotificationChannelPreferences was found.
func (uncpq *UserNotificationChannelPreferencesQuery) First(ctx context.Context) (*UserNotificationChannelPreferences, error) {
	nodes, err := uncpq.Limit(1).All(setContextOp(ctx, uncpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usernotificationchannelpreferences.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uncpq *UserNotificationChannelPreferencesQuery) FirstX(ctx context.Context) *UserNotificationChannelPreferences {
	node, err := uncpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserNotificationChannelPreferences ID from the query.
// Returns a *NotFoundError when no UserNotificationChannelPreferences ID was found.
func (uncpq *UserNotificationChannelPreferencesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uncpq.Limit(1).IDs(setContextOp(ctx, uncpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usernotificationchannelpreferences.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uncpq *UserNotificationChannelPreferencesQuery) FirstIDX(ctx context.Context) int {
	id, err := uncpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserNotificationChannelPreferences entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserNotificationChannelPreferences entity is found.
// Returns a *NotFoundError when no UserNotificationChannelPreferences entities are found.
func (uncpq *UserNotificationChannelPreferencesQuery) Only(ctx context.Context) (*UserNotificationChannelPreferences, error) {
	nodes, err := uncpq.Limit(2).All(setContextOp(ctx, uncpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usernotificationchannelpreferences.Label}
	default:
		return nil, &NotSingularError{usernotificationchannelpreferences.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uncpq *UserNotificationChannelPreferencesQuery) OnlyX(ctx context.Context) *UserNotificationChannelPreferences {
	node, err := uncpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserNotificationChannelPreferences ID in the query.
// Returns a *NotSingularError when more than one UserNotificationChannelPreferences ID is found.
// Returns a *NotFoundError when no entities are found.
func (uncpq *UserNotificationChannelPreferencesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uncpq.Limit(2).IDs(setContextOp(ctx, uncpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usernotificationchannelpreferences.Label}
	default:
		err = &NotSingularError{usernotificationchannelpreferences.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uncpq *UserNotificationChannelPreferencesQuery) OnlyIDX(ctx context.Context) int {
	id, err := uncpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserNotificationChannelPreferencesSlice.
func (uncpq *UserNotificationChannelPreferencesQuery) All(ctx context.Context) ([]*UserNotificationChannelPreferences, error) {
	ctx = setContextOp(ctx, uncpq.ctx, "All")
	if err := uncpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserNotificationChannelPreferences, *UserNotificationChannelPreferencesQuery]()
	return withInterceptors[[]*UserNotificationChannelPreferences](ctx, uncpq, qr, uncpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uncpq *UserNotificationChannelPreferencesQuery) AllX(ctx context.Context) []*UserNotificationChannelPreferences {
	nodes, err := uncpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserNotificationChannelPreferences IDs.
func (uncpq *UserNotificationChannelPreferencesQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = setContextOp(ctx, uncpq.ctx, "IDs")
	if err := uncpq.Select(usernotificationchannelpreferences.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uncpq *UserNotificationChannelPreferencesQuery) IDsX(ctx context.Context) []int {
	ids, err := uncpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uncpq *UserNotificationChannelPreferencesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uncpq.ctx, "Count")
	if err := uncpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uncpq, querierCount[*UserNotificationChannelPreferencesQuery](), uncpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uncpq *UserNotificationChannelPreferencesQuery) CountX(ctx context.Context) int {
	count, err := uncpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uncpq *UserNotificationChannelPreferencesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uncpq.ctx, "Exist")
	switch _, err := uncpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uncpq *UserNotificationChannelPreferencesQuery) ExistX(ctx context.Context) bool {
	exist, err := uncpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserNotificationChannelPreferencesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uncpq *UserNotificationChannelPreferencesQuery) Clone() *UserNotificationChannelPreferencesQuery {
	if uncpq == nil {
		return nil
	}
	return &UserNotificationChannelPreferencesQuery{
		config:          uncpq.config,
		ctx:             uncpq.ctx.Clone(),
		order:           append([]OrderFunc{}, uncpq.order...),
		inters:          append([]Interceptor{}, uncpq.inters...),
		predicates:      append([]predicate.UserNotificationChannelPreferences{}, uncpq.predicates...),
		withChanelUsers: uncpq.withChanelUsers.Clone(),
		// clone intermediate query.
		sql:  uncpq.sql.Clone(),
		path: uncpq.path,
	}
}

// WithChanelUsers tells the query-builder to eager-load the nodes that are connected to
// the "chanel_users" edge. The optional arguments are used to configure the query builder of the edge.
func (uncpq *UserNotificationChannelPreferencesQuery) WithChanelUsers(opts ...func(*UserQuery)) *UserNotificationChannelPreferencesQuery {
	query := (&UserClient{config: uncpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uncpq.withChanelUsers = query
	return uncpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Chanel usernotificationchannelpreferences.Chanel `json:"chanel,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserNotificationChannelPreferences.Query().
//		GroupBy(usernotificationchannelpreferences.FieldChanel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uncpq *UserNotificationChannelPreferencesQuery) GroupBy(field string, fields ...string) *UserNotificationChannelPreferencesGroupBy {
	uncpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserNotificationChannelPreferencesGroupBy{build: uncpq}
	grbuild.flds = &uncpq.ctx.Fields
	grbuild.label = usernotificationchannelpreferences.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Chanel usernotificationchannelpreferences.Chanel `json:"chanel,omitempty"`
//	}
//
//	client.UserNotificationChannelPreferences.Query().
//		Select(usernotificationchannelpreferences.FieldChanel).
//		Scan(ctx, &v)
func (uncpq *UserNotificationChannelPreferencesQuery) Select(fields ...string) *UserNotificationChannelPreferencesSelect {
	uncpq.ctx.Fields = append(uncpq.ctx.Fields, fields...)
	sbuild := &UserNotificationChannelPreferencesSelect{UserNotificationChannelPreferencesQuery: uncpq}
	sbuild.label = usernotificationchannelpreferences.Label
	sbuild.flds, sbuild.scan = &uncpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserNotificationChannelPreferencesSelect configured with the given aggregations.
func (uncpq *UserNotificationChannelPreferencesQuery) Aggregate(fns ...AggregateFunc) *UserNotificationChannelPreferencesSelect {
	return uncpq.Select().Aggregate(fns...)
}

func (uncpq *UserNotificationChannelPreferencesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uncpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uncpq); err != nil {
				return err
			}
		}
	}
	for _, f := range uncpq.ctx.Fields {
		if !usernotificationchannelpreferences.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uncpq.path != nil {
		prev, err := uncpq.path(ctx)
		if err != nil {
			return err
		}
		uncpq.sql = prev
	}
	return nil
}

func (uncpq *UserNotificationChannelPreferencesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserNotificationChannelPreferences, error) {
	var (
		nodes       = []*UserNotificationChannelPreferences{}
		withFKs     = uncpq.withFKs
		_spec       = uncpq.querySpec()
		loadedTypes = [1]bool{
			uncpq.withChanelUsers != nil,
		}
	)
	if uncpq.withChanelUsers != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, usernotificationchannelpreferences.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserNotificationChannelPreferences).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserNotificationChannelPreferences{config: uncpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(uncpq.modifiers) > 0 {
		_spec.Modifiers = uncpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uncpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uncpq.withChanelUsers; query != nil {
		if err := uncpq.loadChanelUsers(ctx, query, nodes, nil,
			func(n *UserNotificationChannelPreferences, e *User) { n.Edges.ChanelUsers = e }); err != nil {
			return nil, err
		}
	}
	for i := range uncpq.loadTotal {
		if err := uncpq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uncpq *UserNotificationChannelPreferencesQuery) loadChanelUsers(ctx context.Context, query *UserQuery, nodes []*UserNotificationChannelPreferences, init func(*UserNotificationChannelPreferences), assign func(*UserNotificationChannelPreferences, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserNotificationChannelPreferences)
	for i := range nodes {
		if nodes[i].user_notification_channels == nil {
			continue
		}
		fk := *nodes[i].user_notification_channels
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_notification_channels" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (uncpq *UserNotificationChannelPreferencesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uncpq.querySpec()
	if len(uncpq.modifiers) > 0 {
		_spec.Modifiers = uncpq.modifiers
	}
	_spec.Node.Columns = uncpq.ctx.Fields
	if len(uncpq.ctx.Fields) > 0 {
		_spec.Unique = uncpq.ctx.Unique != nil && *uncpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uncpq.driver, _spec)
}

func (uncpq *UserNotificationChannelPreferencesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usernotificationchannelpreferences.Table,
			Columns: usernotificationchannelpreferences.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usernotificationchannelpreferences.FieldID,
			},
		},
		From:   uncpq.sql,
		Unique: true,
	}
	if unique := uncpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uncpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usernotificationchannelpreferences.FieldID)
		for i := range fields {
			if fields[i] != usernotificationchannelpreferences.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uncpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uncpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uncpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uncpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uncpq *UserNotificationChannelPreferencesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uncpq.driver.Dialect())
	t1 := builder.Table(usernotificationchannelpreferences.Table)
	columns := uncpq.ctx.Fields
	if len(columns) == 0 {
		columns = usernotificationchannelpreferences.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uncpq.sql != nil {
		selector = uncpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uncpq.ctx.Unique != nil && *uncpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uncpq.predicates {
		p(selector)
	}
	for _, p := range uncpq.order {
		p(selector)
	}
	if offset := uncpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uncpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserNotificationChannelPreferencesGroupBy is the group-by builder for UserNotificationChannelPreferences entities.
type UserNotificationChannelPreferencesGroupBy struct {
	selector
	build *UserNotificationChannelPreferencesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uncpgb *UserNotificationChannelPreferencesGroupBy) Aggregate(fns ...AggregateFunc) *UserNotificationChannelPreferencesGroupBy {
	uncpgb.fns = append(uncpgb.fns, fns...)
	return uncpgb
}

// Scan applies the selector query and scans the result into the given value.
func (uncpgb *UserNotificationChannelPreferencesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uncpgb.build.ctx, "GroupBy")
	if err := uncpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserNotificationChannelPreferencesQuery, *UserNotificationChannelPreferencesGroupBy](ctx, uncpgb.build, uncpgb, uncpgb.build.inters, v)
}

func (uncpgb *UserNotificationChannelPreferencesGroupBy) sqlScan(ctx context.Context, root *UserNotificationChannelPreferencesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(uncpgb.fns))
	for _, fn := range uncpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*uncpgb.flds)+len(uncpgb.fns))
		for _, f := range *uncpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*uncpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uncpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserNotificationChannelPreferencesSelect is the builder for selecting fields of UserNotificationChannelPreferences entities.
type UserNotificationChannelPreferencesSelect struct {
	*UserNotificationChannelPreferencesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uncps *UserNotificationChannelPreferencesSelect) Aggregate(fns ...AggregateFunc) *UserNotificationChannelPreferencesSelect {
	uncps.fns = append(uncps.fns, fns...)
	return uncps
}

// Scan applies the selector query and scans the result into the given value.
func (uncps *UserNotificationChannelPreferencesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uncps.ctx, "Select")
	if err := uncps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserNotificationChannelPreferencesQuery, *UserNotificationChannelPreferencesSelect](ctx, uncps.UserNotificationChannelPreferencesQuery, uncps, uncps.inters, v)
}

func (uncps *UserNotificationChannelPreferencesSelect) sqlScan(ctx context.Context, root *UserNotificationChannelPreferencesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uncps.fns))
	for _, fn := range uncps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uncps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uncps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}