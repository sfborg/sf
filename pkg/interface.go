package sf

import "github.com/sfborg/sflib/pkg/sfga"

// FromX provides methods for converting a file in a supported
// format to Species File Group Archive format.
type FromX interface {
	// Download checks if input string is a URL. If yes, it downloads the file to
	// internal cache and returns the path to the downloaded file. If it is a
	// local file, it does nothing and returns the same path as the input. It
	// returns an error is some part of the process fails.
	Download(src string) (string, error)

	// Extract takes a path to the source file. If the file is compressed, it
	// determines the type of compression and extracts data to a separate cache
	// directory. If not, it copies the file to the cache directory. It returns
	// an error if some part of the process fails.
	Extract(src string) error

	// InitSfga creates an empty sfga Archive database. The
	// SFGA object's database should already have a connection. If any step
	// of the initiation is broken, it returns the corresponding error
	// message.
	InitSfga() (sfga.Archive, error)

	// Import takes a source file (either a local path or a URL) and an output
	// (a local path). It transforms the source data into the
	// Species File Group Archive (SFGA) format and saves the results in the
	// output directory. The output directory will contain files with '.sql'
	// extension for SQL dump and '.sqlite' extension for a SQLite binary
	// database.  If compression is enabled, these files will be further
	// compressed into '.zip' archives.
	// It returns an error if some part of the process fails.
	Import(src, output string) error
}
