// Code generated by ent, DO NOT EDIT.

package submit

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the submit type in the database.
	Label = "submit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIPAddr holds the string denoting the ip_addr field in the database.
	FieldIPAddr = "ip_addr"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldSubmitedAt holds the string denoting the submited_at field in the database.
	FieldSubmitedAt = "submited_at"
	// FieldCompletedAt holds the string denoting the completed_at field in the database.
	FieldCompletedAt = "completed_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeTagResults holds the string denoting the tagresults edge name in mutations.
	EdgeTagResults = "tagResults"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgeContests holds the string denoting the contests edge name in mutations.
	EdgeContests = "contests"
	// Table holds the table name of the submit in the database.
	Table = "submits"
	// TagResultsTable is the table that holds the tagResults relation/edge.
	TagResultsTable = "task_results"
	// TagResultsInverseTable is the table name for the TaskResult entity.
	// It exists in this package in order to avoid circular dependency with the "taskresult" package.
	TagResultsInverseTable = "task_results"
	// TagResultsColumn is the table column denoting the tagResults relation/edge.
	TagResultsColumn = "submit_tag_results"
	// GroupsTable is the table that holds the groups relation/edge. The primary key declared below.
	GroupsTable = "group_submits"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// ContestsTable is the table that holds the contests relation/edge. The primary key declared below.
	ContestsTable = "contest_submits"
	// ContestsInverseTable is the table name for the Contest entity.
	// It exists in this package in order to avoid circular dependency with the "contest" package.
	ContestsInverseTable = "contests"
)

// Columns holds all SQL columns for submit fields.
var Columns = []string{
	FieldID,
	FieldIPAddr,
	FieldYear,
	FieldScore,
	FieldLanguage,
	FieldSubmitedAt,
	FieldCompletedAt,
	FieldUpdatedAt,
}

var (
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"group_id", "submit_id"}
	// ContestsPrimaryKey and ContestsColumn2 are the table columns denoting the
	// primary key for the contests relation (M2M).
	ContestsPrimaryKey = []string{"contest_id", "submit_id"}
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
	// YearValidator is a validator for the "year" field. It is called by the builders before save.
	YearValidator func(int) error
	// ScoreValidator is a validator for the "score" field. It is called by the builders before save.
	ScoreValidator func(int) error
)

// Language defines the type for the "language" enum field.
type Language string

// Language values.
const (
	LanguagePhp        Language = "php"
	LanguageGo         Language = "go"
	LanguageRust       Language = "rust"
	LanguageJavascript Language = "javascript"
	LanguageCsharp     Language = "csharp"
	LanguageCpp        Language = "cpp"
	LanguageRuby       Language = "ruby"
	LanguagePython     Language = "python"
)

func (l Language) String() string {
	return string(l)
}

// LanguageValidator is a validator for the "language" field enum values. It is called by the builders before save.
func LanguageValidator(l Language) error {
	switch l {
	case LanguagePhp, LanguageGo, LanguageRust, LanguageJavascript, LanguageCsharp, LanguageCpp, LanguageRuby, LanguagePython:
		return nil
	default:
		return fmt.Errorf("submit: invalid enum value for language field: %q", l)
	}
}

// OrderOption defines the ordering options for the Submit queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIPAddr orders the results by the ip_addr field.
func ByIPAddr(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIPAddr, opts...).ToFunc()
}

// ByYear orders the results by the year field.
func ByYear(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldYear, opts...).ToFunc()
}

// ByScore orders the results by the score field.
func ByScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScore, opts...).ToFunc()
}

// ByLanguage orders the results by the language field.
func ByLanguage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLanguage, opts...).ToFunc()
}

// BySubmitedAt orders the results by the submited_at field.
func BySubmitedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubmitedAt, opts...).ToFunc()
}

// ByCompletedAt orders the results by the completed_at field.
func ByCompletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompletedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByTagResultsCount orders the results by tagResults count.
func ByTagResultsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagResultsStep(), opts...)
	}
}

// ByTagResults orders the results by tagResults terms.
func ByTagResults(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagResultsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByGroupsCount orders the results by groups count.
func ByGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGroupsStep(), opts...)
	}
}

// ByGroups orders the results by groups terms.
func ByGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByContestsCount orders the results by contests count.
func ByContestsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newContestsStep(), opts...)
	}
}

// ByContests orders the results by contests terms.
func ByContests(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContestsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTagResultsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagResultsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TagResultsTable, TagResultsColumn),
	)
}
func newGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, GroupsTable, GroupsPrimaryKey...),
	)
}
func newContestsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContestsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ContestsTable, ContestsPrimaryKey...),
	)
}
