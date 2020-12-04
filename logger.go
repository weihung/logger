package logger

import (
	"log"
	"os"
	"time"
)

type Logger struct {
	Date       string
	Path       string
	FileName   string
	NameFormat string
	File       *os.File
	logger     *log.Logger
	Prefix     string
	ShowLog    bool
}

func New(path, filename, prefix string, showLog ...bool) *Logger {
	show := false
	if len(showLog) > 0 {
		show = showLog[0]
	}
	logger := Logger{Path: path, FileName: filename, Prefix: prefix, ShowLog: show, NameFormat: "2006-01-02"}
	return &logger
}

func (l *Logger) CheckDate() {
	now := time.Now()
	date := now.Format(l.NameFormat)
	if date != l.Date {
		if l.File != nil {
			l.File.Close()
		}
		l.Date = date
		filename := l.Path + "/" + l.FileName + "-" + date + ".log"
		logFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}

		l.logger = log.New(logFile, l.Prefix, log.LstdFlags)
	}
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.CheckDate()
	l.logger.Printf(format, v...)
	if l.ShowLog {
		log.Printf(format, v...)
	}
}

func (l *Logger) Print(v ...interface{}) {
	l.CheckDate()
	l.logger.Print(v...)
	if l.ShowLog {
		log.Print(v...)
	}
}

func (l *Logger) Println(v ...interface{}) {
	l.CheckDate()
	l.logger.Println(v...)
	if l.ShowLog {
		log.Println(v...)
	}
}

func (l *Logger) Fatal(v ...interface{}) {
	l.CheckDate()
	l.logger.Fatal(v...)
	if l.ShowLog {
		log.Fatal(v...)
	}
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.CheckDate()
	l.logger.Fatalf(format, v...)
	if l.ShowLog {
		log.Fatalf(format, v...)
	}
}

func (l *Logger) Fatalln(v ...interface{}) {
	l.CheckDate()
	l.logger.Fatalln(v...)
	if l.ShowLog {
		log.Fatalln(v...)
	}
}

// Panic is equivalent to Print() followed by a call to panic().
func (l *Logger) Panic(v ...interface{}) {
	l.CheckDate()
	l.logger.Print(v...)
	if l.ShowLog {
		log.Panic(v...)
	}
	os.Exit(1)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.CheckDate()
	l.logger.Printf(format, v...)
	if l.ShowLog {
		log.Panicf(format, v...)
	}
	os.Exit(1)
}

// Panicln is equivalent to Println() followed by a call to panic().
func (l *Logger) Panicln(v ...interface{}) {
	l.CheckDate()
	l.logger.Println(v...)
	if l.ShowLog {
		log.Println(v...)
	}
}

// E is log with error tag
func (l *Logger) E(v ...interface{}) {
	l.CheckDate()
	label := []interface{}{"[ERROR]"}
	v = append(label[:1], v...)
	l.logger.Println(v...)
	if l.ShowLog {
		log.Println(v...)
	}
}

// W is log with warring tag
func (l *Logger) W(v ...interface{}) {
	l.CheckDate()
	label := []interface{}{"[WARRING]"}
	v = append(label[:1], v...)
	l.logger.Println(v...)
	if l.ShowLog {
		log.Println(v...)
	}
}

// I is log with info tag
func (l *Logger) I(v ...interface{}) {
	l.CheckDate()
	label := []interface{}{"[INFO]"}
	v = append(label[:1], v...)
	l.logger.Println(v...)
	if l.ShowLog {
		log.Println(v...)
	}
}

// D is log with debig tag
func (l *Logger) D(v ...interface{}) {
	l.CheckDate()
	label := []interface{}{"[DEBUG]"}
	v = append(label[:1], v...)
	l.logger.Println(v...)
	if l.ShowLog {
		log.Println(v...)
	}
}
