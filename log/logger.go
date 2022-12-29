package log

import (
	"log"
	"os"
)

var (
	WarnLog  *log.Logger
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

func init() {
	p, err := logFilePathSelection()
	if err != nil {
		log.Println(err.Error())
	}

	file, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}

	format := log.Ldate | log.Ltime | log.Lshortfile

	WarnLog = log.New(file, "Warning: ", format)
	InfoLog = log.New(file, "Info: ", format)
	ErrorLog = log.New(file, "Error: ", format)
}
