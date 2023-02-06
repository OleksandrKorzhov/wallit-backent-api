// Code generated by ent, DO NOT EDIT.

package user

import (
	"wallit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// IdentityProviderID applies equality check predicate on the "identity_provider_id" field. It's identical to IdentityProviderIDEQ.
func IdentityProviderID(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIdentityProviderID, v))
}

// HomeCountry applies equality check predicate on the "home_country" field. It's identical to HomeCountryEQ.
func HomeCountry(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHomeCountry, v))
}

// HomeState applies equality check predicate on the "home_state" field. It's identical to HomeStateEQ.
func HomeState(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHomeState, v))
}

// HomeCity applies equality check predicate on the "home_city" field. It's identical to HomeCityEQ.
func HomeCity(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHomeCity, v))
}

// IdentityProviderIDEQ applies the EQ predicate on the "identity_provider_id" field.
func IdentityProviderIDEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldIdentityProviderID, v))
}

// IdentityProviderIDNEQ applies the NEQ predicate on the "identity_provider_id" field.
func IdentityProviderIDNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldIdentityProviderID, v))
}

// IdentityProviderIDIn applies the In predicate on the "identity_provider_id" field.
func IdentityProviderIDIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldIdentityProviderID, vs...))
}

// IdentityProviderIDNotIn applies the NotIn predicate on the "identity_provider_id" field.
func IdentityProviderIDNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldIdentityProviderID, vs...))
}

// IdentityProviderIDGT applies the GT predicate on the "identity_provider_id" field.
func IdentityProviderIDGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldIdentityProviderID, v))
}

// IdentityProviderIDGTE applies the GTE predicate on the "identity_provider_id" field.
func IdentityProviderIDGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldIdentityProviderID, v))
}

// IdentityProviderIDLT applies the LT predicate on the "identity_provider_id" field.
func IdentityProviderIDLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldIdentityProviderID, v))
}

// IdentityProviderIDLTE applies the LTE predicate on the "identity_provider_id" field.
func IdentityProviderIDLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldIdentityProviderID, v))
}

// IdentityProviderIDContains applies the Contains predicate on the "identity_provider_id" field.
func IdentityProviderIDContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldIdentityProviderID, v))
}

// IdentityProviderIDHasPrefix applies the HasPrefix predicate on the "identity_provider_id" field.
func IdentityProviderIDHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldIdentityProviderID, v))
}

// IdentityProviderIDHasSuffix applies the HasSuffix predicate on the "identity_provider_id" field.
func IdentityProviderIDHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldIdentityProviderID, v))
}

// IdentityProviderIDIsNil applies the IsNil predicate on the "identity_provider_id" field.
func IdentityProviderIDIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldIdentityProviderID))
}

// IdentityProviderIDNotNil applies the NotNil predicate on the "identity_provider_id" field.
func IdentityProviderIDNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldIdentityProviderID))
}

// IdentityProviderIDEqualFold applies the EqualFold predicate on the "identity_provider_id" field.
func IdentityProviderIDEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldIdentityProviderID, v))
}

// IdentityProviderIDContainsFold applies the ContainsFold predicate on the "identity_provider_id" field.
func IdentityProviderIDContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldIdentityProviderID, v))
}

// OfferFrequencyEQ applies the EQ predicate on the "offer_frequency" field.
func OfferFrequencyEQ(v OfferFrequency) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOfferFrequency, v))
}

// OfferFrequencyNEQ applies the NEQ predicate on the "offer_frequency" field.
func OfferFrequencyNEQ(v OfferFrequency) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldOfferFrequency, v))
}

// OfferFrequencyIn applies the In predicate on the "offer_frequency" field.
func OfferFrequencyIn(vs ...OfferFrequency) predicate.User {
	return predicate.User(sql.FieldIn(FieldOfferFrequency, vs...))
}

// OfferFrequencyNotIn applies the NotIn predicate on the "offer_frequency" field.
func OfferFrequencyNotIn(vs ...OfferFrequency) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldOfferFrequency, vs...))
}

