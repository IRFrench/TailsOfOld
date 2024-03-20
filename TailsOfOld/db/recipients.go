package db

const (
	DB_RECIPIENTS = "recipients"

	FULL_NAME_COLUMN = "full_name"
)

type RecipientInfo struct {
	FullName string
	Email    string
	Created  string
	Updated  string
}
