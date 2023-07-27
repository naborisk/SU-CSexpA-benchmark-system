// Code generated by ent, DO NOT EDIT.

package contest

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldTitle, v))
}

// StartAt applies equality check predicate on the "start_at" field. It's identical to StartAtEQ.
func StartAt(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldStartAt, v))
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldEndAt, v))
}

// SubmitLimit applies equality check predicate on the "submit_limit" field. It's identical to SubmitLimitEQ.
func SubmitLimit(v int) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldSubmitLimit, v))
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldSlug, v))
}

// Validator applies equality check predicate on the "validator" field. It's identical to ValidatorEQ.
func Validator(v string) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldValidator, v))
}

// TimeLimitPerTask applies equality check predicate on the "time_limit_per_task" field. It's identical to TimeLimitPerTaskEQ.
func TimeLimitPerTask(v int64) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldTimeLimitPerTask, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldUpdatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Contest {
	return predicate.Contest(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Contest {
	return predicate.Contest(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Contest {
	return predicate.Contest(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Contest {
	return predicate.Contest(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Contest {
	return predicate.Contest(sql.FieldContainsFold(FieldTitle, v))
}

// StartAtEQ applies the EQ predicate on the "start_at" field.
func StartAtEQ(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldStartAt, v))
}

// StartAtNEQ applies the NEQ predicate on the "start_at" field.
func StartAtNEQ(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldStartAt, v))
}

// StartAtIn applies the In predicate on the "start_at" field.
func StartAtIn(vs ...time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldStartAt, vs...))
}

// StartAtNotIn applies the NotIn predicate on the "start_at" field.
func StartAtNotIn(vs ...time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldStartAt, vs...))
}

// StartAtGT applies the GT predicate on the "start_at" field.
func StartAtGT(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldStartAt, v))
}

// StartAtGTE applies the GTE predicate on the "start_at" field.
func StartAtGTE(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldStartAt, v))
}

// StartAtLT applies the LT predicate on the "start_at" field.
func StartAtLT(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldStartAt, v))
}

// StartAtLTE applies the LTE predicate on the "start_at" field.
func StartAtLTE(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldStartAt, v))
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldEndAt, v))
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldEndAt, v))
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldEndAt, vs...))
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldEndAt, vs...))
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldEndAt, v))
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldEndAt, v))
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldEndAt, v))
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldEndAt, v))
}

// SubmitLimitEQ applies the EQ predicate on the "submit_limit" field.
func SubmitLimitEQ(v int) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldSubmitLimit, v))
}

// SubmitLimitNEQ applies the NEQ predicate on the "submit_limit" field.
func SubmitLimitNEQ(v int) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldSubmitLimit, v))
}

// SubmitLimitIn applies the In predicate on the "submit_limit" field.
func SubmitLimitIn(vs ...int) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldSubmitLimit, vs...))
}

// SubmitLimitNotIn applies the NotIn predicate on the "submit_limit" field.
func SubmitLimitNotIn(vs ...int) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldSubmitLimit, vs...))
}

// SubmitLimitGT applies the GT predicate on the "submit_limit" field.
func SubmitLimitGT(v int) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldSubmitLimit, v))
}

// SubmitLimitGTE applies the GTE predicate on the "submit_limit" field.
func SubmitLimitGTE(v int) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldSubmitLimit, v))
}

// SubmitLimitLT applies the LT predicate on the "submit_limit" field.
func SubmitLimitLT(v int) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldSubmitLimit, v))
}

// SubmitLimitLTE applies the LTE predicate on the "submit_limit" field.
func SubmitLimitLTE(v int) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldSubmitLimit, v))
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldSlug, v))
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldSlug, v))
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldSlug, vs...))
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldSlug, vs...))
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldSlug, v))
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldSlug, v))
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldSlug, v))
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldSlug, v))
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.Contest {
	return predicate.Contest(sql.FieldContains(FieldSlug, v))
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.Contest {
	return predicate.Contest(sql.FieldHasPrefix(FieldSlug, v))
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.Contest {
	return predicate.Contest(sql.FieldHasSuffix(FieldSlug, v))
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.Contest {
	return predicate.Contest(sql.FieldEqualFold(FieldSlug, v))
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.Contest {
	return predicate.Contest(sql.FieldContainsFold(FieldSlug, v))
}

// TagSelectionLogicEQ applies the EQ predicate on the "tag_selection_logic" field.
func TagSelectionLogicEQ(v TagSelectionLogic) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldTagSelectionLogic, v))
}

// TagSelectionLogicNEQ applies the NEQ predicate on the "tag_selection_logic" field.
func TagSelectionLogicNEQ(v TagSelectionLogic) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldTagSelectionLogic, v))
}

// TagSelectionLogicIn applies the In predicate on the "tag_selection_logic" field.
func TagSelectionLogicIn(vs ...TagSelectionLogic) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldTagSelectionLogic, vs...))
}

// TagSelectionLogicNotIn applies the NotIn predicate on the "tag_selection_logic" field.
func TagSelectionLogicNotIn(vs ...TagSelectionLogic) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldTagSelectionLogic, vs...))
}