// HomeCountryEQ applies the EQ predicate on the "home_country" field.
func HomeCountryEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHomeCountry, v))
}

// HomeCountryNEQ applies the NEQ predicate on the "home_country" field.
func HomeCountryNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldHomeCountry, v))
}

// HomeCountryIn applies the In predicate on the "home_country" field.
func HomeCountryIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldHomeCountry, vs...))
}

// HomeCountryNotIn applies the NotIn predicate on the "home_country" field.
func HomeCountryNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldHomeCountry, vs...))
}

// HomeCountryGT applies the GT predicate on the "home_country" field.
func HomeCountryGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldHomeCountry, v))
}

// HomeCountryGTE applies the GTE predicate on the "home_country" field.
func HomeCountryGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldHomeCountry, v))
}

// HomeCountryLT applies the LT predicate on the "home_country" field.
func HomeCountryLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldHomeCountry, v))
}

// HomeCountryLTE applies the LTE predicate on the "home_country" field.
func HomeCountryLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldHomeCountry, v))
}

// HomeCountryContains applies the Contains predicate on the "home_country" field.
func HomeCountryContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldHomeCountry, v))
}

// HomeCountryHasPrefix applies the HasPrefix predicate on the "home_country" field.
func HomeCountryHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldHomeCountry, v))
}

// HomeCountryHasSuffix applies the HasSuffix predicate on the "home_country" field.
func HomeCountryHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldHomeCountry, v))
}

// HomeCountryIsNil applies the IsNil predicate on the "home_country" field.
func HomeCountryIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldHomeCountry))
}

// HomeCountryNotNil applies the NotNil predicate on the "home_country" field.
func HomeCountryNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldHomeCountry))
}

// HomeCountryEqualFold applies the EqualFold predicate on the "home_country" field.
func HomeCountryEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldHomeCountry, v))
}

// HomeCountryContainsFold applies the ContainsFold predicate on the "home_country" field.
func HomeCountryContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldHomeCountry, v))
}

// HomeStateEQ applies the EQ predicate on the "home_state" field.
func HomeStateEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHomeState, v))
}

// HomeStateNEQ applies the NEQ predicate on the "home_state" field.
func HomeStateNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldHomeState, v))
}

// HomeStateIn applies the In predicate on the "home_state" field.
func HomeStateIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldHomeState, vs...))
}

// HomeStateNotIn applies the NotIn predicate on the "home_state" field.
func HomeStateNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldHomeState, vs...))
}

// HomeStateGT applies the GT predicate on the "home_state" field.
func HomeStateGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldHomeState, v))
}

// HomeStateGTE applies the GTE predicate on the "home_state" field.
func HomeStateGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldHomeState, v))
}

// HomeStateLT applies the LT predicate on the "home_state" field.
func HomeStateLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldHomeState, v))
}

// HomeStateLTE applies the LTE predicate on the "home_state" field.
func HomeStateLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldHomeState, v))
}

// HomeStateContains applies the Contains predicate on the "home_state" field.
func HomeStateContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldHomeState, v))
}

// HomeStateHasPrefix applies the HasPrefix predicate on the "home_state" field.
func HomeStateHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldHomeState, v))
}

// HomeStateHasSuffix applies the HasSuffix predicate on the "home_state" field.
func HomeStateHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldHomeState, v))
}

// HomeStateIsNil applies the IsNil predicate on the "home_state" field.
func HomeStateIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldHomeState))
}

// HomeStateNotNil applies the NotNil predicate on the "home_state" field.
func HomeStateNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldHomeState))
}

// HomeStateEqualFold applies the EqualFold predicate on the "home_state" field.
func HomeStateEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldHomeState, v))
}

// HomeStateContainsFold applies the ContainsFold predicate on the "home_state" field.
func HomeStateContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldHomeState, v))
}

