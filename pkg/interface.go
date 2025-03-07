package sf

// Importer provides methods for converting a file in a supported
// format to Species File Group Archive format.
type Importer interface {
	// Import takes a src (either a local path or URL), and output
	// (a local path).
	Import(src, output string) error

	// Download is a helper command. It takes a src string  and checks if
	// it is a http-based URL. If yes, it downloads the file from the URL
	// and returns back local path to the downloaded file.
	Download(src string) (string, error)

	// Extract checks if a file is compressed and if yes, it extracts its
	// content to config.Config.ImportSrcDir.
	Extract(src string) (string, error)

	// InitSfga creates an empty sfga.Archive, connects to it and saves the
	// archie to the object's internal state.
	InitSfga() error

	// ExportSfga uses outputPath to create sql dump (.sql extension) and
	// binary database (.sqlite extension) from the resulting sfga.Archive.
	// If config.Config.WithZip is true, it also compresses the files to
	// Zip archives.
	ExportSfga(outputPath string) error
}
