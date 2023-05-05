// Code generated by ent, DO NOT EDIT.

package submit

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Submit {
	return predicate.Submit(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Submit {
	return predicate.Submit(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Submit {
	return predicate.Submit(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Submit {
	return predicate.Submit(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Submit {
	return predicate.Submit(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Submit {
	return predicate.Submit(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Submit {
	return predicate.Submit(sql.FieldLTE(FieldID, id))
}

// IPAddr applies equality check predicate on the "ip_addr" field. It's identical to IPAddrEQ.
func IPAddr(v string) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldIPAddr, v))
}

// Year applies equality check predicate on the "year" field. It's identical to YearEQ.
func Year(v int) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldYear, v))
}

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v int) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldScore, v))
}

// SubmitedAt applies equality check predicate on the "submited_at" field. It's identical to SubmitedAtEQ.
func SubmitedAt(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldSubmitedAt, v))
}

// CompletedAt applies equality check predicate on the "completed_at" field. It's identical to CompletedAtEQ.
func CompletedAt(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldCompletedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldUpdatedAt, v))
}

// IPAddrEQ applies the EQ predicate on the "ip_addr" field.
func IPAddrEQ(v string) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldIPAddr, v))
}

// IPAddrNEQ applies the NEQ predicate on the "ip_addr" field.
func IPAddrNEQ(v string) predicate.Submit {
	return predicate.Submit(sql.FieldNEQ(FieldIPAddr, v))
}

// IPAddrIn applies the In predicate on the "ip_addr" field.
func IPAddrIn(vs ...string) predicate.Submit {
	return predicate.Submit(sql.FieldIn(FieldIPAddr, vs...))
}

// IPAddrNotIn applies the NotIn predicate on the "ip_addr" field.
func IPAddrNotIn(vs ...string) predicate.Submit {
	return predicate.Submit(sql.FieldNotIn(FieldIPAddr, vs...))
}

// IPAddrGT applies the GT predicate on the "ip_addr" field.
func IPAddrGT(v string) predicate.Submit {
	return predicate.Submit(sql.FieldGT(FieldIPAddr, v))
}

// IPAddrGTE applies the GTE predicate on the "ip_addr" field.
func IPAddrGTE(v string) predicate.Submit {
	return predicate.Submit(sql.FieldGTE(FieldIPAddr, v))
}

// IPAddrLT applies the LT predicate on the "ip_addr" field.
func IPAddrLT(v string) predicate.Submit {
	return predicate.Submit(sql.FieldLT(FieldIPAddr, v))
}

// IPAddrLTE applies the LTE predicate on the "ip_addr" field.
func IPAddrLTE(v string) predicate.Submit {
	return predicate.Submit(sql.FieldLTE(FieldIPAddr, v))
}

// IPAddrContains applies the Contains predicate on the "ip_addr" field.
func IPAddrContains(v string) predicate.Submit {
	return predicate.Submit(sql.FieldContains(FieldIPAddr, v))
}

// IPAddrHasPrefix applies the HasPrefix predicate on the "ip_addr" field.
func IPAddrHasPrefix(v string) predicate.Submit {
	return predicate.Submit(sql.FieldHasPrefix(FieldIPAddr, v))
}

// IPAddrHasSuffix applies the HasSuffix predicate on the "ip_addr" field.
func IPAddrHasSuffix(v string) predicate.Submit {
	return predicate.Submit(sql.FieldHasSuffix(FieldIPAddr, v))
}

// IPAddrEqualFold applies the EqualFold predicate on the "ip_addr" field.
func IPAddrEqualFold(v string) predicate.Submit {
	return predicate.Submit(sql.FieldEqualFold(FieldIPAddr, v))
}

// IPAddrContainsFold applies the ContainsFold predicate on the "ip_addr" field.
func IPAddrContainsFold(v string) predicate.Submit {
	return predicate.Submit(sql.FieldContainsFold(FieldIPAddr, v))
}