// HomeCityEQ applies the EQ predicate on the "home_city" field.
func HomeCityEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldHomeCity, v))
}

// HomeCityNEQ applies the NEQ predicate on the "home_city" field.
func HomeCityNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldHomeCity, v))
}

// HomeCityIn applies the In predicate on the "home_city" field.
func HomeCityIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldHomeCity, vs...))
}

// HomeCityNotIn applies the NotIn predicate on the "home_city" field.
func HomeCityNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldHomeCity, vs...))
}

// HomeCityGT applies the GT predicate on the "home_city" field.
func HomeCityGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldHomeCity, v))
}

// HomeCityGTE applies the GTE predicate on the "home_city" field.
func HomeCityGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldHomeCity, v))
}

// HomeCityLT applies the LT predicate on the "home_city" field.
func HomeCityLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldHomeCity, v))
}

// HomeCityLTE applies the LTE predicate on the "home_city" field.
func HomeCityLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldHomeCity, v))
}

// HomeCityContains applies the Contains predicate on the "home_city" field.
func HomeCityContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldHomeCity, v))
}

// HomeCityHasPrefix applies the HasPrefix predicate on the "home_city" field.
func HomeCityHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldHomeCity, v))
}

// HomeCityHasSuffix applies the HasSuffix predicate on the "home_city" field.
func HomeCityHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldHomeCity, v))
}

// HomeCityIsNil applies the IsNil predicate on the "home_city" field.
func HomeCityIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldHomeCity))
}

// HomeCityNotNil applies the NotNil predicate on the "home_city" field.
func HomeCityNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldHomeCity))
}

// HomeCityEqualFold applies the EqualFold predicate on the "home_city" field.
func HomeCityEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldHomeCity, v))
}

// HomeCityContainsFold applies the ContainsFold predicate on the "home_city" field.
func HomeCityContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldHomeCity, v))
}

// HasPlaidItems applies the HasEdge predicate on the "plaid_items" edge.
func HasPlaidItems() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PlaidItemsTable, PlaidItemsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlaidItemsWith applies the HasEdge predicate on the "plaid_items" edge with a given conditions (other predicates).
func HasPlaidItemsWith(preds ...predicate.PlaidItem) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlaidItemsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PlaidItemsTable, PlaidItemsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSpendingCategories applies the HasEdge predicate on the "spending_categories" edge.
func HasSpendingCategories() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SpendingCategoriesTable, SpendingCategoriesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSpendingCategoriesWith applies the HasEdge predicate on the "spending_categories" edge with a given conditions (other predicates).
func HasSpendingCategoriesWith(preds ...predicate.SpendingCategory) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SpendingCategoriesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SpendingCategoriesTable, SpendingCategoriesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotificationChannels applies the HasEdge predicate on the "notification_channels" edge.
func HasNotificationChannels() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NotificationChannelsTable, NotificationChannelsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotificationChannelsWith applies the HasEdge predicate on the "notification_channels" edge with a given conditions (other predicates).
func HasNotificationChannelsWith(preds ...predicate.UserNotificationChannelPreferences) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NotificationChannelsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NotificationChannelsTable, NotificationChannelsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotifications applies the HasEdge predicate on the "notifications" edge.
func HasNotifications() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NotificationsTable, NotificationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotificationsWith applies the HasEdge predicate on the "notifications" edge with a given conditions (other predicates).
func HasNotificationsWith(preds ...predicate.Notification) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NotificationsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NotificationsTable, NotificationsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAvailableDiscountOffers applies the HasEdge predicate on the "available_discount_offers" edge.
func HasAvailableDiscountOffers() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AvailableDiscountOffersTable, AvailableDiscountOffersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAvailableDiscountOffersWith applies the HasEdge predicate on the "available_discount_offers" edge with a given conditions (other predicates).
func HasAvailableDiscountOffersWith(preds ...predicate.DiscountOffer) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AvailableDiscountOffersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AvailableDiscountOffersTable, AvailableDiscountOffersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}