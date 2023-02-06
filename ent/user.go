// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"wallit/ent/user"

	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// IdentityProviderID holds the value of the "identity_provider_id" field.
	IdentityProviderID string `json:"identity_provider_id,omitempty"`
	// OfferFrequency holds the value of the "offer_frequency" field.
	OfferFrequency user.OfferFrequency `json:"offer_frequency,omitempty"`
	// HomeCountry holds the value of the "home_country" field.
	HomeCountry string `json:"home_country,omitempty"`
	// HomeState holds the value of the "home_state" field.
	HomeState string `json:"home_state,omitempty"`
	// HomeCity holds the value of the "home_city" field.
	HomeCity string `json:"home_city,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// PlaidItems holds the value of the plaid_items edge.
	PlaidItems []*PlaidItem `json:"plaid_items,omitempty"`
	// SpendingCategories holds the value of the spending_categories edge.
	SpendingCategories []*SpendingCategory `json:"spending_categories,omitempty"`
	// NotificationChannels holds the value of the notification_channels edge.
	NotificationChannels []*UserNotificationChannelPreferences `json:"notification_channels,omitempty"`
	// Notifications holds the value of the notifications edge.
	Notifications []*Notification `json:"notifications,omitempty"`
	// AvailableDiscountOffers holds the value of the available_discount_offers edge.
	AvailableDiscountOffers []*DiscountOffer `json:"available_discount_offers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
	// totalCount holds the count of the edges above.
	totalCount [5]map[string]int

	namedPlaidItems              map[string][]*PlaidItem
	namedSpendingCategories      map[string][]*SpendingCategory
	namedNotificationChannels    map[string][]*UserNotificationChannelPreferences
	namedNotifications           map[string][]*Notification
	namedAvailableDiscountOffers map[string][]*DiscountOffer
}

// PlaidItemsOrErr returns the PlaidItems value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PlaidItemsOrErr() ([]*PlaidItem, error) {
	if e.loadedTypes[0] {
		return e.PlaidItems, nil
	}
	return nil, &NotLoadedError{edge: "plaid_items"}
}

// SpendingCategoriesOrErr returns the SpendingCategories value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SpendingCategoriesOrErr() ([]*SpendingCategory, error) {
	if e.loadedTypes[1] {
		return e.SpendingCategories, nil
	}
	return nil, &NotLoadedError{edge: "spending_categories"}
}

// NotificationChannelsOrErr returns the NotificationChannels value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) NotificationChannelsOrErr() ([]*UserNotificationChannelPreferences, error) {
	if e.loadedTypes[2] {
		return e.NotificationChannels, nil
	}
	return nil, &NotLoadedError{edge: "notification_channels"}
}

// NotificationsOrErr returns the Notifications value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) NotificationsOrErr() ([]*Notification, error) {
	if e.loadedTypes[3] {
		return e.Notifications, nil
	}
	return nil, &NotLoadedError{edge: "notifications"}
}

// AvailableDiscountOffersOrErr returns the AvailableDiscountOffers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AvailableDiscountOffersOrErr() ([]*DiscountOffer, error) {
	if e.loadedTypes[4] {
		return e.AvailableDiscountOffers, nil
	}
	return nil, &NotLoadedError{edge: "available_discount_offers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldIdentityProviderID, user.FieldOfferFrequency, user.FieldHomeCountry, user.FieldHomeState, user.FieldHomeCity:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldIdentityProviderID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field identity_provider_id", values[i])
			} else if value.Valid {
				u.IdentityProviderID = value.String
			}
		case user.FieldOfferFrequency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field offer_frequency", values[i])
			} else if value.Valid {
				u.OfferFrequency = user.OfferFrequency(value.String)
			}
		case user.FieldHomeCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field home_country", values[i])
			} else if value.Valid {
				u.HomeCountry = value.String
			}
		case user.FieldHomeState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field home_state", values[i])
			} else if value.Valid {
				u.HomeState = value.String
			}
		case user.FieldHomeCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field home_city", values[i])
			} else if value.Valid {
				u.HomeCity = value.String
			}
		}
	}
	return nil
}

// QueryPlaidItems queries the "plaid_items" edge of the User entity.
func (u *User) QueryPlaidItems() *PlaidItemQuery {
	return NewUserClient(u.config).QueryPlaidItems(u)
}

// QuerySpendingCategories queries the "spending_categories" edge of the User entity.
func (u *User) QuerySpendingCategories() *SpendingCategoryQuery {
	return NewUserClient(u.config).QuerySpendingCategories(u)
}

