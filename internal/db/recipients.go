package db

import "github.com/pocketbase/pocketbase/models"

const (
	DB_RECIPIENTS = "recipients"

	FULL_NAME_COLUMN = "full_name"
)

type RecipientInfo struct {
	FullName string
	Email    string
}

func (d *DatabaseClient) GetVerifiedRecipients() ([]RecipientInfo, error) {
	recipients, err := d.Db.Dao().FindRecordsByExpr(DB_RECIPIENTS)
	if err != nil {
		return nil, err
	}

	allRecipients := make([]RecipientInfo, len(recipients))
	for index := range recipients {
		allRecipients[index] = parseRecipient(recipients[index])
	}

	return allRecipients, nil
}

func parseRecipient(recipient *models.Record) RecipientInfo {
	return RecipientInfo{
		FullName: recipient.GetString(FULL_NAME_COLUMN),
		Email:    recipient.Email(),
	}
}