// YearEQ applies the EQ predicate on the "year" field.
func YearEQ(v int) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldYear, v))
}

// YearNEQ applies the NEQ predicate on the "year" field.
func YearNEQ(v int) predicate.Submit {
	return predicate.Submit(sql.FieldNEQ(FieldYear, v))
}

// YearIn applies the In predicate on the "year" field.
func YearIn(vs ...int) predicate.Submit {
	return predicate.Submit(sql.FieldIn(FieldYear, vs...))
}

// YearNotIn applies the NotIn predicate on the "year" field.
func YearNotIn(vs ...int) predicate.Submit {
	return predicate.Submit(sql.FieldNotIn(FieldYear, vs...))
}

// YearGT applies the GT predicate on the "year" field.
func YearGT(v int) predicate.Submit {
	return predicate.Submit(sql.FieldGT(FieldYear, v))
}

// YearGTE applies the GTE predicate on the "year" field.
func YearGTE(v int) predicate.Submit {
	return predicate.Submit(sql.FieldGTE(FieldYear, v))
}

// YearLT applies the LT predicate on the "year" field.
func YearLT(v int) predicate.Submit {
	return predicate.Submit(sql.FieldLT(FieldYear, v))
}

// YearLTE applies the LTE predicate on the "year" field.
func YearLTE(v int) predicate.Submit {
	return predicate.Submit(sql.FieldLTE(FieldYear, v))
}

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v int) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldScore, v))
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v int) predicate.Submit {
	return predicate.Submit(sql.FieldNEQ(FieldScore, v))
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...int) predicate.Submit {
	return predicate.Submit(sql.FieldIn(FieldScore, vs...))
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...int) predicate.Submit {
	return predicate.Submit(sql.FieldNotIn(FieldScore, vs...))
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v int) predicate.Submit {
	return predicate.Submit(sql.FieldGT(FieldScore, v))
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v int) predicate.Submit {
	return predicate.Submit(sql.FieldGTE(FieldScore, v))
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v int) predicate.Submit {
	return predicate.Submit(sql.FieldLT(FieldScore, v))
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v int) predicate.Submit {
	return predicate.Submit(sql.FieldLTE(FieldScore, v))
}

// ScoreIsNil applies the IsNil predicate on the "score" field.
func ScoreIsNil() predicate.Submit {
	return predicate.Submit(sql.FieldIsNull(FieldScore))
}

// ScoreNotNil applies the NotNil predicate on the "score" field.
func ScoreNotNil() predicate.Submit {
	return predicate.Submit(sql.FieldNotNull(FieldScore))
}

// LanguageEQ applies the EQ predicate on the "language" field.
func LanguageEQ(v Language) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldLanguage, v))
}

// LanguageNEQ applies the NEQ predicate on the "language" field.
func LanguageNEQ(v Language) predicate.Submit {
	return predicate.Submit(sql.FieldNEQ(FieldLanguage, v))
}

// LanguageIn applies the In predicate on the "language" field.
func LanguageIn(vs ...Language) predicate.Submit {
	return predicate.Submit(sql.FieldIn(FieldLanguage, vs...))
}

// LanguageNotIn applies the NotIn predicate on the "language" field.
func LanguageNotIn(vs ...Language) predicate.Submit {
	return predicate.Submit(sql.FieldNotIn(FieldLanguage, vs...))
}

// LanguageIsNil applies the IsNil predicate on the "language" field.
func LanguageIsNil() predicate.Submit {
	return predicate.Submit(sql.FieldIsNull(FieldLanguage))
}

// LanguageNotNil applies the NotNil predicate on the "language" field.
func LanguageNotNil() predicate.Submit {
	return predicate.Submit(sql.FieldNotNull(FieldLanguage))
}

// SubmitedAtEQ applies the EQ predicate on the "submited_at" field.
func SubmitedAtEQ(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldSubmitedAt, v))
}