// QueryNotificationChannels queries the "notification_channels" edge of the User entity.
func (u *User) QueryNotificationChannels() *UserNotificationChannelPreferencesQuery {
	return NewUserClient(u.config).QueryNotificationChannels(u)
}

// QueryNotifications queries the "notifications" edge of the User entity.
func (u *User) QueryNotifications() *NotificationQuery {
	return NewUserClient(u.config).QueryNotifications(u)
}

// QueryAvailableDiscountOffers queries the "available_discount_offers" edge of the User entity.
func (u *User) QueryAvailableDiscountOffers() *DiscountOfferQuery {
	return NewUserClient(u.config).QueryAvailableDiscountOffers(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("identity_provider_id=")
	builder.WriteString(u.IdentityProviderID)
	builder.WriteString(", ")
	builder.WriteString("offer_frequency=")
	builder.WriteString(fmt.Sprintf("%v", u.OfferFrequency))
	builder.WriteString(", ")
	builder.WriteString("home_country=")
	builder.WriteString(u.HomeCountry)
	builder.WriteString(", ")
	builder.WriteString("home_state=")
	builder.WriteString(u.HomeState)
	builder.WriteString(", ")
	builder.WriteString("home_city=")
	builder.WriteString(u.HomeCity)
	builder.WriteByte(')')
	return builder.String()
}

// NamedPlaidItems returns the PlaidItems named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedPlaidItems(name string) ([]*PlaidItem, error) {
	if u.Edges.namedPlaidItems == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedPlaidItems[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedPlaidItems(name string, edges ...*PlaidItem) {
	if u.Edges.namedPlaidItems == nil {
		u.Edges.namedPlaidItems = make(map[string][]*PlaidItem)
	}
	if len(edges) == 0 {
		u.Edges.namedPlaidItems[name] = []*PlaidItem{}
	} else {
		u.Edges.namedPlaidItems[name] = append(u.Edges.namedPlaidItems[name], edges...)
	}
}

// NamedSpendingCategories returns the SpendingCategories named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedSpendingCategories(name string) ([]*SpendingCategory, error) {
	if u.Edges.namedSpendingCategories == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedSpendingCategories[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedSpendingCategories(name string, edges ...*SpendingCategory) {
	if u.Edges.namedSpendingCategories == nil {
		u.Edges.namedSpendingCategories = make(map[string][]*SpendingCategory)
	}
	if len(edges) == 0 {
		u.Edges.namedSpendingCategories[name] = []*SpendingCategory{}
	} else {
		u.Edges.namedSpendingCategories[name] = append(u.Edges.namedSpendingCategories[name], edges...)
	}
}

// NamedNotificationChannels returns the NotificationChannels named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedNotificationChannels(name string) ([]*UserNotificationChannelPreferences, error) {
	if u.Edges.namedNotificationChannels == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedNotificationChannels[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedNotificationChannels(name string, edges ...*UserNotificationChannelPreferences) {
	if u.Edges.namedNotificationChannels == nil {
		u.Edges.namedNotificationChannels = make(map[string][]*UserNotificationChannelPreferences)
	}
	if len(edges) == 0 {
		u.Edges.namedNotificationChannels[name] = []*UserNotificationChannelPreferences{}
	} else {
		u.Edges.namedNotificationChannels[name] = append(u.Edges.namedNotificationChannels[name], edges...)
	}
}

// NamedNotifications returns the Notifications named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedNotifications(name string) ([]*Notification, error) {
	if u.Edges.namedNotifications == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedNotifications[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedNotifications(name string, edges ...*Notification) {
	if u.Edges.namedNotifications == nil {
		u.Edges.namedNotifications = make(map[string][]*Notification)
	}
	if len(edges) == 0 {
		u.Edges.namedNotifications[name] = []*Notification{}
	} else {
		u.Edges.namedNotifications[name] = append(u.Edges.namedNotifications[name], edges...)
	}
}

// NamedAvailableDiscountOffers returns the AvailableDiscountOffers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedAvailableDiscountOffers(name string) ([]*DiscountOffer, error) {
	if u.Edges.namedAvailableDiscountOffers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedAvailableDiscountOffers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedAvailableDiscountOffers(name string, edges ...*DiscountOffer) {
	if u.Edges.namedAvailableDiscountOffers == nil {
		u.Edges.namedAvailableDiscountOffers = make(map[string][]*DiscountOffer)
	}
	if len(edges) == 0 {
		u.Edges.namedAvailableDiscountOffers[name] = []*DiscountOffer{}
	} else {
		u.Edges.namedAvailableDiscountOffers[name] = append(u.Edges.namedAvailableDiscountOffers[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
