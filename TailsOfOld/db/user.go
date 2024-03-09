package db

const (
	DB_USERS = "users"

	USERNAME_COLUMN = "username"
	EMAIL_COLUMN    = "email"
	NAME_COLUMN     = "name"
	AVATAR_COLUMN   = "avatar"
)

type UserInfo struct {
	Username   string
	Name       string
	Email      string
	AvatarPath string
}
