package diff

import "database/sql"

// Diff defines methods to compare two SFGA datasets.
type Diff interface {
	// Compare takes two SFGA archives and compares their data.
	// The result is saved internally to SFGA database.
	Compare(src, ref, out string) error
}

type Matcher interface {
	Init(*sql.DB, []Record) error
	Match(Record) ([]Record, error)
	MatchExact(string) ([]Record, error)
	MatchFuzzy(string, string) ([]Record, error)
}

type DBase interface {
	Init(*sql.DB)
	Select(string) ([]Record, error)
}

type Exact interface {
	Init([]Record)
	Find(string) bool
}

type Fuzzy interface {
	Init([]Record) error
	FindExact(string) []string
	FindFuzzy(string) []string
}
