package initializer

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Logging(environment string, hostname string, serviceName string, gitHash string) (*zap.Logger, error) {
	var loggerConfig zap.Config
	if environment == "development" {
		loggerConfig = zap.NewDevelopmentConfig()
	} else {
		loggerConfig = zap.NewProductionConfig()
	}

	// Elasticsearch expects @timestamp and RFC3339
	loggerConfig.EncoderConfig.TimeKey = "@timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logger, err := loggerConfig.Build()
	if err != nil {
		return nil, err
	}

	// Our centralized logging system should expect these
	logger = logger.With(
		zap.String("service", serviceName),
		zap.String("environment", environment),
		zap.String("build", gitHash),
		zap.String("hostname", hostname),
	)

	return logger, nil
}
