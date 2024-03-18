package main

import "context"

type Demo struct{}

// Returns a container that echoes whatever string argument is provided
func (m *Demo) Demo(ctx context.Context) (*Directory, error) {
	releaser := dag.Releaser()

	archives, err := releaser.Release(
		ctx,
		"hello",
		"v0.0.1",
		[]string{
			"linux/amd64",
			"windows/amd64",
			"darwin/arm64",
		},
		[]*ReleaserFileBuilder{
			dag.Gobinary("hello", dag.CurrentModule().Source().Directory("hello")).AsReleaserFileBuilder(),
			dag.Staticfile(dag.CurrentModule().Source().Directory("hello").File("README.md")).AsReleaserFileBuilder(),
		},
		[]*ReleaserDirectoryBuilder{},
	)
	if err != nil {
		return nil, err
	}

	dir := dag.Directory()

	for _, archive := range archives {
		dir = dir.WithFile("", &archive)
	}

	return dir, nil
}
