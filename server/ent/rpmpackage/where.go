// Code generated by ent, DO NOT EDIT.

package rpmpackage

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/FyraLabs/subatomic/server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEQ(FieldName, v))
}

// Epoch applies equality check predicate on the "epoch" field. It's identical to EpochEQ.
func Epoch(v int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEQ(FieldEpoch, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEQ(FieldVersion, v))
}

// Release applies equality check predicate on the "release" field. It's identical to ReleaseEQ.
func Release(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEQ(FieldRelease, v))
}

// Arch applies equality check predicate on the "arch" field. It's identical to ArchEQ.
func Arch(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEQ(FieldArch, v))
}

// FilePath applies equality check predicate on the "file_path" field. It's identical to FilePathEQ.
func FilePath(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEQ(FieldFilePath, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldContainsFold(FieldName, v))
}

// EpochEQ applies the EQ predicate on the "epoch" field.
func EpochEQ(v int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEQ(FieldEpoch, v))
}

// EpochNEQ applies the NEQ predicate on the "epoch" field.
func EpochNEQ(v int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldNEQ(FieldEpoch, v))
}

// EpochIn applies the In predicate on the "epoch" field.
func EpochIn(vs ...int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldIn(FieldEpoch, vs...))
}

// EpochNotIn applies the NotIn predicate on the "epoch" field.
func EpochNotIn(vs ...int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldNotIn(FieldEpoch, vs...))
}

// EpochGT applies the GT predicate on the "epoch" field.
func EpochGT(v int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldGT(FieldEpoch, v))
}

// EpochGTE applies the GTE predicate on the "epoch" field.
func EpochGTE(v int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldGTE(FieldEpoch, v))
}

// EpochLT applies the LT predicate on the "epoch" field.
func EpochLT(v int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldLT(FieldEpoch, v))
}

// EpochLTE applies the LTE predicate on the "epoch" field.
func EpochLTE(v int) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldLTE(FieldEpoch, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldLTE(FieldVersion, v))
}

// VersionContains applies the Contains predicate on the "version" field.
func VersionContains(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldContains(FieldVersion, v))
}

// VersionHasPrefix applies the HasPrefix predicate on the "version" field.
func VersionHasPrefix(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldHasPrefix(FieldVersion, v))
}

// VersionHasSuffix applies the HasSuffix predicate on the "version" field.
func VersionHasSuffix(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldHasSuffix(FieldVersion, v))
}

// VersionEqualFold applies the EqualFold predicate on the "version" field.
func VersionEqualFold(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEqualFold(FieldVersion, v))
}

// VersionContainsFold applies the ContainsFold predicate on the "version" field.
func VersionContainsFold(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldContainsFold(FieldVersion, v))
}

// ReleaseEQ applies the EQ predicate on the "release" field.
func ReleaseEQ(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEQ(FieldRelease, v))
}

// ReleaseNEQ applies the NEQ predicate on the "release" field.
func ReleaseNEQ(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldNEQ(FieldRelease, v))
}

// ReleaseIn applies the In predicate on the "release" field.
func ReleaseIn(vs ...string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldIn(FieldRelease, vs...))
}

// ReleaseNotIn applies the NotIn predicate on the "release" field.
func ReleaseNotIn(vs ...string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldNotIn(FieldRelease, vs...))
}

// ReleaseGT applies the GT predicate on the "release" field.
func ReleaseGT(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldGT(FieldRelease, v))
}

// ReleaseGTE applies the GTE predicate on the "release" field.
func ReleaseGTE(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldGTE(FieldRelease, v))
}

// ReleaseLT applies the LT predicate on the "release" field.
func ReleaseLT(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldLT(FieldRelease, v))
}

// ReleaseLTE applies the LTE predicate on the "release" field.
func ReleaseLTE(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldLTE(FieldRelease, v))
}

// ReleaseContains applies the Contains predicate on the "release" field.
func ReleaseContains(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldContains(FieldRelease, v))
}

// ReleaseHasPrefix applies the HasPrefix predicate on the "release" field.
func ReleaseHasPrefix(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldHasPrefix(FieldRelease, v))
}

// ReleaseHasSuffix applies the HasSuffix predicate on the "release" field.
func ReleaseHasSuffix(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldHasSuffix(FieldRelease, v))
}

// ReleaseEqualFold applies the EqualFold predicate on the "release" field.
func ReleaseEqualFold(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEqualFold(FieldRelease, v))
}

// ReleaseContainsFold applies the ContainsFold predicate on the "release" field.
func ReleaseContainsFold(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldContainsFold(FieldRelease, v))
}

