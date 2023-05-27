// Code generated by ent, DO NOT EDIT.

package contest

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the contest type in the database.
	Label = "contest"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldStartAt holds the string denoting the start_at field in the database.
	FieldStartAt = "start_at"
	// FieldEndAt holds the string denoting the end_at field in the database.
	FieldEndAt = "end_at"
	// FieldSubmitLimit holds the string denoting the submit_limit field in the database.
	FieldSubmitLimit = "submit_limit"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// FieldTagSelection holds the string denoting the tag_selection field in the database.
	FieldTagSelection = "tag_selection"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeSubmits holds the string denoting the submits edge name in mutations.
	EdgeSubmits = "submits"
	// Table holds the table name of the contest in the database.
	Table = "contests"
	// SubmitsTable is the table that holds the submits relation/edge.
	SubmitsTable = "submits"
	// SubmitsInverseTable is the table name for the Submit entity.
	// It exists in this package in order to avoid circular dependency with the "submit" package.
	SubmitsInverseTable = "submits"
	// SubmitsColumn is the table column denoting the submits relation/edge.
	SubmitsColumn = "contest_submits"
)

// Columns holds all SQL columns for contest fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldStartAt,
	FieldEndAt,
	FieldSubmitLimit,
	FieldYear,
	FieldTagSelection,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// YearValidator is a validator for the "year" field. It is called by the builders before save.
	YearValidator func(int) error
)

// TagSelection defines the type for the "tag_selection" enum field.
type TagSelection string

// TagSelection values.
const (
	TagSelectionAuto   TagSelection = "auto"
	TagSelectionManual TagSelection = "manual"
)

func (ts TagSelection) String() string {
	return string(ts)
}

// TagSelectionValidator is a validator for the "tag_selection" field enum values. It is called by the builders before save.
func TagSelectionValidator(ts TagSelection) error {
	switch ts {
	case TagSelectionAuto, TagSelectionManual:
		return nil
	default:
		return fmt.Errorf("contest: invalid enum value for tag_selection field: %q", ts)
	}
}

// OrderOption defines the ordering options for the Contest queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByStartAt orders the results by the start_at field.
func ByStartAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartAt, opts...).ToFunc()
}

// ByEndAt orders the results by the end_at field.
func ByEndAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndAt, opts...).ToFunc()
}

// BySubmitLimit orders the results by the submit_limit field.
func BySubmitLimit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubmitLimit, opts...).ToFunc()
}

// ByYear orders the results by the year field.
func ByYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldYear, opts...).ToFunc()
}

// ByTagSelection orders the results by the tag_selection field.
func ByTagSelection(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTagSelection, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// BySubmitsCount orders the results by submits count.
func BySubmitsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubmitsStep(), opts...)
	}
}

// BySubmits orders the results by submits terms.
func BySubmits(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubmitsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSubmitsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubmitsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubmitsTable, SubmitsColumn),
	)
}
