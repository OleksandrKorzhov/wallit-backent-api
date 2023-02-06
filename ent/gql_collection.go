// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"wallit/ent/plaidinstitutionaccount"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (do *DiscountOfferQuery) CollectFields(ctx context.Context, satisfies ...string) (*DiscountOfferQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return do, nil
	}
	if err := do.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return do, nil
}

func (do *DiscountOfferQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "ownerMerchant":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&MerchantClient{config: do.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			do.withOwnerMerchant = query
		case "discountEligibleUsers":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: do.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			do.WithNamedDiscountEligibleUsers(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "discountOfferNotification":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&NotificationClient{config: do.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			do.withDiscountOfferNotification = query
		}
	}
	return nil
}

type discountofferPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []DiscountOfferPaginateOption
}

func newDiscountOfferPaginateArgs(rv map[string]interface{}) *discountofferPaginateArgs {
	args := &discountofferPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &DiscountOfferOrder{Field: &DiscountOfferOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithDiscountOfferOrder(order))
			}
		case *DiscountOfferOrder:
			if v != nil {
				args.opts = append(args.opts, WithDiscountOfferOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*DiscountOfferWhereInput); ok {
		args.opts = append(args.opts, WithDiscountOfferFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (m *MerchantQuery) CollectFields(ctx context.Context, satisfies ...string) (*MerchantQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return m, nil
	}
	if err := m.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *MerchantQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "discountOffers":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&DiscountOfferClient{config: m.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			m.WithNamedDiscountOffers(alias, func(wq *DiscountOfferQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type merchantPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []MerchantPaginateOption
}

func newMerchantPaginateArgs(rv map[string]interface{}) *merchantPaginateArgs {
	args := &merchantPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*MerchantWhereInput); ok {
		args.opts = append(args.opts, WithMerchantFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (n *NotificationQuery) CollectFields(ctx context.Context, satisfies ...string) (*NotificationQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return n, nil
	}
	if err := n.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return n, nil
}

func (n *NotificationQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "notificationRecipient":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: n.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			n.withNotificationRecipient = query
		case "notificationDiscountOffer":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&DiscountOfferClient{config: n.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			n.withNotificationDiscountOffer = query
		}
	}
	return nil
}

type notificationPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []NotificationPaginateOption
}

func newNotificationPaginateArgs(rv map[string]interface{}) *notificationPaginateArgs {
	args := &notificationPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &NotificationOrder{Field: &NotificationOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithNotificationOrder(order))
			}
		case *NotificationOrder:
			if v != nil {
				args.opts = append(args.opts, WithNotificationOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*NotificationWhereInput); ok {
		args.opts = append(args.opts, WithNotificationFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pi *PlaidInstitutionQuery) CollectFields(ctx context.Context, satisfies ...string) (*PlaidInstitutionQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pi, nil
	}
	if err := pi.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pi, nil
}

func (pi *PlaidInstitutionQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "plaidItem":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PlaidItemClient{config: pi.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			pi.withPlaidItem = query
		case "accounts":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PlaidInstitutionAccountClient{config: pi.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			pi.WithNamedAccounts(alias, func(wq *PlaidInstitutionAccountQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type plaidinstitutionPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PlaidInstitutionPaginateOption
}

func newPlaidInstitutionPaginateArgs(rv map[string]interface{}) *plaidinstitutionPaginateArgs {
	args := &plaidinstitutionPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*PlaidInstitutionWhereInput); ok {
		args.opts = append(args.opts, WithPlaidInstitutionFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pia *PlaidInstitutionAccountQuery) CollectFields(ctx context.Context, satisfies ...string) (*PlaidInstitutionAccountQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pia, nil
	}
	if err := pia.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pia, nil
}

func (pia *PlaidInstitutionAccountQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "parentInstitution":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PlaidInstitutionClient{config: pia.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			pia.withParentInstitution = query
		case "transactions":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TransactionClient{config: pia.config}).Query()
			)
			args := newTransactionPaginateArgs(fieldArgs(ctx, new(TransactionWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newTransactionPager(args.opts)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					pia.loadTotal = append(pia.loadTotal, func(ctx context.Context, nodes []*PlaidInstitutionAccount) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"plaid_institution_account_transactions"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(plaidinstitutionaccount.TransactionsColumn, ids...))
						})
						if err := query.GroupBy(plaidinstitutionaccount.TransactionsColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				} else {
					pia.loadTotal = append(pia.loadTotal, func(_ context.Context, nodes []*PlaidInstitutionAccount) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Transactions)
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}

			query = pager.applyCursors(query, args.after, args.before)
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(plaidinstitutionaccount.TransactionsColumn, limit, pager.orderExpr(args.last != nil))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query, args.last != nil)
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, op, *field, path, satisfies...); err != nil {
					return err
				}
			}
			pia.WithNamedTransactions(alias, func(wq *TransactionQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type plaidinstitutionaccountPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PlaidInstitutionAccountPaginateOption
}

func newPlaidInstitutionAccountPaginateArgs(rv map[string]interface{}) *plaidinstitutionaccountPaginateArgs {
	args := &plaidinstitutionaccountPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*PlaidInstitutionAccountWhereInput); ok {
		args.opts = append(args.opts, WithPlaidInstitutionAccountFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pi *PlaidItemQuery) CollectFields(ctx context.Context, satisfies ...string) (*PlaidItemQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pi, nil
	}
	if err := pi.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pi, nil
}

func (pi *PlaidItemQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "owner":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: pi.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			pi.withOwner = query
		case "transactionSyncs":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TransactionSyncClient{config: pi.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			pi.WithNamedTransactionSyncs(alias, func(wq *TransactionSyncQuery) {
				*wq = *query
			})
		case "institution":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PlaidInstitutionClient{config: pi.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			pi.withInstitution = query
		}
	}
	return nil
}

type plaiditemPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PlaidItemPaginateOption
}

func newPlaidItemPaginateArgs(rv map[string]interface{}) *plaiditemPaginateArgs {
	args := &plaiditemPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &PlaidItemOrder{Field: &PlaidItemOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithPlaidItemOrder(order))
			}
		case *PlaidItemOrder:
			if v != nil {
				args.opts = append(args.opts, WithPlaidItemOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*PlaidItemWhereInput); ok {
		args.opts = append(args.opts, WithPlaidItemFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (sc *SpendingCategoryQuery) CollectFields(ctx context.Context, satisfies ...string) (*SpendingCategoryQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return sc, nil
	}
	if err := sc.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return sc, nil
}

func (sc *SpendingCategoryQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "categoryTransactions":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TransactionClient{config: sc.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			sc.WithNamedCategoryTransactions(alias, func(wq *TransactionQuery) {
				*wq = *query
			})
		case "interestedUsers":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: sc.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			sc.WithNamedInterestedUsers(alias, func(wq *UserQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type spendingcategoryPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []SpendingCategoryPaginateOption
}

func newSpendingCategoryPaginateArgs(rv map[string]interface{}) *spendingcategoryPaginateArgs {
	args := &spendingcategoryPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*SpendingCategoryWhereInput); ok {
		args.opts = append(args.opts, WithSpendingCategoryFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TransactionQuery) CollectFields(ctx context.Context, satisfies ...string) (*TransactionQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TransactionQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "institutionAccount":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PlaidInstitutionAccountClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			t.withInstitutionAccount = query
		case "transactionCategories":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&SpendingCategoryClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			t.WithNamedTransactionCategories(alias, func(wq *SpendingCategoryQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type transactionPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TransactionPaginateOption
}

func newTransactionPaginateArgs(rv map[string]interface{}) *transactionPaginateArgs {
	args := &transactionPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &TransactionOrder{Field: &TransactionOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithTransactionOrder(order))
			}
		case *TransactionOrder:
			if v != nil {
				args.opts = append(args.opts, WithTransactionOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*TransactionWhereInput); ok {
		args.opts = append(args.opts, WithTransactionFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ts *TransactionSyncQuery) CollectFields(ctx context.Context, satisfies ...string) (*TransactionSyncQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ts, nil
	}
	if err := ts.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ts, nil
}

func (ts *TransactionSyncQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "item":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PlaidItemClient{config: ts.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			ts.withItem = query
		}
	}
	return nil
}

type transactionsyncPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TransactionSyncPaginateOption
}

func newTransactionSyncPaginateArgs(rv map[string]interface{}) *transactionsyncPaginateArgs {
	args := &transactionsyncPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*TransactionSyncWhereInput); ok {
		args.opts = append(args.opts, WithTransactionSyncFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "plaidItems":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PlaidItemClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedPlaidItems(alias, func(wq *PlaidItemQuery) {
				*wq = *query
			})
		case "spendingCategories":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&SpendingCategoryClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedSpendingCategories(alias, func(wq *SpendingCategoryQuery) {
				*wq = *query
			})
		case "notificationChannels":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserNotificationChannelPreferencesClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedNotificationChannels(alias, func(wq *UserNotificationChannelPreferencesQuery) {
				*wq = *query
			})
		case "notifications":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&NotificationClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedNotifications(alias, func(wq *NotificationQuery) {
				*wq = *query
			})
		case "availableDiscountOffers":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&DiscountOfferClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedAvailableDiscountOffers(alias, func(wq *DiscountOfferQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]interface{}) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (uncp *UserNotificationChannelPreferencesQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserNotificationChannelPreferencesQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return uncp, nil
	}
	if err := uncp.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return uncp, nil
}

func (uncp *UserNotificationChannelPreferencesQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "chanelUsers":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: uncp.config}).Query()
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			uncp.withChanelUsers = query
		}
	}
	return nil
}

type usernotificationchannelpreferencesPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserNotificationChannelPreferencesPaginateOption
}

func newUserNotificationChannelPreferencesPaginateArgs(rv map[string]interface{}) *usernotificationchannelpreferencesPaginateArgs {
	args := &usernotificationchannelpreferencesPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*UserNotificationChannelPreferencesWhereInput); ok {
		args.opts = append(args.opts, WithUserNotificationChannelPreferencesFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput interface{}, path ...string) map[string]interface{} {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	for _, name := range path {
		var field *graphql.CollectedField
		for _, f := range graphql.CollectFields(oc, fc.Field.Selections, nil) {
			if f.Alias == name {
				field = &f
				break
			}
		}
		if field == nil {
			return nil
		}
		cf, err := fc.Child(ctx, *field)
		if err != nil {
			args := field.ArgumentMap(oc.Variables)
			return unmarshalArgs(ctx, whereInput, args)
		}
		fc = cf
	}
	return fc.Args
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput interface{}, args map[string]interface{}) map[string]interface{} {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}