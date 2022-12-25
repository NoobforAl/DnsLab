package log

import (
	"log"
)

var (
	WarnLog  *log.Logger
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

func init() {
	format := log.Ldate | log.Ltime | log.Lshortfile

	WarnLog = log.New(nil, "Warning: ", format)
	InfoLog = log.New(nil, "Info: ", format)
	ErrorLog = log.New(nil, "Error: ", format)
}
