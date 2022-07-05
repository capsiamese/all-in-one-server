// Code generated by entc, DO NOT EDIT.

package extensionclient

import (
	"notification/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// ExtensionID applies equality check predicate on the "extension_id" field. It's identical to ExtensionIDEQ.
func ExtensionID(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExtensionID), v))
	})
}

// LastAccessTime applies equality check predicate on the "last_access_time" field. It's identical to LastAccessTimeEQ.
func LastAccessTime(v time.Time) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastAccessTime), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ExtensionClient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtensionClient(func(s *sql.Selector) {
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
func NameNotIn(vs ...string) predicate.ExtensionClient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtensionClient(func(s *sql.Selector) {
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
func NameGT(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ExtensionIDEQ applies the EQ predicate on the "extension_id" field.
func ExtensionIDEQ(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExtensionID), v))
	})
}

// ExtensionIDNEQ applies the NEQ predicate on the "extension_id" field.
func ExtensionIDNEQ(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExtensionID), v))
	})
}

// ExtensionIDIn applies the In predicate on the "extension_id" field.
func ExtensionIDIn(vs ...string) predicate.ExtensionClient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtensionClient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExtensionID), v...))
	})
}

// ExtensionIDNotIn applies the NotIn predicate on the "extension_id" field.
func ExtensionIDNotIn(vs ...string) predicate.ExtensionClient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtensionClient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExtensionID), v...))
	})
}

// ExtensionIDGT applies the GT predicate on the "extension_id" field.
func ExtensionIDGT(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExtensionID), v))
	})
}

// ExtensionIDGTE applies the GTE predicate on the "extension_id" field.
func ExtensionIDGTE(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExtensionID), v))
	})
}

// ExtensionIDLT applies the LT predicate on the "extension_id" field.
func ExtensionIDLT(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExtensionID), v))
	})
}

// ExtensionIDLTE applies the LTE predicate on the "extension_id" field.
func ExtensionIDLTE(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExtensionID), v))
	})
}

// ExtensionIDContains applies the Contains predicate on the "extension_id" field.
func ExtensionIDContains(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExtensionID), v))
	})
}

// ExtensionIDHasPrefix applies the HasPrefix predicate on the "extension_id" field.
func ExtensionIDHasPrefix(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExtensionID), v))
	})
}

// ExtensionIDHasSuffix applies the HasSuffix predicate on the "extension_id" field.
func ExtensionIDHasSuffix(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExtensionID), v))
	})
}

// ExtensionIDEqualFold applies the EqualFold predicate on the "extension_id" field.
func ExtensionIDEqualFold(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExtensionID), v))
	})
}

// ExtensionIDContainsFold applies the ContainsFold predicate on the "extension_id" field.
func ExtensionIDContainsFold(v string) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExtensionID), v))
	})
}

// LastAccessTimeEQ applies the EQ predicate on the "last_access_time" field.
func LastAccessTimeEQ(v time.Time) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastAccessTime), v))
	})
}

// LastAccessTimeNEQ applies the NEQ predicate on the "last_access_time" field.
func LastAccessTimeNEQ(v time.Time) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastAccessTime), v))
	})
}

// LastAccessTimeIn applies the In predicate on the "last_access_time" field.
func LastAccessTimeIn(vs ...time.Time) predicate.ExtensionClient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtensionClient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastAccessTime), v...))
	})
}

// LastAccessTimeNotIn applies the NotIn predicate on the "last_access_time" field.
func LastAccessTimeNotIn(vs ...time.Time) predicate.ExtensionClient {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ExtensionClient(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastAccessTime), v...))
	})
}

// LastAccessTimeGT applies the GT predicate on the "last_access_time" field.
func LastAccessTimeGT(v time.Time) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastAccessTime), v))
	})
}

// LastAccessTimeGTE applies the GTE predicate on the "last_access_time" field.
func LastAccessTimeGTE(v time.Time) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastAccessTime), v))
	})
}

// LastAccessTimeLT applies the LT predicate on the "last_access_time" field.
func LastAccessTimeLT(v time.Time) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastAccessTime), v))
	})
}

// LastAccessTimeLTE applies the LTE predicate on the "last_access_time" field.
func LastAccessTimeLTE(v time.Time) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastAccessTime), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ExtensionClient) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ExtensionClient) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
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
func Not(p predicate.ExtensionClient) predicate.ExtensionClient {
	return predicate.ExtensionClient(func(s *sql.Selector) {
		p(s.Not())
	})
}
