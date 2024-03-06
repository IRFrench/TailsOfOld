package main

import (
	"TailsOfOld/TailsOfOld/articles"
	"flag"
	"time"
)

func main() {
	var title string
	var author string
	var imagePath string
	var section string

	flag.StringVar(&title, "t", "", "Title of the article")
	flag.StringVar(&author, "a", "", "Author of the article")
	flag.StringVar(&imagePath, "i", "", "Path to an image for the article")
	flag.StringVar(&section, "s", "", "Section of website to add article")

	flag.Parse()

	articleTracker := articles.NewArticleTracker("articles.json", "TailsOfOld/static/templates/articles")
	if err := articleTracker.ReadArticlesFile(); err != nil {
		panic(err)
	}

	if section == string(articles.GAMES_SECTION) {
		if err := articleTracker.CreateArticle(time.Now(), title, author, imagePath, articles.GAMES_SECTION); err != nil {
			panic(err)
		}
		return
	}
	if section == string(articles.PROGRAMMING_SECTION) {
		if err := articleTracker.CreateArticle(time.Now(), title, author, imagePath, articles.PROGRAMMING_SECTION); err != nil {
			panic(err)
		}
		return
	}
	if section == string(articles.TALES_SECTION) {
		if err := articleTracker.CreateArticle(time.Now(), title, author, imagePath, articles.TALES_SECTION); err != nil {
			panic(err)
		}
		return
	}
	panic("Section does not exist")
}