// ValidatorEQ applies the EQ predicate on the "validator" field.
func ValidatorEQ(v string) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldValidator, v))
}

// ValidatorNEQ applies the NEQ predicate on the "validator" field.
func ValidatorNEQ(v string) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldValidator, v))
}

// ValidatorIn applies the In predicate on the "validator" field.
func ValidatorIn(vs ...string) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldValidator, vs...))
}

// ValidatorNotIn applies the NotIn predicate on the "validator" field.
func ValidatorNotIn(vs ...string) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldValidator, vs...))
}

// ValidatorGT applies the GT predicate on the "validator" field.
func ValidatorGT(v string) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldValidator, v))
}

// ValidatorGTE applies the GTE predicate on the "validator" field.
func ValidatorGTE(v string) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldValidator, v))
}

// ValidatorLT applies the LT predicate on the "validator" field.
func ValidatorLT(v string) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldValidator, v))
}

// ValidatorLTE applies the LTE predicate on the "validator" field.
func ValidatorLTE(v string) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldValidator, v))
}

// ValidatorContains applies the Contains predicate on the "validator" field.
func ValidatorContains(v string) predicate.Contest {
	return predicate.Contest(sql.FieldContains(FieldValidator, v))
}

// ValidatorHasPrefix applies the HasPrefix predicate on the "validator" field.
func ValidatorHasPrefix(v string) predicate.Contest {
	return predicate.Contest(sql.FieldHasPrefix(FieldValidator, v))
}

// ValidatorHasSuffix applies the HasSuffix predicate on the "validator" field.
func ValidatorHasSuffix(v string) predicate.Contest {
	return predicate.Contest(sql.FieldHasSuffix(FieldValidator, v))
}

// ValidatorEqualFold applies the EqualFold predicate on the "validator" field.
func ValidatorEqualFold(v string) predicate.Contest {
	return predicate.Contest(sql.FieldEqualFold(FieldValidator, v))
}

// ValidatorContainsFold applies the ContainsFold predicate on the "validator" field.
func ValidatorContainsFold(v string) predicate.Contest {
	return predicate.Contest(sql.FieldContainsFold(FieldValidator, v))
}

// TimeLimitPerTaskEQ applies the EQ predicate on the "time_limit_per_task" field.
func TimeLimitPerTaskEQ(v int64) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldTimeLimitPerTask, v))
}

// TimeLimitPerTaskNEQ applies the NEQ predicate on the "time_limit_per_task" field.
func TimeLimitPerTaskNEQ(v int64) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldTimeLimitPerTask, v))
}

// TimeLimitPerTaskIn applies the In predicate on the "time_limit_per_task" field.
func TimeLimitPerTaskIn(vs ...int64) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldTimeLimitPerTask, vs...))
}

// TimeLimitPerTaskNotIn applies the NotIn predicate on the "time_limit_per_task" field.
func TimeLimitPerTaskNotIn(vs ...int64) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldTimeLimitPerTask, vs...))
}

// TimeLimitPerTaskGT applies the GT predicate on the "time_limit_per_task" field.
func TimeLimitPerTaskGT(v int64) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldTimeLimitPerTask, v))
}

// TimeLimitPerTaskGTE applies the GTE predicate on the "time_limit_per_task" field.
func TimeLimitPerTaskGTE(v int64) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldTimeLimitPerTask, v))
}

// TimeLimitPerTaskLT applies the LT predicate on the "time_limit_per_task" field.
func TimeLimitPerTaskLT(v int64) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldTimeLimitPerTask, v))
}

// TimeLimitPerTaskLTE applies the LTE predicate on the "time_limit_per_task" field.
func TimeLimitPerTaskLTE(v int64) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldTimeLimitPerTask, v))
}

// TimeLimitPerTaskIsNil applies the IsNil predicate on the "time_limit_per_task" field.
func TimeLimitPerTaskIsNil() predicate.Contest {
	return predicate.Contest(sql.FieldIsNull(FieldTimeLimitPerTask))
}

// TimeLimitPerTaskNotNil applies the NotNil predicate on the "time_limit_per_task" field.
func TimeLimitPerTaskNotNil() predicate.Contest {
	return predicate.Contest(sql.FieldNotNull(FieldTimeLimitPerTask))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Contest {
	return predicate.Contest(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Contest {
	return predicate.Contest(sql.FieldNotNull(FieldUpdatedAt))
}

// HasSubmits applies the HasEdge predicate on the "submits" edge.
func HasSubmits() predicate.Contest {
	return predicate.Contest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubmitsTable, SubmitsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubmitsWith applies the HasEdge predicate on the "submits" edge with a given conditions (other predicates).
func HasSubmitsWith(preds ...predicate.Submit) predicate.Contest {
	return predicate.Contest(func(s *sql.Selector) {
		step := newSubmitsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Contest) predicate.Contest {
	return predicate.Contest(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Contest) predicate.Contest {
	return predicate.Contest(func(s *sql.Selector) {
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
func Not(p predicate.Contest) predicate.Contest {
	return predicate.Contest(func(s *sql.Selector) {
		p(s.Not())
	})
}
