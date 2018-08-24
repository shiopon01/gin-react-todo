package main

import (
	"net/http"
	"strings"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/shiopon01/gin-react-todo/server/env"
)

// BinaryFileSystem ..
type BinaryFileSystem interface{}

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

// ReadAsset ...
func ReadAsset(root string) *binaryFileSystem {
	fs := &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    root,
	}

	return &binaryFileSystem{
		fs,
	}
}

func main() {
	// Select Development mode or Release mode
	if !env.DEBUG {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Front page display
	r.Use(static.Serve("/", ReadAsset("assets")))

	// API routing
	RouteCollections(r)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
