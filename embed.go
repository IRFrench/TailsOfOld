package filesystem

import "embed"

//go:embed TailsOfOld/static/*
var FileSystem embed.FS
