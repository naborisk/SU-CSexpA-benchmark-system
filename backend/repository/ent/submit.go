// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/repository/ent/submit"
)

// Submit is the model entity for the Submit schema.
type Submit struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Year holds the value of the "year" field.
	Year int `json:"year,omitempty"`
	// Score holds the value of the "score" field.
	Score int `json:"score,omitempty"`
	// Language holds the value of the "language" field.
	Language submit.Language `json:"language,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// TaskNum holds the value of the "task_num" field.
	TaskNum int `json:"task_num,omitempty"`
	// SubmitedAt holds the value of the "submited_at" field.
	SubmitedAt time.Time `json:"submited_at,omitempty"`
	// CompletedAt holds the value of the "completed_at" field.
	CompletedAt time.Time `json:"completed_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubmitQuery when eager-loading is set.
	Edges           SubmitEdges `json:"edges"`
	contest_submits *int
	group_submits   *int
	selectValues    sql.SelectValues
}

// SubmitEdges holds the relations/edges for other nodes in the graph.
type SubmitEdges struct {
	// TaskResults holds the value of the taskResults edge.
	TaskResults []*TaskResult `json:"taskResults,omitempty"`
	// Groups holds the value of the groups edge.
	Groups *Group `json:"groups,omitempty"`
	// Contests holds the value of the contests edge.
	Contests *Contest `json:"contests,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// TaskResultsOrErr returns the TaskResults value or an error if the edge
// was not loaded in eager-loading.
func (e SubmitEdges) TaskResultsOrErr() ([]*TaskResult, error) {
	if e.loadedTypes[0] {
		return e.TaskResults, nil
	}
	return nil, &NotLoadedError{edge: "taskResults"}
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubmitEdges) GroupsOrErr() (*Group, error) {
	if e.loadedTypes[1] {
		if e.Groups == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// ContestsOrErr returns the Contests value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubmitEdges) ContestsOrErr() (*Contest, error) {
	if e.loadedTypes[2] {
		if e.Contests == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: contest.Label}
		}
		return e.Contests, nil
	}
	return nil, &NotLoadedError{edge: "contests"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Submit) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case submit.FieldID, submit.FieldYear, submit.FieldScore, submit.FieldTaskNum:
			values[i] = new(sql.NullInt64)
		case submit.FieldURL, submit.FieldLanguage, submit.FieldMessage, submit.FieldStatus:
			values[i] = new(sql.NullString)
		case submit.FieldSubmitedAt, submit.FieldCompletedAt, submit.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case submit.ForeignKeys[0]: // contest_submits
			values[i] = new(sql.NullInt64)
		case submit.ForeignKeys[1]: // group_submits
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Submit fields.
func (s *Submit) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case submit.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case submit.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				s.URL = value.String
			}
		case submit.FieldYear:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field year", values[i])
			} else if value.Valid {
				s.Year = int(value.Int64)
			}
		case submit.FieldScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				s.Score = int(value.Int64)
			}
		case submit.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				s.Language = submit.Language(value.String)
			}
		case submit.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				s.Message = value.String
			}
		case submit.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = value.String
			}
		case submit.FieldTaskNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field task_num", values[i])
			} else if value.Valid {
				s.TaskNum = int(value.Int64)
			}
		case submit.FieldSubmitedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field submited_at", values[i])
			} else if value.Valid {
				s.SubmitedAt = value.Time
			}
		case submit.FieldCompletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field completed_at", values[i])
			} else if value.Valid {
				s.CompletedAt = value.Time
			}
		case submit.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case submit.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field contest_submits", value)
			} else if value.Valid {
				s.contest_submits = new(int)
				*s.contest_submits = int(value.Int64)
			}
		case submit.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field group_submits", value)
			} else if value.Valid {
				s.group_submits = new(int)
				*s.group_submits = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Submit.
// This includes values selected through modifiers, order, etc.
func (s *Submit) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryTaskResults queries the "taskResults" edge of the Submit entity.
func (s *Submit) QueryTaskResults() *TaskResultQuery {
	return NewSubmitClient(s.config).QueryTaskResults(s)
}

// QueryGroups queries the "groups" edge of the Submit entity.
func (s *Submit) QueryGroups() *GroupQuery {
	return NewSubmitClient(s.config).QueryGroups(s)
}

// QueryContests queries the "contests" edge of the Submit entity.
func (s *Submit) QueryContests() *ContestQuery {
	return NewSubmitClient(s.config).QueryContests(s)
}

// Update returns a builder for updating this Submit.
// Note that you need to call Submit.Unwrap() before calling this method if this Submit
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Submit) Update() *SubmitUpdateOne {
	return NewSubmitClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Submit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Submit) Unwrap() *Submit {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Submit is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Submit) String() string {
	var builder strings.Builder
	builder.WriteString("Submit(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("url=")
	builder.WriteString(s.URL)
	builder.WriteString(", ")
	builder.WriteString("year=")
	builder.WriteString(fmt.Sprintf("%v", s.Year))
	builder.WriteString(", ")
	builder.WriteString("score=")
	builder.WriteString(fmt.Sprintf("%v", s.Score))
	builder.WriteString(", ")
	builder.WriteString("language=")
	builder.WriteString(fmt.Sprintf("%v", s.Language))
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(s.Message)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(s.Status)
	builder.WriteString(", ")
	builder.WriteString("task_num=")
	builder.WriteString(fmt.Sprintf("%v", s.TaskNum))
	builder.WriteString(", ")
	builder.WriteString("submited_at=")
	builder.WriteString(s.SubmitedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("completed_at=")
	builder.WriteString(s.CompletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Submits is a parsable slice of Submit.
type Submits []*Submit
