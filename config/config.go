package config

import "go.uber.org/zap"

func InitLogs() {
	z, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(z)
	_ = z.Sync()
}
