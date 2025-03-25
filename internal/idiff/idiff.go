package idiff

import (
	"github.com/google/uuid"
	"github.com/sfborg/sf/pkg/config"
	"github.com/sfborg/sf/pkg/diff"
	"github.com/sfborg/sflib/pkg/sfga"
)

type idiff struct {
	cfg      config.Config
	src, ref sfga.Archive
	refUUID  uuid.UUID
	workDir  string
	matcher  diff.Matcher
}

func New(cfg config.Config) diff.Diff {
	res := idiff{cfg: cfg}
	return &res
}
