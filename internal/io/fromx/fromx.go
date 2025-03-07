package fromx

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/gnames/gnsys"
	sf "github.com/sfborg/sf/pkg"
	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sflib/ent/sfga"
	"github.com/sfborg/sflib/io/sfgaio"
)

type fromx struct {
	cfg config.Config
}

func New(cfg config.Config) sf.FromX {
	res := fromx{cfg: cfg}
	return &res
}

// Import is a placeholder to correspond to From interface
func (f *fromx) Import(_, _ string) error {
	return nil
}

func (f *fromx) Download(src string) (string, error) {
	var err error
	if strings.HasPrefix(src, "http") {
		slog.Info("Downloading from URL", "url", src)
		src, err = gnsys.Download(src, f.cfg.DownloadDir, true)
		if err != nil {
			return "", err
		}
	}
	return src, nil
}

func (f *fromx) Extract(src string) error {
	var err error
	var e gnsys.Extractor
	switch gnsys.GetFileType(src) {
	case gnsys.ZipFT:
		e = gnsys.ExtractZip
	case gnsys.TarFT:
		e = gnsys.ExtractTar
	case gnsys.TarGzFT:
		e = gnsys.ExtractTarGz
	case gnsys.TarBzFT:
		e = gnsys.ExtractTarBz2
	case gnsys.TarXzFt:
		e = gnsys.ExtractTarXz
	default:
		err = f.copy(src)
		if err != nil {
			return err
		}
		return nil
	}
	err = e(src, f.cfg.DataDir)
	if err != nil {
		return err
	}
	return nil
}

func (f *fromx) InitSfga() (sfga.Archive, error) {
	slog.Info("Creating SFGA database")
	sfga := sfgaio.New()
	err := sfga.Create(f.cfg.SfgaDir)
	if err != nil {
		return nil, err
	}
	_, err = sfga.Connect()
	if err != nil {
		return nil, err
	}
	return sfga, nil
}

func (f *fromx) copy(src string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstPath := filepath.Join(f.cfg.DataDir, filepath.Base(src))

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
