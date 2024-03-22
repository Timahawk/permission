// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// EdgeWithin holds the string denoting the within edge name in mutations.
	EdgeWithin = "within"
	// EdgeDirectRoles holds the string denoting the direct_roles edge name in mutations.
	EdgeDirectRoles = "direct_roles"
	// Table holds the table name of the user in the database.
	Table = "users"
	// WithinTable is the table that holds the within relation/edge. The primary key declared below.
	WithinTable = "group_members"
	// WithinInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	WithinInverseTable = "groups"
	// DirectRolesTable is the table that holds the direct_roles relation/edge. The primary key declared below.
	DirectRolesTable = "role_users"
	// DirectRolesInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	DirectRolesInverseTable = "roles"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldBirthday,
}

var (
	// WithinPrimaryKey and WithinColumn2 are the table columns denoting the
	// primary key for the within relation (M2M).
	WithinPrimaryKey = []string{"group_id", "user_id"}
	// DirectRolesPrimaryKey and DirectRolesColumn2 are the table columns denoting the
	// primary key for the direct_roles relation (M2M).
	DirectRolesPrimaryKey = []string{"role_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByBirthday orders the results by the birthday field.
func ByBirthday(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthday, opts...).ToFunc()
}

// ByWithinCount orders the results by within count.
func ByWithinCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWithinStep(), opts...)
	}
}

// ByWithin orders the results by within terms.
func ByWithin(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWithinStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDirectRolesCount orders the results by direct_roles count.
func ByDirectRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDirectRolesStep(), opts...)
	}
}

// ByDirectRoles orders the results by direct_roles terms.
func ByDirectRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDirectRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newWithinStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WithinInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, WithinTable, WithinPrimaryKey...),
	)
}
func newDirectRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DirectRolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, DirectRolesTable, DirectRolesPrimaryKey...),
	)
}
