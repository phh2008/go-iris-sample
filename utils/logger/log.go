package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
)

//日志输出类型
type OutType string

const (
	//文件
	File OutType = "file"
	//控制台
	Console OutType = "console"
)

//日志级别
type Level uint8

const (
	// DebugLevel defines debug log level.
	DebugLevel Level = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel
	// NoLevel defines an absent log level.
	NoLevel
	// Disabled disables the logger.
	Disabled
)

type config struct {
	//输出类型：file,console
	outType OutType
	//时间格式
	timeFormat string
	//日志极别
	level Level
	//字段
	field map[string]interface{}
}

//default config
func NewConfig() *config {
	return &config{
		outType:    Console,
		timeFormat: "2006-01-02 15:04:05.000",
		level:      InfoLevel,
	}
}

//设置输出类型
func (a *config) OutType(outType OutType) *config {
	a.outType = outType
	return a
}

//设置时间格式
func (a *config) TimeFormat(format string) *config {
	a.timeFormat = format
	return a
}

//设置日志极别
func (a *config) Level(level Level) *config {
	a.level = level
	return a
}

//字段
func (a *config) Field(field map[string]interface{}) *config {
	a.field = field
	return a
}

//初始化日志配置
func (a *config) Build() {
	zeroLogInit(*a)
}

func zeroLogInit(cfg config) {
	//时间格式
	zerolog.TimeFieldFormat = cfg.timeFormat
	//全局level
	lv := uint8(cfg.level)
	zerolog.SetGlobalLevel(zerolog.Level(lv))
	var logger zerolog.Logger
	if cfg.outType == File {
		//日志文件输出
		lumLog := &lumberjack.Logger{
			Filename:   "/var/log/myBlog/log.log",
			MaxSize:    5, // megabytes
			MaxBackups: 3,
			MaxAge:     31,    //days
			Compress:   false, // disabled by default
		}
		logger = zerolog.New(lumLog)
	} else {
		//控制台输出
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: cfg.timeFormat}
		output.FormatTimestamp = func(i interface{}) string {
			return fmt.Sprintf("[%s]", i)
		}
		output.FormatLevel = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
		}
		output.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		}
		output.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("%s:", i)
		}
		output.FormatFieldValue = func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		}
		logger = zerolog.New(output)
	}
	logger = logger.With().Fields(cfg.field).Timestamp().Logger()
	log.Logger = logger
}
