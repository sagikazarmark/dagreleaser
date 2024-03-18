package main

import (
	"context"
)

type Staticfile struct {
	File *File
}

func New(file *File) *Staticfile {
	return &Staticfile{
		File: file,
	}
}

func (m *Staticfile) Build(ctx context.Context, platform Platform) *File {
	return m.File
}
