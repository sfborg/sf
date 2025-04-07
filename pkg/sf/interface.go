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

// ToConvertor provides methods for coverting an SFGA file to one of the
// supported formats.
type ToConvertor interface {
	// Export takes a source SFGA file (either a local path or URL) and an
	// output (a local path). It transforms SFGA into a required format and
	// saves the result into the output file. If the result contains several
	// files, they are compressed with Zip. If the result is one file, the
	// Zip conversion is optional.
	// It returns an error if some part of the process fails.
	Export(src, output string) error
}
