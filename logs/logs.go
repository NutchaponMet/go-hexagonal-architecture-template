package logs

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogEntry struct {
	Level     string `bson:"level"`
	Message   string `bson:"message"`
	TimeStamp string `bson:"timestamp"`
}

var log *zap.Logger
var cli *mongo.Client
var collection *mongo.Collection

func LogInit() {

	cli = initLogMongoClient()
	collection = cli.Database("api_dev_db").Collection("kpi_clik_logs")

	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

}

func Info(message string, field ...zapcore.Field) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, _ = collection.InsertOne(ctx, LogEntry{
		Level:     zapcore.InfoLevel.String(),
		Message:   message,
		TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
	})

	log.Info(message, field...)
}

func Debug(message string, field ...zapcore.Field) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, _ = collection.InsertOne(ctx, LogEntry{
		Level:     zapcore.InfoLevel.String(),
		Message:   message,
		TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
	})

	log.Debug(message, field...)
}
func Error(message interface{}, field ...zapcore.Field) {
	switch v := message.(type) {
	case error:
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, _ = collection.InsertOne(ctx, LogEntry{
			Level:     zapcore.InfoLevel.String(),
			Message:   v.Error(),
			TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
		})

		log.Error(v.Error(), field...)
	case string:
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, _ = collection.InsertOne(ctx, LogEntry{
			Level:     zapcore.InfoLevel.String(),
			Message:   v,
			TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
		})

		log.Error(v, field...)
	}
}

func initLogMongoClient() (client *mongo.Client) {
	DBURL := fmt.Sprintf("%v://%v:%v@%v:%v/%v",
		viper.GetString("mongodb.driver"),
		viper.GetString("mongodb.username"),
		viper.GetString("mongodb.password"),
		viper.GetString("mongodb.host"),
		viper.GetInt("mongodb.port"),
		viper.GetString("mongodb.dbname"),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DBURL))
	if err != nil {
		panic(err)
	}
	return
}
