package diffio

import (
	"github.com/google/uuid"
	"github.com/sfborg/sf/internal/ent/diff"
	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sflib/pkg/sfga"
)

type diffio struct {
	cfg      config.Config
	src, ref sfga.Archive
	refUUID  uuid.UUID
	workDir  string
	matcher  diff.Matcher
}

func New(cfg config.Config) diff.Diff {
	res := diffio{cfg: cfg}
	return &res
}
