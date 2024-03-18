package main

import (
	"context"
	"dagger/demo/internal/dagger"
)

type Demo struct{}

func (m *Demo) Release(ctx context.Context) (*Directory, error) {
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

func (m *Demo) ReleaseAndPublish(ctx context.Context, token *dagger.Secret) error {
	releaser := dag.Releaser()

	_, err := releaser.ReleaseAndPublish(
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
		dag.Github(token, "sagikazarmark/dagreleaser", "v0.0.1").AsReleaserPublisher(),
	)
	if err != nil {
		return err
	}

	return nil
}
