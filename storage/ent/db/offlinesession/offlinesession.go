// Code generated by ent, DO NOT EDIT.

package offlinesession

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the offlinesession type in the database.
	Label = "offline_session"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldConnID holds the string denoting the conn_id field in the database.
	FieldConnID = "conn_id"
	// FieldRefresh holds the string denoting the refresh field in the database.
	FieldRefresh = "refresh"
	// FieldConnectorData holds the string denoting the connector_data field in the database.
	FieldConnectorData = "connector_data"
	// Table holds the table name of the offlinesession in the database.
	Table = "offline_sessions"
)

// Columns holds all SQL columns for offlinesession fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldConnID,
	FieldRefresh,
	FieldConnectorData,
}

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
	// UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	UserIDValidator func(string) error
	// ConnIDValidator is a validator for the "conn_id" field. It is called by the builders before save.
	ConnIDValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the OfflineSession queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByConnID orders the results by the conn_id field.
func ByConnID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConnID, opts...).ToFunc()
}
