package db

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

const (
	DB_RECIPIENTS = "recipients"

	FULL_NAME_COLUMN = "full_name"
	EMAIL_COLUMN     = "email"
	VERFIIED_COLUMN  = "verified"
)

var (
	ErrIdAndEmailDontMatch = errors.New("id and email do not match")
)

type RecipientInfo struct {
	Id       string
	FullName string
	Email    string
	Verified bool
}

func (d *DatabaseClient) GetVerifiedRecipients() ([]RecipientInfo, error) {
	recipients, err := d.Db.Dao().FindRecordsByFilter(
		DB_RECIPIENTS,
		fmt.Sprintf("%v = true", VERFIIED_COLUMN),
		"-created",
		0,
		0,
	)
	if err != nil {
		return nil, err
	}

	allRecipients := make([]RecipientInfo, len(recipients))
	for index := range recipients {
		allRecipients[index] = parseRecipient(recipients[index])
	}

	return allRecipients, nil
}

func (d *DatabaseClient) CreateRecipient(fullName, email string) error {
	collection, err := d.Db.Dao().FindCollectionByNameOrId(DB_RECIPIENTS)
	if err != nil {
		return err
	}
	record := models.NewRecord(collection)
	form := forms.NewRecordUpsert(d.Db, record)

	form.LoadData(
		map[string]any{
			FULL_NAME_COLUMN: fullName,
			EMAIL_COLUMN:     strings.ToLower(email),
		},
	)
	return form.Submit()
}

func (d *DatabaseClient) DeleteRecipient(id, email string) error {
	record, err := d.Db.Dao().FindRecordById(DB_RECIPIENTS, id)
	if err != nil {
		return err
	}

	if record.Email() != strings.ToLower(email) {
		return ErrIdAndEmailDontMatch
	}

	return d.Db.Dao().DeleteRecord(record)
}

func (d *DatabaseClient) VerifyRecipient(id, email string) error {
	record, err := d.Db.Dao().FindRecordById(DB_RECIPIENTS, id)
	if err != nil {
		return err
	}

	if record.Email() != strings.ToLower(email) {
		return ErrIdAndEmailDontMatch
	}

	record.Set(VERFIIED_COLUMN, true)

	return d.Db.Dao().SaveRecord(record)
}

func parseRecipient(recipient *models.Record) RecipientInfo {
	return RecipientInfo{
		Id:       recipient.Id,
		FullName: recipient.GetString(FULL_NAME_COLUMN),
		Email:    recipient.Email(),
		Verified: recipient.GetBool(VERFIIED_COLUMN),
	}
}
