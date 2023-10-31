package filesystem

import (
	"embed"
	"io/fs"
)

//go:embed TailsOfOld/static/*
var FileSystem embed.FS

func Static() (fs.FS, error) {
	return fs.Sub(FileSystem, "TailsOfOld")
}
