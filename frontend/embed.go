package frontend

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var distFS embed.FS

// DistDirFS contains the embedded dist directory files (without the "dist" prefix)
var DistFS, _ = fs.Sub(distFS, "dist")
