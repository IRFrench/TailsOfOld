package articles

import (
	filesystem "TailsOfOld"
	"TailsOfOld/internal/db"
	"TailsOfOld/tailsofold/handlers"
	weberrors "TailsOfOld/tailsofold/handlers/web_errors"
	"html/template"
	"net/http"

	"github.com/rs/zerolog/log"
)

type OverviewHandler struct {
	Database *db.DatabaseClient
	Section  string
}

type OverviewVariables struct {
	Section  string
	Articles []db.ArticleInfo
}

func (o OverviewHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// Build template
	templatePath := "TailsOfOld/static/templates/articles/overview.html"
	template := template.New("overview.html")

	template, err := template.ParseFS(filesystem.FileSystem, handlers.BASE_TEMPLATES, templatePath)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse filesystem into template")
		weberrors.Borked(response, request)
		return
	}

	// Gather articles
	databaseArticles, err := o.Database.GetEntireSectionArticleInfo(o.Section)
	if err != nil {
		log.Error().Err(err).Msg("failed to find section articles in the database")
		weberrors.Borked(response, request)
		return
	}

	for _, article := range databaseArticles {
		// Update the author
		articleAuthor, err := o.Database.GetAuthor(article.Author)
		if err != nil {
			log.Error().Err(err).Str("article title", article.Title).Msg("failed to find article author")
			weberrors.Borked(response, request)
			return
		}
		article.Author = articleAuthor.Name
	}

	vars := OverviewVariables{
		Section:  o.Section,
		Articles: databaseArticles,
	}

	if err := template.ExecuteTemplate(response, "base", vars); err != nil {
		log.Error().Err(err).Msg("failed to execute the template")
		weberrors.Borked(response, request)
		return
	}
}
