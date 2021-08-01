package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger // 几乎任何东西
	Info    *log.Logger // 重要信息
	Warning *log.Logger // 警告
	Error   *log.Logger // 错误
)

// 配置日志
// func init() {
// 	// 日志前缀
// 	log.SetPrefix("Away0x")
// 	// 配置输出
// 	// log.SetOutput(os.Stderr)
// 	f, _ := os.OpenFile("./demo.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
// 	log.SetOutput(f)
// 	// 配置输出格式
// 	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
// }

// 自定义 log
func init() {
	logFlags := log.Ldate | log.Ltime | log.Lshortfile
	file, err := os.OpenFile("customlog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("无法打开 log 文件: ", err)
	}

	Trace = log.New(ioutil.Discard, "Trace: ", logFlags)
	Info = log.New(os.Stdout, "Info: ", logFlags)
	Warning = log.New(os.Stdout, "Trace: ", logFlags)
	// MultiWriter: 将两个 writer 组合到一起，这样可使日志可以同时记录到终端和文件中
	Error = log.New(io.MultiWriter(file, os.Stderr), "Trace: ", logFlags)
}

func main() {
	// log.Println("123") // 记录日志
	// log.Fatalln("") // 记录日志并调用 os.Exit(1) 退出程序
	// log.Panicln("") // 记录日志并 panic

	Trace.Println("一些鸡毛蒜皮的小事")
	Info.Println("一些特别的信息")
	Warning.Println("这是一个警告")
	Error.Println("出现了故障")

}
