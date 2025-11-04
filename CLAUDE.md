# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`sf` is a CLI tool for converting biodiversity data between various formats and the Species File Group Archive (SFGA) format. It handles format conversions (from/to SFGA) and performs taxonomic data comparisons between SFGA files.

## Build and Development Commands

The project uses [just](https://github.com/casey/just) as a command runner (similar to make). All commands are defined in `Justfile`.

**Building:**
```bash
just build              # Build with version info from git
just buildrel           # Build release version (stripped, no debug info)
just install            # Build and install to $GOPATH/bin
```

**Testing:**
```bash
just test               # Run all tests with coverage
just coverage           # Run tests and display coverage report
```

**Dependency Management:**
```bash
just deps               # Download dependencies
just tools              # Install development tools from tools.go
```

**Release:**
```bash
just release            # Build cross-platform binaries (Linux, macOS, Windows) to /tmp
just clean              # Remove generated files
```

**Other:**
```bash
just version            # Display current version from git tags
just help               # Show all available commands
```

## Architecture

### Command Structure

Built with [cobra](https://github.com/spf13/cobra), organized hierarchically:

- `sf` - Root command
  - `from` - Convert various formats TO SFGA
    - `text` - From plain text (one name per line)
    - `xsv` - From CSV/TSV/PSV files (with Darwin Core or CoLDP headers)
    - `coldp` - From Catalogue of Life Data Package
    - `dwca` - From Darwin Core Archive
  - `to` - Convert SFGA TO other formats
    - `coldp` - To Catalogue of Life Data Package
    - `dwca` - To Darwin Core Archive
  - `diff` - Compare two SFGA files and output differences
  - `apply` - Apply diff results to an SFGA file
  - `update` - Update SFGA file with new data

All command implementations are in `cmd/` directory.

### Core Abstractions

**Interfaces** (`pkg/sf/interface.go`):
- `FromConvertor` - Implements `Import(src, output string) error` to convert formats to SFGA
- `ToConvertor` - Implements `Export(src, output string) error` to convert SFGA to other formats

**Configuration** (`config/config.go`):
- Central `Config` struct manages all app settings
- Uses functional options pattern (`Option` type)
- Manages cache directories (under `~/.cache/sfborg/sf/`)
- Key settings: `BatchSize`, `JobsNum`, `WithQuotes`, `WithZipOutput`, `WithDetails`

### Data Flow

1. **Import Flow** (from → SFGA):
   - Source file downloaded/accessed → `CacheDir/import/download/`
   - Extracted to → `CacheDir/import/src/`
   - Converted to SFGA → `CacheDir/import/sfga/`
   - Output saved (SQLite + SQL dump, optionally zipped)

2. **Export Flow** (SFGA → to):
   - SFGA file loaded from source (local path or URL)
   - Converted to target format
   - Saved to output location

3. **Diff Flow**:
   - Two SFGA files compared (can filter by specific taxon)
   - Uses bloom filters and suffix tries (stored in `DiffWorkDir`)
   - Multiple matching strategies: exact, fuzzy (edit distance), database-backed
   - Outputs comparison results as new SFGA file

### Package Organization

- `cmd/` - Cobra command definitions
- `pkg/sf/` - Core interfaces and types
- `pkg/from/` - Converters TO SFGA format:
  - `ftext/` - Text file imports
  - `fxsv/` - CSV/TSV/PSV imports
  - `fcoldp/` - CoLDP imports
  - `fdwca/` - DwCA imports
  - `fsfga/` - SFGA-to-SFGA operations
  - `shared.go` - Common import functionality
- `pkg/to/` - Converters FROM SFGA:
  - `tcoldp/` - Export to CoLDP
  - `shared.go` - Common export functionality
- `pkg/diff/` - Diff interfaces
- `internal/idiff/` - Diff implementation:
  - `exactio/` - Exact matching
  - `fuzzyio/` - Fuzzy matching with edit distance
  - `matchio/` - General matching logic
  - `dbio/` - Database-backed matching
- `config/` - Configuration management
- `internal/util/` - System utilities

### Dependencies

The project heavily depends on `github.com/sfborg/sflib` for core SFGA operations:
- `sflib.NewDwca()` - DwCA archive handler
- `sflib.NewColdp()` - CoLDP archive handler
- `sflib/pkg/sfga` - SFGA format implementation
- `sflib/pkg/parser` - Scientific name parsing (uses GNparser)

Other key dependencies:
- `github.com/gnames/gnparser` - Scientific name parsing
- `modernc.org/sqlite` - Pure Go SQLite implementation
- `github.com/devopsfaith/bloomfilter` - Bloom filters for diff operations
- `github.com/dvirsky/levenshtein` - Edit distance calculations

### Key Patterns

1. **Functional Options**: All constructors use `Option` functions for configuration
2. **Interface Segregation**: Small, focused interfaces (`FromConvertor`, `ToConvertor`)
3. **Shared Functionality**: Common logic extracted to `Shared` embedded structs
4. **Pool Pattern**: Parser pools for concurrent operations (`parserPool` in fcoldp)
5. **Cache Management**: Persistent work directories vs temporary cache directories

### Scientific Name Parsing

When `WithDetails: true` (default), the tool uses GNparser to populate detailed name fields:
- Uninomial, Genus, SpecificEpithet, InfraspecificEpithet
- Authorship fields (combination, basionym)
- Nomenclatural code detection
- Rank and status information

Can be disabled with `--no-parse-details` flag for faster imports.

### File Handling

- Supports local files and remote URLs (HTTP/HTTPS)
- Handles compressed archives (.zip, .tar.gz)
- CSV/TSV/PSV files can optionally have quoted fields (`--quotes-allowed`)
- Output can be compressed with `--zip-output` flag

### Diff Implementation Details

The diff system uses multiple matching strategies to handle taxonomic name variations:
1. **Exact matching** - Direct string comparison
2. **Fuzzy matching** - Levenshtein distance for near-matches
3. **Match storage** - Database-backed persistent matching results

Supports scoped comparison by specifying root taxon in each file using `--source-taxon` and `--target-taxon` flags.
