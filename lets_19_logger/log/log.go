package log

import (
	"os"
)

const (
	INFO  = "[INFO]:"
	ERROR = "[Error]:"
	DeBug = "[DeBug]:"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
	DeBug(msg string)
}

//file文件输出log
type Flog struct {
	FilePath string //log file path
}

func (f Flog) Info(msg string) {
	//写入日志 info模式
	out(f, msg, INFO)
}

func (f Flog) Error(msg string) {
	out(f, msg, ERROR)
}

func (f Flog) DeBug(msg string) {
	out(f, msg, DeBug)
}

//console控制台log
type Clog struct {
}

func (c Clog) Info(msg string) {
	os.Stdout.WriteString(INFO + msg + "\n")
}
func (c Clog) Error(msg string) {
	os.Stdout.WriteString(ERROR + msg + "\n")
}
func (c Clog) DeBug(msg string) {
	os.Stdout.WriteString(DeBug + msg + "\n")
}

func NewFlog(file string) Logger {
	f := &Flog{file}
	logger := Logger(f)
	return logger
}
func NewClog() Logger {
	logger := Logger(&Clog{})
	return logger
}

/*
	打开目标文件
*/
func OpenWriter(f Flog) *os.File {
	file, err := os.OpenFile(f.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("日志输出错误...")
	}
	return file
}

//操作输出内容
func out(f Flog, msg string, model string) {
	file := OpenWriter(f)
	file.WriteString(model + msg + "\n")
	file.Close()
}
