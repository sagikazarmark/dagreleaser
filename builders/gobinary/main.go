package main

import (
	"context"
)

type Gobinary struct {
	Name   string
	Source *Directory
}

func New(name string, source *Directory) *Gobinary {
	return &Gobinary{
		Name:   name,
		Source: source,
	}
}

func (m *Gobinary) Build(ctx context.Context, platform Platform) *File {
	return dag.Go().Build(m.Source, GoBuildOpts{
		Name:     m.Name,
		Trimpath: true,
		Platform: string(platform),

		// Add other args
	})
}
