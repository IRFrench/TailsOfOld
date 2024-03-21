package filesystem

import (
	"embed"
	"io/fs"
)

//go:embed tailsofold/static/*
var FileSystem embed.FS

//go:embed database/pb_data
var DatabaseFiles embed.FS

//go:embed tailsofold/static/templates/mail/newsletter.txt
var Newsletter string

//go:embed tailsofold/static/templates/mail/verify.txt
var Verify string

func Static() (fs.FS, error) {
	return fs.Sub(FileSystem, "tailsofold")
}
