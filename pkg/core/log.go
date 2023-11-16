package core

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LogrusTextFormatter struct {
}

func SetupLog() *logrus.Logger {
	createLogFolder(constants.LOG_FOLDER)

	return setupLogrus()
}

//lint:ignore U1000 You can use github.com/rs/zerolog for logging
func setupZerolog() zerolog.Logger {
	consoleWriter := zerolog.ConsoleWriter{
		NoColor: true,
		Out:     createLogWriter(),
		FormatTimestamp: func(i interface{}) string {
			return fmt.Sprintf("[%s]", i)
		},
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s:%s]", viper.GetString(constants.APP_ENV), i))
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf(": %s", i)
		},
		FormatCaller: func(i interface{}) string {
			fmt.Println(i)
			return fmt.Sprintf("[%s]", i)
		},
	}
	logger := zerolog.New(consoleWriter).Level(zerolog.TraceLevel).With().Timestamp().Caller().Stack().Logger()
	return logger
}

func setupLogrus() *logrus.Logger {
	l := logrus.New()

	l.SetReportCaller(true)
	l.SetFormatter(new(LogrusTextFormatter))
	l.SetOutput(createLogWriter())

	return l
}

func createLogWriter() *os.File {
	//#nosec G302 -- File permission need to be 0666
	logFile, err := os.OpenFile(
		fmt.Sprintf("%s/%s.log", constants.LOG_FOLDER, time.Now().Format(constants.FORMAT_DATE)),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0o666,
	)

	if err != nil {
		log.Panic(err)
	}

	return logFile
}

func createLogFolder(logFolder string) {
	if _, err := os.Stat(logFolder); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(logFolder, os.ModePerm)
		if err != nil {
			log.Panic(err)
		}
	}
}

func (f *LogrusTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	lc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		lc = time.UTC
	}
	datetime := entry.Time.In(lc).Format(time.RFC3339)
	level := strings.ToUpper(entry.Level.String())
	msg := entry.Message
	file := entry.Caller.File
	line := entry.Caller.Line
	env := strings.ToUpper(viper.GetString(constants.APP_ENV))
	logMsg := fmt.Sprintf("[%s][%s:%s][%s:%d] : %s", datetime, env, level, file, line, msg)
	return append([]byte(logMsg), '\n'), nil
}
