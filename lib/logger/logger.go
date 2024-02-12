package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

type Logger struct {
	Logger     *log.Logger
	FileDir    string
	FileYear   int
	FileMonth  int
	FileDay    int
	FileIndex  int
	IsDebug    bool
	Decorators []string
}

func NewLogger(logDir string, debug bool) *Logger {
	today := time.Now()
	year, month, day := today.Date()
	index := 0

	logFile := path.Join(logDir, fmt.Sprintf("%d-%02d-%02d_%02d.log", year, month, day, index))

	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error opening log file:", err)
		return nil
	}

	for fileTooBig(&Logger{Logger: log.New(f, "", log.Ldate|log.Ltime)}) {
		index++
		logFile = path.Join(logDir, fmt.Sprintf("%d-%02d-%02d_%02d.log", year, month, day, index))
		f, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Println("Error opening log file:", err)
			return nil
		}
	}

	return &Logger{
		Logger:     log.New(f, "", log.Ldate|log.Ltime),
		FileDir:    logDir,
		FileYear:   year,
		FileMonth:  int(month),
		FileDay:    day,
		FileIndex:  index,
		IsDebug:    debug,
		Decorators: []string{},
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.checkFile()
	v = append([]interface{}{"INFO"}, v...)
	for _, d := range l.Decorators {
		v = append([]interface{}{d}, v...)
	}
	l.Logger.Println(v...)
}

func (l *Logger) Debug(v ...interface{}) {
	if l.IsDebug {
		l.checkFile()
		v = append([]interface{}{"DEBUG"}, v...)
		for _, d := range l.Decorators {
			v = append([]interface{}{d}, v...)
		}
		l.Logger.Println(v...)
	}
}

func (l *Logger) Warning(v ...interface{}) {
	l.checkFile()
	v = append([]interface{}{"WARNING"}, v...)
	for _, d := range l.Decorators {
		v = append([]interface{}{d}, v...)
	}
	l.Logger.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.checkFile()
	v = append([]interface{}{"ERROR"}, v...)
	for _, d := range l.Decorators {
		v = append([]interface{}{d}, v...)
	}
	l.Logger.Println(v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.checkFile()
	v = append([]interface{}{"FATAL"}, v...)
	for _, d := range l.Decorators {
		v = append([]interface{}{d}, v...)
	}
	l.Logger.Fatalln(v...)
}

func (l *Logger) AddDecorator(decorator string) {
	l.Decorators = append(l.Decorators, decorator)
}

func (l *Logger) RemoveDecorator(decorator string) {
	for i, d := range l.Decorators {
		if d == decorator {
			l.Decorators = append(l.Decorators[:i], l.Decorators[i+1:]...)
			break
		}
	}
}

func (l *Logger) checkFile() {
	if dateChanged(l) {
		l.updateLogDirectory()
	}

	if fileTooBig(l) {
		l.updateFileIndex()
	}
}

func dateChanged(l *Logger) bool {
	today := time.Now()
	year, month, day := today.Date()
	if year != l.FileYear || int(month) != l.FileMonth || day != l.FileDay {
		return true
	}
	return false
}

func (l *Logger) updateLogDirectory() {
	today := time.Now()
	year, month, day := today.Date()
	l.FileYear = year
	l.FileMonth = int(month)
	l.FileDay = day
	l.FileIndex = 0
	l.FileDir = path.Join(l.FileDir, fmt.Sprintf("%d-%02d-%02d", year, month, day))
	os.MkdirAll(l.FileDir, os.ModePerm)

	logFile := path.Join(l.FileDir, fmt.Sprintf("%d-%02d-%02d_%02d.log", year, month, day, l.FileIndex))
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error opening new log file:", err)
		return
	}

	l.Logger = log.New(f, "", log.Ldate|log.Ltime)
}

func fileTooBig(l *Logger) bool {
	fileInfo, err := os.Stat(l.Logger.Writer().(*os.File).Name())
	if err != nil {
		return false
	}
	if fileInfo.Size() > 10*1024*1024 { // 10MB
		return true
	}
	return false
}

func (l *Logger) updateFileIndex() {
	l.FileIndex++
	logFile := path.Join(l.FileDir, fmt.Sprintf("%d-%02d-%02d_%02d.log", l.FileYear, l.FileMonth, l.FileDay, l.FileIndex))
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error opening new log file:", err)
		return
	}

	l.Logger = log.New(f, "", log.Ldate|log.Ltime)
}
