package utils

import (
	zrl "github.com/pnkj-kmr/zap-rotate-logger"

	"go.uber.org/zap"
)

// L - application logger
var L *zap.Logger

// SetLogger setting the logger
func SetLogger(debug bool) {
	L = zrl.New(
		zrl.WithFileName("lg"),
		zrl.WithMaxSize(1),
		zrl.WithMaxBackups(30),
		zrl.WithMaxAge(60),
		zrl.WithCompress(true),
		zrl.WithDebug(debug),
	)
}
