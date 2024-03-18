package main

import (
	"context"
	"fmt"
	"strings"

	platformlist "github.com/containerd/containerd/platforms"
)

type Releaser struct {
	// FileBuilders      []FileBuilder
	// DirectoryBuilders []DirectoryBuilder
}

// func New(
// 	fileBuilders []FileBuilder,
// 	directoryBuilders []DirectoryBuilder,
// ) *Releaser {
// 	return &Releaser{
// 		FileBuilders:      fileBuilders,
// 		DirectoryBuilders: directoryBuilders,
// 	}
// }

// FileBuilder executes a build and returns a list of files.
type FileBuilder interface {
	DaggerObject

	Build(ctx context.Context, platform Platform) *File
}

// DirectoryBuilder executes a build and returns a directory.
type DirectoryBuilder interface {
	DaggerObject

	Build(ctx context.Context, platform Platform) *Directory
}

// Archiver archives a directory of files and returns a single archive.
type Archiver interface {
	DaggerObject

	Archive(name string, source *Directory) *File
}

// Publisher publishes release artifacts.
// type Publisher interface {
// 	Publish(artifacts []*File)
// }

func (m *Releaser) Release(
	name string,
	version string,
	platforms []string,

	fileBuilders []FileBuilder,
	directoryBuilders []DirectoryBuilder,
) []*File {
	archives := make([]*File, 0, len(platforms))

	for _, platform := range platforms {
		platform := Platform(platform)

		// Store build artifacts in a new directory
		buildDir := dag.Directory()

		// Run file builders and store the results
		for _, builder := range fileBuilders {
			file := builder.Build(context.Background(), platform)
			buildDir = buildDir.WithFile("", file)
		}

		// Run directory builders and store the results
		for _, builder := range directoryBuilders {
			dir := builder.Build(context.Background(), platform)
			buildDir = buildDir.WithDirectory("", dir)
		}

		parsedPlatform := platformlist.MustParse(string(platform))

		archiveName := fmt.Sprintf("%s_%s_%s", name, parsedPlatform.OS, parsedPlatform.Architecture)

		archiver := m.selectArchiver(platform)

		archive := archiver.Archive(archiveName, buildDir)

		archives = append(archives, archive)
	}

	return archives

	// Publish

	// return nil
}

func (m *Releaser) selectArchiver(platform Platform) Archiver {
	// TODO: custom archiver selector

	if strings.HasPrefix(string(platform), "windows/") {
		return dag.Archivist().Zip()
	}

	return dag.Archivist().TarGz()
}
