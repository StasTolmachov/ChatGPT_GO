package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
)

// MyFormatter кастомный форматтер для цветного вывода в терминале
type MyFormatter struct{}

// ANSI коды для цветного вывода в терминале
var (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
)

// Цвета для уровней логирования
var levelColors = map[logrus.Level]string{
	logrus.PanicLevel: Red,
	logrus.FatalLevel: Red,
	logrus.ErrorLevel: Red,
	logrus.WarnLevel:  Yellow,
	logrus.InfoLevel:  Green,
	logrus.DebugLevel: Blue,
	logrus.TraceLevel: Cyan,
}

// levelList для отображения уровня логирования в понятном формате
var levelList = []string{
	"PANIC", "FATAL", "ERROR", "WARN", "INFO", "DEBUG", "TRACE",
}

var Log *logrus.Logger

// Format метод для кастомного цветного форматирования логов
func (mf *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// Получаем название файла
	strList := strings.Split(entry.Caller.File, "/")
	fileName := strList[len(strList)-1]

	// Получаем имя функции
	pc := entry.Caller.PC
	funcName := runtime.FuncForPC(pc).Name()
	funcNameShort := funcName[strings.LastIndex(funcName, "/")+1:]

	// Цвет уровня
	levelColor := levelColors[entry.Level]
	level := levelList[int(entry.Level)]

	// Форматируем строку лога с цветами
	b.WriteString(fmt.Sprintf("%s[%s]%s | %s%s%s | %s:%d | %s",
		levelColor, level, Reset, Cyan, funcNameShort, Reset,
		fileName, entry.Caller.Line, entry.Message))

	// Только если есть данные в map, добавляем их
	if len(entry.Data) > 0 {
		b.WriteString(fmt.Sprintf(" | %s%v%s", Magenta, entry.Data, Reset))
	}

	b.WriteString("\n")
	return b.Bytes(), nil
}

// MakeLogger создает логгер с выводом только в терминал
func MakeLogger() {
	Log = logrus.New()

	// Устанавливаем вывод только в терминал
	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&MyFormatter{}) // Цветной формат для терминала

	Log.SetReportCaller(true)

}
