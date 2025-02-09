package api

import (
	"context"
	"encoding/json"
)

// RepositoryRaw is the store model for a Repository.
// Fields have exactly the same meanings as Repository.
type RepositoryRaw struct {
	ID int

	// Standard fields
	CreatorID int
	CreatedTs int64
	UpdaterID int
	UpdatedTs int64

	// Related fields
	VCSID     int
	ProjectID int

	// Domain specific fields
	Name               string
	FullPath           string
	WebURL             string
	BranchFilter       string
	BaseDirectory      string
	FilePathTemplate   string
	SchemaPathTemplate string
	ExternalID         string
	ExternalWebhookID  string
	WebhookURLHost     string
	WebhookEndpointID  string
	WebhookSecretToken string
	AccessToken        string
	ExpiresTs          int64
	RefreshToken       string
}

// ToRepository creates an instance of Repository based on the RepositoryRaw.
// This is intended to be called when we need to compose a Repository relationship.
func (raw *RepositoryRaw) ToRepository() *Repository {
	return &Repository{
		ID: raw.ID,

		CreatorID: raw.CreatorID,
		CreatedTs: raw.CreatedTs,
		UpdaterID: raw.UpdaterID,
		UpdatedTs: raw.UpdatedTs,

		VCSID:     raw.VCSID,
		ProjectID: raw.ProjectID,

		Name:               raw.Name,
		FullPath:           raw.FullPath,
		WebURL:             raw.WebURL,
		BranchFilter:       raw.BranchFilter,
		BaseDirectory:      raw.BaseDirectory,
		FilePathTemplate:   raw.FilePathTemplate,
		SchemaPathTemplate: raw.SchemaPathTemplate,
		ExternalID:         raw.ExternalID,
		ExternalWebhookID:  raw.ExternalWebhookID,
		WebhookURLHost:     raw.WebhookURLHost,
		WebhookEndpointID:  raw.WebhookEndpointID,
		WebhookSecretToken: raw.WebhookSecretToken,
		AccessToken:        raw.AccessToken,
		ExpiresTs:          raw.ExpiresTs,
		RefreshToken:       raw.RefreshToken,
	}
}

// Repository is the API message for a repository.
type Repository struct {
	ID int `jsonapi:"primary,repository"`

	// Standard fields
	CreatorID int
	Creator   *Principal `jsonapi:"relation,creator"`
	CreatedTs int64      `jsonapi:"attr,createdTs"`
	UpdaterID int
	Updater   *Principal `jsonapi:"relation,updater"`
	UpdatedTs int64      `jsonapi:"attr,updatedTs"`

	// Related fields
	VCSID     int
	VCS       *VCS `jsonapi:"relation,vcs"`
	ProjectID int
	Project   *Project `jsonapi:"relation,project"`

	// Domain specific fields
	Name          string `jsonapi:"attr,name"`
	FullPath      string `jsonapi:"attr,fullPath"`
	WebURL        string `jsonapi:"attr,webUrl"`
	BranchFilter  string `jsonapi:"attr,branchFilter"`
	BaseDirectory string `jsonapi:"attr,baseDirectory"`
	// The file path template for matching the committed migration script.
	FilePathTemplate string `jsonapi:"attr,filePathTemplate"`
	// The file path template for storing the latest schema auto-generated by Bytebase after migration.
	// If empty, then Bytebase won't auto generate it.
	SchemaPathTemplate string `jsonapi:"attr,schemaPathTemplate"`
	ExternalID         string `jsonapi:"attr,externalId"`
	ExternalWebhookID  string
	WebhookURLHost     string
	WebhookEndpointID  string
	WebhookSecretToken string
	// These will be exclusively used on the server side and we don't return it to the client.
	AccessToken  string
	ExpiresTs    int64
	RefreshToken string
}

// RepositoryCreate is the API message for creating a repository.
type RepositoryCreate struct {
	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	CreatorID int

	// Related fields
	VCSID     int `jsonapi:"attr,vcsId"`
	ProjectID int

	// Domain specific fields
	Name               string `jsonapi:"attr,name"`
	FullPath           string `jsonapi:"attr,fullPath"`
	WebURL             string `jsonapi:"attr,webUrl"`
	BranchFilter       string `jsonapi:"attr,branchFilter"`
	BaseDirectory      string `jsonapi:"attr,baseDirectory"`
	FilePathTemplate   string `jsonapi:"attr,filePathTemplate"`
	SchemaPathTemplate string `jsonapi:"attr,schemaPathTemplate"`
	ExternalID         string `jsonapi:"attr,externalId"`
	// Token belonged by the user linking the project to the VCS repository. We store this token together
	// with the refresh token in the new repository record so we can use it to call VCS API on
	// behalf of that user to perform tasks like webhook CRUD later.
	AccessToken        string `jsonapi:"attr,accessToken"`
	ExpiresTs          int64  `jsonapi:"attr,expiresTs"`
	RefreshToken       string `jsonapi:"attr,refreshToken"`
	ExternalWebhookID  string
	WebhookURLHost     string
	WebhookEndpointID  string
	WebhookSecretToken string
}

// RepositoryFind is the API message for finding repositories.
type RepositoryFind struct {
	ID *int

	// Related fields
	VCSID     *int
	ProjectID *int

	// Domain specific fields
	WebhookEndpointID *string
}

func (find *RepositoryFind) String() string {
	str, err := json.Marshal(*find)
	if err != nil {
		return err.Error()
	}
	return string(str)
}

// RepositoryPatch is the API message for patching a repository.
type RepositoryPatch struct {
	ID int `jsonapi:"primary,repositoryPatch"`

	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	UpdaterID int

	// Domain specific fields
	BranchFilter       *string `jsonapi:"attr,branchFilter"`
	BaseDirectory      *string `jsonapi:"attr,baseDirectory"`
	FilePathTemplate   *string `jsonapi:"attr,filePathTemplate"`
	SchemaPathTemplate *string `jsonapi:"attr,schemaPathTemplate"`
	AccessToken        *string
	ExpiresTs          *int64
	RefreshToken       *string
}

// RepositoryDelete is the API message for deleting a repository.
type RepositoryDelete struct {
	// Related fields
	// When deleting the repository, we need to update the corresponding project workflow type to "UI",
	// thus we use ProjectID here.
	ProjectID int

	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	DeleterID int
}

// RepositoryService is the service for repositories.
type RepositoryService interface {
	CreateRepository(ctx context.Context, create *RepositoryCreate) (*RepositoryRaw, error)
	FindRepositoryList(ctx context.Context, find *RepositoryFind) ([]*RepositoryRaw, error)
	FindRepository(ctx context.Context, find *RepositoryFind) (*RepositoryRaw, error)
	PatchRepository(ctx context.Context, patch *RepositoryPatch) (*RepositoryRaw, error)
	DeleteRepository(ctx context.Context, delete *RepositoryDelete) error
}
