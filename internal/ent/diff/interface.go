package diff

// Diff defines methods to compare two SFGA datasets.
type Diff interface {
	// Compare takes two SFGA archives and compares their data.
	// The result is saved internally to SFGA database.
	Compare(src, dst string) error
}