// ArchEQ applies the EQ predicate on the "arch" field.
func ArchEQ(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEQ(FieldArch, v))
}

// ArchNEQ applies the NEQ predicate on the "arch" field.
func ArchNEQ(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldNEQ(FieldArch, v))
}

// ArchIn applies the In predicate on the "arch" field.
func ArchIn(vs ...string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldIn(FieldArch, vs...))
}

// ArchNotIn applies the NotIn predicate on the "arch" field.
func ArchNotIn(vs ...string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldNotIn(FieldArch, vs...))
}

// ArchGT applies the GT predicate on the "arch" field.
func ArchGT(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldGT(FieldArch, v))
}

// ArchGTE applies the GTE predicate on the "arch" field.
func ArchGTE(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldGTE(FieldArch, v))
}

// ArchLT applies the LT predicate on the "arch" field.
func ArchLT(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldLT(FieldArch, v))
}

// ArchLTE applies the LTE predicate on the "arch" field.
func ArchLTE(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldLTE(FieldArch, v))
}

// ArchContains applies the Contains predicate on the "arch" field.
func ArchContains(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldContains(FieldArch, v))
}

// ArchHasPrefix applies the HasPrefix predicate on the "arch" field.
func ArchHasPrefix(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldHasPrefix(FieldArch, v))
}

// ArchHasSuffix applies the HasSuffix predicate on the "arch" field.
func ArchHasSuffix(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldHasSuffix(FieldArch, v))
}

// ArchEqualFold applies the EqualFold predicate on the "arch" field.
func ArchEqualFold(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEqualFold(FieldArch, v))
}

// ArchContainsFold applies the ContainsFold predicate on the "arch" field.
func ArchContainsFold(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldContainsFold(FieldArch, v))
}

// FilePathEQ applies the EQ predicate on the "file_path" field.
func FilePathEQ(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEQ(FieldFilePath, v))
}

// FilePathNEQ applies the NEQ predicate on the "file_path" field.
func FilePathNEQ(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldNEQ(FieldFilePath, v))
}

// FilePathIn applies the In predicate on the "file_path" field.
func FilePathIn(vs ...string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldIn(FieldFilePath, vs...))
}

// FilePathNotIn applies the NotIn predicate on the "file_path" field.
func FilePathNotIn(vs ...string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldNotIn(FieldFilePath, vs...))
}

// FilePathGT applies the GT predicate on the "file_path" field.
func FilePathGT(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldGT(FieldFilePath, v))
}

// FilePathGTE applies the GTE predicate on the "file_path" field.
func FilePathGTE(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldGTE(FieldFilePath, v))
}

// FilePathLT applies the LT predicate on the "file_path" field.
func FilePathLT(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldLT(FieldFilePath, v))
}

// FilePathLTE applies the LTE predicate on the "file_path" field.
func FilePathLTE(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldLTE(FieldFilePath, v))
}

// FilePathContains applies the Contains predicate on the "file_path" field.
func FilePathContains(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldContains(FieldFilePath, v))
}

// FilePathHasPrefix applies the HasPrefix predicate on the "file_path" field.
func FilePathHasPrefix(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldHasPrefix(FieldFilePath, v))
}

// FilePathHasSuffix applies the HasSuffix predicate on the "file_path" field.
func FilePathHasSuffix(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldHasSuffix(FieldFilePath, v))
}

// FilePathEqualFold applies the EqualFold predicate on the "file_path" field.
func FilePathEqualFold(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldEqualFold(FieldFilePath, v))
}

// FilePathContainsFold applies the ContainsFold predicate on the "file_path" field.
func FilePathContainsFold(v string) predicate.RpmPackage {
	return predicate.RpmPackage(sql.FieldContainsFold(FieldFilePath, v))
}

// HasRepo applies the HasEdge predicate on the "repo" edge.
func HasRepo() predicate.RpmPackage {
	return predicate.RpmPackage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RepoTable, RepoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepoWith applies the HasEdge predicate on the "repo" edge with a given conditions (other predicates).
func HasRepoWith(preds ...predicate.Repo) predicate.RpmPackage {
	return predicate.RpmPackage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepoInverseTable, RepoFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RepoTable, RepoColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RpmPackage) predicate.RpmPackage {
	return predicate.RpmPackage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RpmPackage) predicate.RpmPackage {
	return predicate.RpmPackage(func(s *sql.Selector) {
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
func Not(p predicate.RpmPackage) predicate.RpmPackage {
	return predicate.RpmPackage(func(s *sql.Selector) {
		p(s.Not())
	})
}
