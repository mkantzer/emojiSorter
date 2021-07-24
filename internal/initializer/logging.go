package initializer

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Initializer that prepares logging for our service
//
// If writer is provided the logger is setup to write to the provided io.Writer, this
// is useful in tests.
func Logging(environment string, hostname string, serviceName string, gitHash string,
	writer io.Writer) (*zap.Logger, error) {

	var loggerConfig zap.Config
	if environment == "development" {
		loggerConfig = zap.NewDevelopmentConfig()
	} else {
		loggerConfig = zap.NewProductionConfig()
	}

	// Elasticsearch expects @timestamp and RFC3339
	loggerConfig.EncoderConfig.TimeKey = "@timestamp"
	rfc3339Milli := "2006-01-02T15:04:05.999Z07:00"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(rfc3339Milli)
	loggerConfig.EncoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
	loggerConfig.EncoderConfig.MessageKey = "message"
	loggerConfig.DisableStacktrace = true

	var logger *zap.Logger
	if writer != nil {
		var encoder zapcore.Encoder
		if environment == "development" {
			encoder = zapcore.NewConsoleEncoder(loggerConfig.EncoderConfig)
		} else {
			encoder = zapcore.NewJSONEncoder(loggerConfig.EncoderConfig)
		}

		logger = zap.New(zapcore.NewCore(encoder, zapcore.AddSync(writer), loggerConfig.Level))
	} else {
		var err error
		logger, err = loggerConfig.Build()
		if err != nil {
			return nil, err
		}
	}

	// Our centralized logging system should expect these
	logger = logger.With(
		zap.String("source", serviceName),
		zap.String("environment", environment),
		zap.String("source_version", gitHash),
		zap.String("hostname", hostname),
		zap.Int("pid", os.Getpid()),
	)

	return logger, nil
}
