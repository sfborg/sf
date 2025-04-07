package from

import (
	"log/slog"

	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sflib"
	"github.com/sfborg/sflib/pkg/sfga"
)

// Shared contains methods that are common between different
// 'from' convertors.
type Shared struct {
	cfg config.Config
}

func New(cfg config.Config) *Shared {
	res := Shared{cfg: cfg}
	return &res
}

func (fs *Shared) InitSfga() (sfga.Archive, error) {
	slog.Info("Creating SFGA database")
	sfga := sflib.NewSfga()
	err := sfga.Create(fs.cfg.OutputDir)
	if err != nil {
		return nil, err
	}
	_, err = sfga.Connect()
	if err != nil {
		return nil, err
	}
	return sfga, nil
}
