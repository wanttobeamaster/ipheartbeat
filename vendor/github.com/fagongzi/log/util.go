package log

import (
	"flag"
)

var (
	crashLog    = flag.String("log-crash", "/tmp/crash.log", "The crash log file.")
	logFile     = flag.String("log-file", "/tmp/run-log", "The external log file. Default log to console.")
	logLevel    = flag.String("log-level", "info", "The log level, default is info")
	logRotateBy = flag.String("log-rotate-by", "day", "The log rotate by [day|hour], default is day")
	logCount    = flag.Int("log-count", 10, "Count of log files, default is 10")
	logHigh     = flag.Bool("log-high", false, "The log highlighting")
)

// Cfg is the log cfg
type Cfg struct {
	LogLevel string
	LogFile  string
}

// InitLog init log
func InitLog() {
	if !flag.Parsed() {
		flag.Parse()
	}

	SetLogCount(*logCount)
	SetHighlighting(*logHigh)
	SetLevelByString(*logLevel)
	if "" != *logFile {
		if *logRotateBy == "hour" {
			SetRotateByHour()
		} else {
			SetRotateByDay()
		}
		SetOutputByName(*logFile)
		CrashLog(*crashLog)
	}

	if !DebugEnabled() {
		SetFlags(Ldate | Ltime | Lmicroseconds)
	}
}
