package filesystem

import (
	"embed"
	"io/fs"
)

//go:embed tailsofold/static/*
var FileSystem embed.FS

func Static() (fs.FS, error) {
	return fs.Sub(FileSystem, "tailsofold")
}