// SubmitedAtNEQ applies the NEQ predicate on the "submited_at" field.
func SubmitedAtNEQ(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldNEQ(FieldSubmitedAt, v))
}

// SubmitedAtIn applies the In predicate on the "submited_at" field.
func SubmitedAtIn(vs ...time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldIn(FieldSubmitedAt, vs...))
}

// SubmitedAtNotIn applies the NotIn predicate on the "submited_at" field.
func SubmitedAtNotIn(vs ...time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldNotIn(FieldSubmitedAt, vs...))
}

// SubmitedAtGT applies the GT predicate on the "submited_at" field.
func SubmitedAtGT(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldGT(FieldSubmitedAt, v))
}

// SubmitedAtGTE applies the GTE predicate on the "submited_at" field.
func SubmitedAtGTE(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldGTE(FieldSubmitedAt, v))
}

// SubmitedAtLT applies the LT predicate on the "submited_at" field.
func SubmitedAtLT(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldLT(FieldSubmitedAt, v))
}

// SubmitedAtLTE applies the LTE predicate on the "submited_at" field.
func SubmitedAtLTE(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldLTE(FieldSubmitedAt, v))
}

// CompletedAtEQ applies the EQ predicate on the "completed_at" field.
func CompletedAtEQ(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldCompletedAt, v))
}

// CompletedAtNEQ applies the NEQ predicate on the "completed_at" field.
func CompletedAtNEQ(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldNEQ(FieldCompletedAt, v))
}

// CompletedAtIn applies the In predicate on the "completed_at" field.
func CompletedAtIn(vs ...time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldIn(FieldCompletedAt, vs...))
}

// CompletedAtNotIn applies the NotIn predicate on the "completed_at" field.
func CompletedAtNotIn(vs ...time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldNotIn(FieldCompletedAt, vs...))
}

// CompletedAtGT applies the GT predicate on the "completed_at" field.
func CompletedAtGT(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldGT(FieldCompletedAt, v))
}

// CompletedAtGTE applies the GTE predicate on the "completed_at" field.
func CompletedAtGTE(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldGTE(FieldCompletedAt, v))
}

// CompletedAtLT applies the LT predicate on the "completed_at" field.
func CompletedAtLT(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldLT(FieldCompletedAt, v))
}

// CompletedAtLTE applies the LTE predicate on the "completed_at" field.
func CompletedAtLTE(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldLTE(FieldCompletedAt, v))
}

// CompletedAtIsNil applies the IsNil predicate on the "completed_at" field.
func CompletedAtIsNil() predicate.Submit {
	return predicate.Submit(sql.FieldIsNull(FieldCompletedAt))
}

// CompletedAtNotNil applies the NotNil predicate on the "completed_at" field.
func CompletedAtNotNil() predicate.Submit {
	return predicate.Submit(sql.FieldNotNull(FieldCompletedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Submit {
	return predicate.Submit(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Submit {
	return predicate.Submit(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Submit {
	return predicate.Submit(sql.FieldNotNull(FieldUpdatedAt))
}

// HasTaskResults applies the HasEdge predicate on the "taskResults" edge.
func HasTaskResults() predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TaskResultsTable, TaskResultsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskResultsWith applies the HasEdge predicate on the "taskResults" edge with a given conditions (other predicates).
func HasTaskResultsWith(preds ...predicate.TaskResult) predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		step := newTaskResultsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GroupsTable, GroupsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.Group) predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		step := newGroupsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasContests applies the HasEdge predicate on the "contests" edge.
func HasContests() predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ContestsTable, ContestsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContestsWith applies the HasEdge predicate on the "contests" edge with a given conditions (other predicates).
func HasContestsWith(preds ...predicate.Contest) predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		step := newContestsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Submit) predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Submit) predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
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
func Not(p predicate.Submit) predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		p(s.Not())
	})
}
