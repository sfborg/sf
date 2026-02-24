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
			if err := fd.importVernacular(i); err != nil {
				return err
			}
		} else if strings.Contains(rowType, "distribution") {
			if err := fd.importDistr(i); err != nil {
				return err
			}
		} else {
			continue
		}
	}
	return nil
}
