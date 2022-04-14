package logger

import (
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var logger = logrus.New()
var Date = time.Now().Format("2006-01")
var OutFileError = getWriter("../logs/errors/ilanver-error(" + Date + ").log")
var OutFileWarning = getWriter("../logs/warnings/ilanver-warning(" + Date + ").log")
var OutFileInfo = getWriter("../logs/infos/ilanver-info(" + Date + ").log")

const (
	splitAfterPkgName = "github.com/barancanatbas"
)

func init() {
	//logger.Out = getWriter()
	logger.Level = logrus.InfoLevel
	logger.Formatter = &prefixed.TextFormatter{
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	}
	logger.SetReportCaller(false)
}

// SetLogLevel sets log level
func SetLogLevel(level logrus.Level) {
	logger.Level = level
}

// Fields sets fields on the logger.
type Fields logrus.Fields

// skipFrameCount 4 olursa kendisinin bir üstünde çağrılan metodu bulur
// skipFrameCount 5 olursa kendisinin çağıran metodu çağıran metodu bulur
// databaseerror kısmında 5 kullanılması gerekiyor. çünkü database error kısmını çağıran metot bizim için gerekli
// ama diğer yerlerde 4 verilecek

// Debugf logs a message at level Debug on the standard logger.
func Debugf(skipFrameCount int, format string, args ...interface{}) {
	if logger.Level >= logrus.DebugLevel {
		entry := newEntry(skipFrameCount)
		entry.Debugf(format, args...)
	}
}

// Infof logs a message at level Info on the standard logger.
func Infof(skipFrameCount int, format string, args ...interface{}) {
	logger.Out = OutFileInfo
	if logger.Level >= logrus.InfoLevel {
		entry := newEntry(skipFrameCount)
		entry.Infof(format, args...)
	}
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(skipFrameCount int, format string, args ...interface{}) {
	logger.Out = OutFileWarning
	if logger.Level >= logrus.WarnLevel {
		entry := newEntry(skipFrameCount)
		entry.Warnf(format, args...)
	}
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(skipFrameCount int, format string, args ...interface{}) {
	logger.Out = OutFileError
	if logger.Level >= logrus.ErrorLevel {
		entry := newEntry(skipFrameCount)
		entry.Errorf(format, args...)
	}
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(skipFrameCount int, format string, args ...interface{}) {
	if logger.Level >= logrus.FatalLevel {
		entry := newEntry(skipFrameCount)
		entry.Fatalf(format, args...)
	}
}

func newEntry(skipFrameCount int) *logrus.Entry {
	file, function, line := callerInfo(skipFrameCount, splitAfterPkgName)

	entry := logger.WithFields(logrus.Fields{})
	entry.Data["file"] = file
	entry.Data["line"] = line
	entry.Data["function"] = function
	return entry
}

// callerInfo grabs caller file, function and line number
func callerInfo(skip int, pkgName string) (file, function string, line int) {

	// Grab frame
	pc := make([]uintptr, 1)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	// Set file, function and line number
	file = trimPkgName(frame.File, pkgName)
	function = trimPkgName(frame.Function, pkgName)
	line = frame.Line

	return
}

// trimPkgName trims string after splitStr
func trimPkgName(frameStr, splitStr string) string {
	count := strings.LastIndex(frameStr, splitStr)
	if count > -1 {
		frameStr = frameStr[count+len(splitStr):]
	}

	return frameStr
}

// getWriter returns io.Writer
func getWriter(fileName string) io.Writer {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Errorf("Failed to open log file: %v", err)
		return os.Stdout
	} else {
		return file
	}
}

func a() {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	logrus.WithField("file", filename).WithField("function", fn)
}
