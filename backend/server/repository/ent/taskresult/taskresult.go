// Code generated by ent, DO NOT EDIT.

package taskresult

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the taskresult type in the database.
	Label = "task_result"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRequestPerSec holds the string denoting the request_per_sec field in the database.
	FieldRequestPerSec = "request_per_sec"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldErrorMessage holds the string denoting the error_message field in the database.
	FieldErrorMessage = "error_message"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldMethod holds the string denoting the method field in the database.
	FieldMethod = "method"
	// FieldRequestContentType holds the string denoting the request_content_type field in the database.
	FieldRequestContentType = "request_content_type"
	// FieldRequestBody holds the string denoting the request_body field in the database.
	FieldRequestBody = "request_body"
	// FieldThreadNum holds the string denoting the thread_num field in the database.
	FieldThreadNum = "thread_num"
	// FieldAttemptCount holds the string denoting the attempt_count field in the database.
	FieldAttemptCount = "attempt_count"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeSubmits holds the string denoting the submits edge name in mutations.
	EdgeSubmits = "submits"
	// Table holds the table name of the taskresult in the database.
	Table = "task_results"
	// SubmitsTable is the table that holds the submits relation/edge.
	SubmitsTable = "task_results"
	// SubmitsInverseTable is the table name for the Submit entity.
	// It exists in this package in order to avoid circular dependency with the "submit" package.
	SubmitsInverseTable = "submits"
	// SubmitsColumn is the table column denoting the submits relation/edge.
	SubmitsColumn = "submit_task_results"
)

// Columns holds all SQL columns for taskresult fields.
var Columns = []string{
	FieldID,
	FieldRequestPerSec,
	FieldStatus,
	FieldErrorMessage,
	FieldURL,
	FieldMethod,
	FieldRequestContentType,
	FieldRequestBody,
	FieldThreadNum,
	FieldAttemptCount,
	FieldCreatedAt,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "task_results"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"submit_task_results",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the TaskResult queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRequestPerSec orders the results by the request_per_sec field.
func ByRequestPerSec(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequestPerSec, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByErrorMessage orders the results by the error_message field.
func ByErrorMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldErrorMessage, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByMethod orders the results by the method field.
func ByMethod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMethod, opts...).ToFunc()
}

// ByRequestContentType orders the results by the request_content_type field.
func ByRequestContentType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequestContentType, opts...).ToFunc()
}

// ByRequestBody orders the results by the request_body field.
func ByRequestBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequestBody, opts...).ToFunc()
}

// ByThreadNum orders the results by the thread_num field.
func ByThreadNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThreadNum, opts...).ToFunc()
}

// ByAttemptCount orders the results by the attempt_count field.
func ByAttemptCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAttemptCount, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// BySubmitsField orders the results by submits field.
func BySubmitsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubmitsStep(), sql.OrderByField(field, opts...))
	}
}
func newSubmitsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubmitsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SubmitsTable, SubmitsColumn),
	)
}
