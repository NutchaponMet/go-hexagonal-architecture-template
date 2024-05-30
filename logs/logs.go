package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func LogInit() {

	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""
	var err error
	Log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

}

func Info(message string, field ...zapcore.Field) {
	Log.Info(message, field...)
}

func Debug(message string, field ...zapcore.Field) {
	Log.Debug(message, field...)
}
func Error(message interface{}, field ...zapcore.Field) {
	switch v := message.(type) {
	case error:
		Log.Error(v.Error(), field...)
	case string:
		Log.Error(v, field...)
	}
}
