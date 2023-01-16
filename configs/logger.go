package configs

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ConfigLog struct {
	// Enable console logging
	ConsoleLoggingEnabled bool

	// EncodeLogsAsJson makes the log framework log JSON
	EncodeLogsAsJson bool
	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// Directory to log to to when filelogging is enabled
	Directory string
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge int
}

func newRollingFile() io.Writer {
	now := time.Now().UTC()
	date := now.Format("20060102")
	folder := EnvConfig()["LOGS_FOLDER"]
	maxsize := StringToInt(EnvConfig()["LOGS_ROTATION_MB"])
	maxage := StringToInt(EnvConfig()["LOGS_RETENTION_DAYS"])
	return &lumberjack.Logger{
		Filename:   path.Join("."+folder+"/", "compare-name_"+date+".log"),
		MaxBackups: 365,     // files
		MaxSize:    maxsize, // megabytes
		MaxAge:     maxage,  // days
	}
}

func ZeroLogger() zerolog.Logger {
	enablefiles := EnvConfig()["LOGS_FILES"]
	var writers []io.Writer
	writers = append(writers, zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf(" %-9s |", i)
		},
		FormatFieldName: func(i interface{}) string {
			return fmt.Sprintf("%s: ", i)
		},
		FormatFieldValue: func(i interface{}) string {
			str, _ := i.(string)
			removeStr2Backslash := strings.ReplaceAll(str, "\\", "")
			return fmt.Sprintf("%s", removeStr2Backslash)
		},
	})

	if enablefiles == "ENABLE" {
		// save to file
		writers = append(writers, newRollingFile())
	}

	zerolog.MessageFieldName = "logs_id"
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"

	mw := io.MultiWriter(writers...)

	log := zerolog.New(mw).With().Timestamp().Logger()

	return log
}

func OpenLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
