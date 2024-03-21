package db

import (
	"fmt"

	"github.com/pocketbase/pocketbase/models"
)

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

func (d *DatabaseClient) GetAuthor(id string) (UserInfo, error) {
	articleAuthor, err := d.Db.Dao().FindRecordById(
		DB_USERS,
		id,
	)
	if err != nil {
		return UserInfo{}, err
	}
	return parseUser(articleAuthor), nil
}

func parseUser(user *models.Record) UserInfo {
	return UserInfo{
		Username:   user.GetString(USERNAME_COLUMN),
		Email:      user.GetString(EMAIL_COLUMN),
		Name:       user.GetString(NAME_COLUMN),
		AvatarPath: fmt.Sprintf("/pb_data/storage/%v/%v", user.BaseFilesPath(), user.GetString(AVATAR_COLUMN)),
	}
}
