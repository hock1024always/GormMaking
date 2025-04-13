package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// 日志实例
var (
	// errorLog 定义了一个用于记录错误日志的日志记录器，日志前缀为红色的 [error]，并包含时间戳和文件行号
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	// infoLog 定义了一个用于记录信息日志的日志记录器，日志前缀为蓝色的 [info ]，并包含时间戳和文件行号
	infoLog = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	// loggers 是一个包含 errorLog 和 infoLog 的日志记录器切片
	loggers = []*log.Logger{errorLog, infoLog}
	// TODO 保证每次只有一个goroutine可以访问这个日志记录器（之前没考虑到）
	mu sync.Mutex
)

// 日志实例
var (
	// Error 使用 errorLog 记录错误信息
	Error = errorLog.Println
	// Errorf 使用 errorLog 记录格式化的错误信息
	Errorf = errorLog.Printf
	// Info 使用 infoLog 记录普通信息
	Info = infoLog.Println
	// Infof 使用 infoLog 记录格式化的普通信息
	Infof = infoLog.Printf
)

// 三个日志层级
const (
	InfoLevel = iota //有点6
	ErrorLevel
	Disabled
)

// SetLevel controls log level
// 设置日志级别
func SetLevel(level int) {

	mu.Lock()
	defer mu.Unlock()

	// 遍历所有日志记录器
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	// 如果错误级别小于传入的级别，则将错误日志输出设置为丢弃
	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}
	// 如果信息级别小于传入的级别，则将信息日志输出设置为丢弃
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}
