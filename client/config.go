package client

import "errors"

type Config struct {
	BaseURL  string
	Username string
	Password string
	// Default workspace used if one isn't provided on request
	Workspace *string
	// Default repository used if one isn't provided
	Repository *string
}

func (c Config) GetBaseURL() string {
	if c.BaseURL != "" {
		return c.BaseURL
	}
	return "https://api.bitbucket.org"
}

func (c Config) GetUsername() string {
	return c.Username
}

func (c Config) GetPassword() string {
	return c.Password
}

func (c Config) GetWorkspace(w *string) (*string, error) {
	if w != nil {
		return w, nil
	}

	if c.Workspace != nil {
		return c.Workspace, nil
	}

	return nil, errors.New("no workspace set")
}

func (c Config) GetRepository(r *string) (*string, error) {
	if r != nil {
		return r, nil
	}

	if c.Repository != nil {
		return c.Repository, nil
	}

	return nil, errors.New("no repository set")
}

// GetWorkspaceAndRepository is a helper to minimise repeating GetWorkspace and GetRepository
func (c Config) GetWorkspaceAndRepository(workspace, repository *string) (w, r *string, err error) {
	w, err = c.GetWorkspace(workspace)
	if err != nil {
		return
	}

	r, err = c.GetRepository(repository)
	return
}