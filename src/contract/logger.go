package contract

type Logger interface {
	Debug(args ...interface{})
	Debugln(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infoln(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnln(args ...interface{})
	Warnf(format string, args ...interface{})

	Warning(args ...interface{})
	Warningln(args ...interface{})
	Warningf(format string, args ...interface{})

	Print(args ...interface{})
	Println(args ...interface{})
	Printf(format string, args ...interface{})

	Error(args ...interface{})
	Errorln(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalf(format string, args ...interface{})

	Panic(args ...interface{})
	Panicln(args ...interface{})
	Panicf(format string, args ...interface{})
}
