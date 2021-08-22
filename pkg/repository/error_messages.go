package repository

const (
	// User messages:
	NotFoundUserWithID       string = "no user with the specified ID was found"
	NotFoundUserWithEmail    string = "no user with the specified email was found"
	NotFoundUserWithUsername string = "no user with the specified username was found"
	WrongUserEmailOrPassword string = "wrong user email address or password"
	PasswordsDoesNotMatch    string = "passwords does not match"

	// Project messages:
	NotFoundProjects              string = "no projects on the platform were found"
	NotFoundProjectsByUser        string = "no projects for the specified user were found"
	NotFoundProjectWithID         string = "no project with the specified ID was found"
	NotFoundProjectWithAlias      string = "no project with the specified alias was found"
	PermissionDeniedUpdateProject string = "permission denied, only the project creator can update his project"
	PermissionDeniedDeleteProject string = "permission denied, only the project creator can delete his project"

	// Task messages:
	NotFoundTasksByProject     string = "no tasks for the specified project were found"
	NotFoundTaskWithID         string = "no task for the specified ID was found"
	PermissionDeniedCreateTask string = "permission denied, only the project creator can create a task for his project"
	PermissionDeniedUpdateTask string = "permission denied, only the task creator can update his task"
	PermissionDeniedDeleteTask string = "permission denied, only the task creator can delete his task"

	// Task messages:
	NotFoundAnswersByProject     string = "no answers for the specified project were found"
	NotFoundAnswersByTask        string = "no answers for the specified task were found"
	NotFoundAnswerWithID         string = "no answer for the specified ID was found"
	PermissionDeniedUpdateAnswer string = "permission denied, only the answer creator can update his answer"
	PermissionDeniedDeleteAnswer string = "permission denied, only the answer creator can delete his answer"

	// Token messages:
	UnauthorizedCredentials    string = "unauthorized, check credentials of your token"
	UnauthorizedSessionEnded   string = "unauthorized, check expiration time of your refresh token"
	UnauthorizedExpirationTime string = "unauthorized, check expiration time of your access token"
)
