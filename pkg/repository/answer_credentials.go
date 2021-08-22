package repository

const (
	// Create Answer credentials:
	AnswerCreateCredential string = "answer:create" // for all roles

	// Update Answer credentials:
	AnswerUpdateCredential    string = "answer:update"     // all exist answers
	AnswerOwnUpdateCredential string = "answer:own:update" // only own answers

	// Delete Answer credentials:
	AnswerDeleteCredential    string = "answer:delete"     // for all roles
	AnswerOwnDeleteCredential string = "answer:own:delete" // only own answers
)
