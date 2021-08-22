package repository

const (
	// Create Project credentials:
	ProjectCreateCredential string = "project:create" // for all roles

	// Update Project credentials:
	ProjectUpdateCredential    string = "project:update"     // all exist projects
	ProjectOwnUpdateCredential string = "project:own:update" // only own projects

	// Delete Project credentials:
	ProjectDeleteCredential    string = "project:delete"     // all exist projects
	ProjectOwnDeleteCredential string = "project:own:delete" // only own projects
)
