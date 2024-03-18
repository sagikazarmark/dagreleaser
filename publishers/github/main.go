package main

import "context"

type Github struct {
	Token   *Secret
	Repo    string
	Version string
}

func New(
	token *Secret,
	repo string,
	version string,
) *Github {
	return &Github{
		Token:   token,
		Repo:    repo,
		Version: version,
	}
}

// Returns a container that echoes whatever string argument is provided
func (m *Github) Publish(ctx context.Context, artifacts []*File) error {
	_, err := dag.Gh(GhOpts{
		Token: m.Token,
		Repo:  m.Repo,
	}).Release().Create(m.Version, m.Version, GhReleaseCreateOpts{
		Files:         artifacts,
		GenerateNotes: true,
		Latest:        true,
		VerifyTag:     true,
	}).Sync(ctx)
	if err != nil {
		return err
	}

	return nil
}
