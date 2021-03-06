// Code generated by entc, DO NOT EDIT.

package group

import (
	"aio/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	uuid "github.com/satori/go.uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v uuid.UUID) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// ShareURL applies equality check predicate on the "share_url" field. It's identical to ShareURLEQ.
func ShareURL(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShareURL), v))
	})
}

// Option applies equality check predicate on the "option" field. It's identical to OptionEQ.
func Option(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOption), v))
	})
}

// Seq applies equality check predicate on the "seq" field. It's identical to SeqEQ.
func Seq(v int32) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeq), v))
	})
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v uuid.UUID) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v uuid.UUID) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUID), v))
	})
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...uuid.UUID) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUID), v...))
	})
}

// UIDNotIn applies the NotIn predicate on the "uid" field.
func UIDNotIn(vs ...uuid.UUID) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUID), v...))
	})
}

// UIDGT applies the GT predicate on the "uid" field.
func UIDGT(v uuid.UUID) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUID), v))
	})
}

// UIDGTE applies the GTE predicate on the "uid" field.
func UIDGTE(v uuid.UUID) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUID), v))
	})
}

// UIDLT applies the LT predicate on the "uid" field.
func UIDLT(v uuid.UUID) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUID), v))
	})
}

// UIDLTE applies the LTE predicate on the "uid" field.
func UIDLTE(v uuid.UUID) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUID), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// ShareURLEQ applies the EQ predicate on the "share_url" field.
func ShareURLEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShareURL), v))
	})
}

// ShareURLNEQ applies the NEQ predicate on the "share_url" field.
func ShareURLNEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldShareURL), v))
	})
}

// ShareURLIn applies the In predicate on the "share_url" field.
func ShareURLIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldShareURL), v...))
	})
}

// ShareURLNotIn applies the NotIn predicate on the "share_url" field.
func ShareURLNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldShareURL), v...))
	})
}

// ShareURLGT applies the GT predicate on the "share_url" field.
func ShareURLGT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldShareURL), v))
	})
}

// ShareURLGTE applies the GTE predicate on the "share_url" field.
func ShareURLGTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldShareURL), v))
	})
}

// ShareURLLT applies the LT predicate on the "share_url" field.
func ShareURLLT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldShareURL), v))
	})
}

// ShareURLLTE applies the LTE predicate on the "share_url" field.
func ShareURLLTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldShareURL), v))
	})
}

// ShareURLContains applies the Contains predicate on the "share_url" field.
func ShareURLContains(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldShareURL), v))
	})
}

// ShareURLHasPrefix applies the HasPrefix predicate on the "share_url" field.
func ShareURLHasPrefix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldShareURL), v))
	})
}

// ShareURLHasSuffix applies the HasSuffix predicate on the "share_url" field.
func ShareURLHasSuffix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldShareURL), v))
	})
}

// ShareURLIsNil applies the IsNil predicate on the "share_url" field.
func ShareURLIsNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldShareURL)))
	})
}

// ShareURLNotNil applies the NotNil predicate on the "share_url" field.
func ShareURLNotNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldShareURL)))
	})
}

// ShareURLEqualFold applies the EqualFold predicate on the "share_url" field.
func ShareURLEqualFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldShareURL), v))
	})
}

// ShareURLContainsFold applies the ContainsFold predicate on the "share_url" field.
func ShareURLContainsFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldShareURL), v))
	})
}

// OptionEQ applies the EQ predicate on the "option" field.
func OptionEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOption), v))
	})
}

// OptionNEQ applies the NEQ predicate on the "option" field.
func OptionNEQ(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOption), v))
	})
}

// OptionIn applies the In predicate on the "option" field.
func OptionIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOption), v...))
	})
}

// OptionNotIn applies the NotIn predicate on the "option" field.
func OptionNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOption), v...))
	})
}

// OptionGT applies the GT predicate on the "option" field.
func OptionGT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOption), v))
	})
}

// OptionGTE applies the GTE predicate on the "option" field.
func OptionGTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOption), v))
	})
}

// OptionLT applies the LT predicate on the "option" field.
func OptionLT(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOption), v))
	})
}

// OptionLTE applies the LTE predicate on the "option" field.
func OptionLTE(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOption), v))
	})
}

// OptionContains applies the Contains predicate on the "option" field.
func OptionContains(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOption), v))
	})
}

// OptionHasPrefix applies the HasPrefix predicate on the "option" field.
func OptionHasPrefix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOption), v))
	})
}

// OptionHasSuffix applies the HasSuffix predicate on the "option" field.
func OptionHasSuffix(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOption), v))
	})
}

// OptionIsNil applies the IsNil predicate on the "option" field.
func OptionIsNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOption)))
	})
}

// OptionNotNil applies the NotNil predicate on the "option" field.
func OptionNotNil() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOption)))
	})
}

// OptionEqualFold applies the EqualFold predicate on the "option" field.
func OptionEqualFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOption), v))
	})
}

// OptionContainsFold applies the ContainsFold predicate on the "option" field.
func OptionContainsFold(v string) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOption), v))
	})
}

// SeqEQ applies the EQ predicate on the "seq" field.
func SeqEQ(v int32) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSeq), v))
	})
}

// SeqNEQ applies the NEQ predicate on the "seq" field.
func SeqNEQ(v int32) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSeq), v))
	})
}

// SeqIn applies the In predicate on the "seq" field.
func SeqIn(vs ...int32) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSeq), v...))
	})
}

// SeqNotIn applies the NotIn predicate on the "seq" field.
func SeqNotIn(vs ...int32) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Group(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSeq), v...))
	})
}

// SeqGT applies the GT predicate on the "seq" field.
func SeqGT(v int32) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSeq), v))
	})
}

// SeqGTE applies the GTE predicate on the "seq" field.
func SeqGTE(v int32) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSeq), v))
	})
}

// SeqLT applies the LT predicate on the "seq" field.
func SeqLT(v int32) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSeq), v))
	})
}

// SeqLTE applies the LTE predicate on the "seq" field.
func SeqLTE(v int32) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSeq), v))
	})
}

// HasTabs applies the HasEdge predicate on the "tabs" edge.
func HasTabs() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TabsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TabsTable, TabsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTabsWith applies the HasEdge predicate on the "tabs" edge with a given conditions (other predicates).
func HasTabsWith(preds ...predicate.Tab) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TabsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TabsTable, TabsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClient applies the HasEdge predicate on the "client" edge.
func HasClient() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClientTable, ClientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClientWith applies the HasEdge predicate on the "client" edge with a given conditions (other predicates).
func HasClientWith(preds ...predicate.ExtensionClient) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClientTable, ClientColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
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
func Not(p predicate.Group) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		p(s.Not())
	})
}
