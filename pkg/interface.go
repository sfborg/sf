package sf

// Importer provides methods for converting a file in a supported
// format to Species File Group Archive format.
type Importer interface {
	// Import takes a src (either a local path or URL), and output
	// (a local path).
	Import(src, output string) error
}
