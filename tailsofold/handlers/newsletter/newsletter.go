package newsletter

import (
	filesystem "TailsOfOld"
	"TailsOfOld/internal/db"
	mailclient "TailsOfOld/internal/mail_client"
	"TailsOfOld/tailsofold/handlers"
	weberrors "TailsOfOld/tailsofold/handlers/web_errors"
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/rs/zerolog/log"
)

const (
	ErrSomethingWentWrong = "something went wrong"
)

type SubscribeHandler struct {
	Database *db.DatabaseClient
	Mail     *mailclient.MailClient
}

type SubscribeData struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

func (s SubscribeHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	var data SubscribeData
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		log.Err(err).Msg("failed to decode subscribe body")
		http.Error(response, ErrSomethingWentWrong, http.StatusInternalServerError)
		return
	}
	defer request.Body.Close()

	if err := s.Database.CreateRecipient(data.FullName, data.Email); err != nil {
		log.Err(err).Msg("failed to create recipient")
		http.Error(response, ErrSomethingWentWrong, http.StatusInternalServerError)
		return
	}

	recipient, err := s.Database.GetRecipientByNameAndEmail(data.FullName, data.Email)
	if err != nil {
		log.Err(err).Msg("failed to find recipient")
		http.Error(response, ErrSomethingWentWrong, http.StatusInternalServerError)
		return
	}

	go func() {
		log.Info().Str("email", recipient.Email).Msg("sending verification email")
		if err := s.Mail.SendVerification(s.Database, recipient); err != nil {
			log.Err(err).Msg("failed to send verification")
			return
		}
		log.Info().Msg("verification sent")
	}()

	if err := json.NewEncoder(response).Encode(true); err != nil {
		log.Err(err).Msg("failed to encode response")
		http.Error(response, ErrSomethingWentWrong, http.StatusInternalServerError)
	}
}

type VerifyHandler struct {
	Database *db.DatabaseClient
}

func (v VerifyHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// Build template
	email := request.URL.Query().Get("email")
	id := request.URL.Query().Get("id")
	log.Info().Str("email", email).Str("id", id).Msg("User verification")

	templatePath := "tailsofold/static/templates/mail/verified.html"
	template := template.New("verified.html")

	template, err := template.ParseFS(filesystem.FileSystem, handlers.BASE_TEMPLATES, templatePath)
	if err != nil {
		log.Err(err).Msg("failed to parse filesystem into template")
		weberrors.Borked(response, request)
		return
	}

	if err := v.Database.VerifyRecipient(id, email); err != nil {
		log.Err(err).Msg("failed to verify user email")
		weberrors.Borked(response, request)
		return
	}

	if err := template.ExecuteTemplate(response, "base", nil); err != nil {
		log.Err(err).Msg("failed to execute template")
		weberrors.Borked(response, request)
		return
	}
}

type UnsubscribeHandler struct {
	Database *db.DatabaseClient
}

func (u UnsubscribeHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// Build template
	email := request.URL.Query().Get("email")
	id := request.URL.Query().Get("id")

	log.Info().Str("email", email).Str("id", id).Msg("User unsubscription")

	templatePath := "tailsofold/static/templates/mail/unsubscribe.html"
	template := template.New("unsubscribe.html")

	template, err := template.ParseFS(filesystem.FileSystem, handlers.BASE_TEMPLATES, templatePath)
	if err != nil {
		log.Err(err).Msg("failed to parse filesystem into template")
		weberrors.Borked(response, request)
		return
	}

	if err := u.Database.DeleteRecipient(id, email); err != nil {
		log.Err(err).Msg("failed to verify user email")
		weberrors.Borked(response, request)
		return
	}

	if err := template.ExecuteTemplate(response, "base", nil); err != nil {
		log.Err(err).Msg("failed to execute template")
		weberrors.Borked(response, request)
		return
	}
}
