package fdwca

import (
	"path/filepath"
	"strings"
)

func (fd *fdwca) importExtensions() error {
	for i, ext := range fd.dwca.Meta().Extensions {
		if ext == nil {
			continue
		}

		rowType := filepath.Base(ext.RowType)
		rowType = strings.ToLower(rowType)
		if strings.Contains(rowType, "vernacular") {
			fd.importVernacular(i, ext)
		} else if strings.Contains(rowType, "distribution") {
			fd.importDistr(i, ext)
		} else {
			continue
		}
	}
	return nil
}
