package jsondb

import (
	"path/filepath"

	"github.com/pnkj-kmr/looking-glass/utils"
	"go.uber.org/zap"
)

const fileExtension string = ".json"

// storagePath default json storage location
var storagePath string = filepath.Join("conf", "data")

func init() {
	_, err := getOrCreateDir(storagePath)
	if err != nil {
		utils.L.Fatal("Permission error", zap.String("err", err.Error()), zap.String("path", storagePath))
	}
}
