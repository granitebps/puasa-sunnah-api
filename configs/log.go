package configs

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/granitebps/puasa-sunnah-api/constants"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Log struct {
	Logger *logrus.Logger
}

func NewLog(logFolder string) *Log {
	return &Log{Logger: initLog(logFolder)}
}

type CustomTextFormatter struct {
}

func (f *CustomTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	lc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		lc = time.UTC
	}
	datetime := entry.Time.In(lc).Format(time.RFC3339)
	level := strings.ToUpper(entry.Level.String())
	msg := entry.Message
	file := entry.Caller.File
	line := entry.Caller.Line
	env := strings.ToUpper(viper.GetString("APP_ENV"))
	logMsg := fmt.Sprintf("[%s][%s:%s][%s:%d] : %s", datetime, env, level, file, line, msg)
	return append([]byte(logMsg), '\n'), nil
}

func initLog(logFolder string) *logrus.Logger {
	createLogFolder(logFolder)

	logger := logrus.New()
	logger.SetFormatter(new(CustomTextFormatter))
	logger.SetReportCaller(true)

	f, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logFolder, time.Now().Format(constants.LOG_FORMAT_DATE)), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panicf("error opening file: %v", err)
	}

	log.SetOutput(f)

	return logger
}

func createLogFolder(logFolder string) {
	if _, err := os.Stat(logFolder); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(logFolder, os.ModePerm)
		if err != nil {
			// log.Panicf("error creating log folder: %v", e.WrapError(err))
			log.Panicf("error creating log folder: %v", err)
		}
	}
}
