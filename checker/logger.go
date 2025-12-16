package checker

import (
	"errors"

	"go.uber.org/zap"
)

var logger *zap.Logger

// func init() {
// 	var err error
// 	logger, err = zap.NewProduction()
// 	if err != nil {
// 		panic(err)
// 	}
// }

func InitLogger(mode string) error {
	var cfg zap.Config
	var err error

	switch mode {
	case "json":
		cfg = zap.NewProductionConfig()
	case "text":
		cfg = zap.NewDevelopmentConfig()
		cfg.Encoding = "console"
	default:
		return errors.New("invalid log format")
	}
	logger, err = cfg.Build()
	if err != nil {
		return err
	}
	return nil
}
