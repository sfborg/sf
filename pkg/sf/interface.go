package sf

// FromConvertor provides methods for converting a file in a supported
// format to Species File Group Archive format.
type FromConvertor interface {
	// Import takes a source file (either a local path or a URL) and an output
	// (a local path). It transforms the source data into the
	// Species File Group Archive (SFGA) format and saves the results in the
	// output directory. The output directory will contain files with '.sql'
	// extension for SQL dump and '.sqlite' extension for a SQLite binary
	// database.  If compression is opted in, these files will be further
	// compressed into '.zip' archives.
	// It returns an error if some part of the process fails.
	Import(src, output string) error
}
