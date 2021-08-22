package repository

const (
	// Create Task credentials:
	TaskCreateCredential string = "task:create" // for all roles

	// Update Task credentials:
	TaskUpdateCredential    string = "task:update"     // all exist tasks
	TaskOwnUpdateCredential string = "task:own:update" // only own tasks

	// Delete Task credentials:
	TaskDeleteCredential    string = "task:delete"     // for all roles
	TaskOwnDeleteCredential string = "task:own:delete" // only own tasks
)
