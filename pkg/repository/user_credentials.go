package repository

const (
	// Create User credentials:
	UserCreateCredential string = "user:create" // for all roles

	// Update User credentials:
	UserUpdateCredential    string = "user:update"     // all exist users
	UserOwnUpdateCredential string = "user:own:update" // only own users

	// Delete User credentials:
	UserDeleteCredential string = "user:delete" // for all roles
)
