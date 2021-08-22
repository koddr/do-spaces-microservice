package utils

import (
	"fmt"

	"Komentory/cdn/pkg/repository"
)

// GetCredentialsByRole func for getting credentials from a role name.
func GetCredentialsByRole(role string) ([]string, error) {
	// Define credentials variable.
	var credentials []string

	// Switch given role.
	switch role {
	case repository.AdminRoleName:
		// Admin credentials (all access).
		credentials = []string{
			// Create, Update, Delete any projects:
			repository.ProjectCreateCredential,
			repository.ProjectUpdateCredential,
			repository.ProjectDeleteCredential,
			// Update, Delete own projects:
			repository.ProjectOwnUpdateCredential,
			repository.ProjectOwnDeleteCredential,
			// Create, Update, Delete any tasks:
			repository.TaskCreateCredential,
			repository.TaskUpdateCredential,
			repository.TaskDeleteCredential,
			// Update, Delete own tasks:
			repository.TaskOwnUpdateCredential,
			repository.TaskOwnDeleteCredential,
			// Create, Update any answers:
			repository.AnswerCreateCredential,
			repository.AnswerUpdateCredential,
			repository.AnswerDeleteCredential,
			// Update, Delete own answers:
			repository.AnswerOwnUpdateCredential,
			repository.AnswerOwnDeleteCredential,
			// Create, Update, Delete any users:
			repository.UserCreateCredential,
			repository.UserUpdateCredential,
			repository.UserDeleteCredential,
			// Update own profile:
			repository.UserOwnUpdateCredential,
		}
	case repository.ModeratorRoleName:
		// Moderator credentials (only project creation and update).
		credentials = []string{
			// Create, Update any projects:
			repository.ProjectCreateCredential,
			repository.ProjectUpdateCredential,
			// Update, Delete own projects:
			repository.ProjectOwnUpdateCredential,
			repository.ProjectOwnDeleteCredential,
			// Create, Update any tasks:
			repository.TaskCreateCredential,
			repository.TaskUpdateCredential,
			// Update, Delete own tasks:
			repository.TaskOwnUpdateCredential,
			repository.TaskOwnDeleteCredential,
			// Create, Update any answers:
			repository.AnswerCreateCredential,
			repository.AnswerUpdateCredential,
			// Update, Delete own answers:
			repository.AnswerOwnUpdateCredential,
			repository.AnswerOwnDeleteCredential,
			// Update any users:
			repository.UserUpdateCredential,
			// Update own profile:
			repository.UserOwnUpdateCredential,
		}
	case repository.UserRoleName:
		// Simple user credentials (only project creation).
		credentials = []string{
			// Create, Update, Delete only own projects:
			repository.ProjectCreateCredential,
			repository.ProjectOwnUpdateCredential,
			repository.ProjectOwnDeleteCredential,
			// Create any tasks:
			repository.TaskCreateCredential,
			// Update, Delete own tasks:
			repository.TaskOwnUpdateCredential,
			repository.TaskOwnDeleteCredential,
			// Create, Update any answers:
			repository.AnswerCreateCredential,
			// Update, Delete own answers:
			repository.AnswerOwnUpdateCredential,
			repository.AnswerOwnDeleteCredential,
			// Update own profile:
			repository.UserOwnUpdateCredential,
		}
	default:
		// Return error message.
		return nil, fmt.Errorf("role '%v' does not exist", role)
	}

	return credentials, nil
}
