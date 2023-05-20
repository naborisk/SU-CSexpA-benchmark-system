// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/submit"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/taskresult"
)

// TaskResult is the model entity for the TaskResult schema.
type TaskResult struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RequestPerSec holds the value of the "request_per_sec" field.
	RequestPerSec int `json:"request_per_sec,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// ErrorMessage holds the value of the "error_message" field.
	ErrorMessage string `json:"error_message,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Method holds the value of the "method" field.
	Method string `json:"method,omitempty"`
	// RequestContentType holds the value of the "request_content_type" field.
	RequestContentType string `json:"request_content_type,omitempty"`
	// RequestBody holds the value of the "request_body" field.
	RequestBody string `json:"request_body,omitempty"`
	// ThreadNum holds the value of the "thread_num" field.
	ThreadNum int `json:"thread_num,omitempty"`
	// AttemptCount holds the value of the "attempt_count" field.
	AttemptCount int `json:"attempt_count,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskResultQuery when eager-loading is set.
	Edges               TaskResultEdges `json:"edges"`
	submit_task_results *int
	selectValues        sql.SelectValues
}

// TaskResultEdges holds the relations/edges for other nodes in the graph.
type TaskResultEdges struct {
	// Submits holds the value of the submits edge.
	Submits *Submit `json:"submits,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SubmitsOrErr returns the Submits value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskResultEdges) SubmitsOrErr() (*Submit, error) {
	if e.loadedTypes[0] {
		if e.Submits == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: submit.Label}
		}
		return e.Submits, nil
	}
	return nil, &NotLoadedError{edge: "submits"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TaskResult) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case taskresult.FieldID, taskresult.FieldRequestPerSec, taskresult.FieldThreadNum, taskresult.FieldAttemptCount:
			values[i] = new(sql.NullInt64)
		case taskresult.FieldStatus, taskresult.FieldErrorMessage, taskresult.FieldURL, taskresult.FieldMethod, taskresult.FieldRequestContentType, taskresult.FieldRequestBody:
			values[i] = new(sql.NullString)
		case taskresult.FieldCreatedAt, taskresult.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case taskresult.ForeignKeys[0]: // submit_task_results
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TaskResult fields.
func (tr *TaskResult) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case taskresult.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tr.ID = int(value.Int64)
		case taskresult.FieldRequestPerSec:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field request_per_sec", values[i])
			} else if value.Valid {
				tr.RequestPerSec = int(value.Int64)
			}
		case taskresult.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				tr.Status = value.String
			}
		case taskresult.FieldErrorMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field error_message", values[i])
			} else if value.Valid {
				tr.ErrorMessage = value.String
			}
		case taskresult.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				tr.URL = value.String
			}
		case taskresult.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				tr.Method = value.String
			}
		case taskresult.FieldRequestContentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field request_content_type", values[i])
			} else if value.Valid {
				tr.RequestContentType = value.String
			}
		case taskresult.FieldRequestBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field request_body", values[i])
			} else if value.Valid {
				tr.RequestBody = value.String
			}
		case taskresult.FieldThreadNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field thread_num", values[i])
			} else if value.Valid {
				tr.ThreadNum = int(value.Int64)
			}
		case taskresult.FieldAttemptCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field attempt_count", values[i])
			} else if value.Valid {
				tr.AttemptCount = int(value.Int64)
			}
		case taskresult.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tr.CreatedAt = value.Time
			}
		case taskresult.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				tr.DeletedAt = value.Time
			}
		case taskresult.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field submit_task_results", value)
			} else if value.Valid {
				tr.submit_task_results = new(int)
				*tr.submit_task_results = int(value.Int64)
			}
		default:
			tr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TaskResult.
// This includes values selected through modifiers, order, etc.
func (tr *TaskResult) Value(name string) (ent.Value, error) {
	return tr.selectValues.Get(name)
}

// QuerySubmits queries the "submits" edge of the TaskResult entity.
func (tr *TaskResult) QuerySubmits() *SubmitQuery {
	return NewTaskResultClient(tr.config).QuerySubmits(tr)
}

// Update returns a builder for updating this TaskResult.
// Note that you need to call TaskResult.Unwrap() before calling this method if this TaskResult
// was returned from a transaction, and the transaction was committed or rolled back.
func (tr *TaskResult) Update() *TaskResultUpdateOne {
	return NewTaskResultClient(tr.config).UpdateOne(tr)
}

// Unwrap unwraps the TaskResult entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tr *TaskResult) Unwrap() *TaskResult {
	_tx, ok := tr.config.driver.(*txDriver)
	if !ok {
		panic("ent: TaskResult is not a transactional entity")
	}
	tr.config.driver = _tx.drv
	return tr
}

// String implements the fmt.Stringer.
func (tr *TaskResult) String() string {
	var builder strings.Builder
	builder.WriteString("TaskResult(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tr.ID))
	builder.WriteString("request_per_sec=")
	builder.WriteString(fmt.Sprintf("%v", tr.RequestPerSec))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(tr.Status)
	builder.WriteString(", ")
	builder.WriteString("error_message=")
	builder.WriteString(tr.ErrorMessage)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(tr.URL)
	builder.WriteString(", ")
	builder.WriteString("method=")
	builder.WriteString(tr.Method)
	builder.WriteString(", ")
	builder.WriteString("request_content_type=")
	builder.WriteString(tr.RequestContentType)
	builder.WriteString(", ")
	builder.WriteString("request_body=")
	builder.WriteString(tr.RequestBody)
	builder.WriteString(", ")
	builder.WriteString("thread_num=")
	builder.WriteString(fmt.Sprintf("%v", tr.ThreadNum))
	builder.WriteString(", ")
	builder.WriteString("attempt_count=")
	builder.WriteString(fmt.Sprintf("%v", tr.AttemptCount))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(tr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(tr.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TaskResults is a parsable slice of TaskResult.
type TaskResults []*TaskResult
