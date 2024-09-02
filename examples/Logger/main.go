package main

////`k8s.io/klog/v2` 是 Kubernetes 使用的日志库，提供了多种日志记录方法。以下是一些常用的方法：
//
//### 常用方法
//
//#### Info
//```go
//log.Info(args ...interface{})
//log.Infof(format string, args ...interface{})
//log.Infoln(args ...interface{})
//```
//
//#### Warning
//```go
//log.Warning(args ...interface{})
//log.Warningf(format string, args ...interface{})
//log.Warningln(args ...interface{})
//```
//
//#### Error
//```go
//log.Error(args ...interface{})
//log.Errorf(format string, args ...interface{})
//log.Errorln(args ...interface{})
//```
//
//#### Fatal
//```go
//log.Fatal(args ...interface{})
//log.Fatalf(format string, args ...interface{})
//log.Fatalln(args ...interface{})
//```
//
//#### V (for verbosity)
//### 示例
//
//以下是一个简单的示例，展示了如何使用 `log.V` 方法来记录不同级别的日志消息：
//
//```go
//package main
//
//import (
//	"flag"
//	log "k8s.io/klog/v2"
//)
//
//func init() {
//	// 初始化 klog
//	flag.Set("logtostderr", "true") // 日志输出到标准错误
//	flag.Set("v", "2")              // 设置日志级别
//	flag.Parse()
//}
//
//func main() {
//	log.V(1).Info("This is a level 1 info message")
//	log.V(2).Infof("This is a level 2 info message with format: %s", "formatted string")
//	log.V(3).Infoln("This is a level 3 info message with newline")
//
//	// 确保所有日志都被刷新
//	log.Flush()
//}
//```
//
//在这个示例中，`log.V(level)` 用于记录不同级别的日志消息。通过设置日志级别（如 `flag.Set("v", "2")`），可以控制哪些级别的日志消息会被输出。
//```go
//log.V(level klog.Level).Info(args ...interface{})
//log.V(level klog.Level).Infof(format string, args ...interface{})
//log.V(level klog.Level).Infoln(args ...interface{})
//```
//
//#### Flush
//```go
//log.Flush()
//```
//
//### 示例代码
//```go
//package main
//
//import (
//	log "k8s.io/klog/v2"
//)
//
//func main() {
//	log.Info("This is an info message")
//	log.Infof("This is an info message with format: %s", "formatted string")
//	log.Warning("This is a warning message")
//	log.Errorf("This is an error message with format: %s", "formatted string")
//	log.Fatal("This is a fatal message")
//}
//```
//
//### 初始化和配置
//```go
//package main
//
//import (
//	"flag"
//	log "k8s.io/klog/v2"
//)
//
//func init() {
//	// 初始化 klog
//	flag.Set("logtostderr", "true") // 日志输出到标准错误
//	flag.Parse()
//}
//
//func main() {
//	log.Info("Klog initialized")
//	log.Flush() // 确保所有日志都被刷新
//}
//```

import (
	log "k8s.io/klog/v2"
)

func main() {
	log.Infof("Hello, World! %s", "klog")
}

////`log.Flush()` 是 `k8s.io/klog/v2` 库中的一个方法，用于确保所有缓冲的日志条目都被写入到输出目标（如文件或标准输出）。在程序结束前调用 `log.Flush()` 可以确保所有日志都被正确输出。
//
//### 示例
//
//以下是一个简单的示例，展示了如何使用 `log.Flush()`：
//
//```go
//package main
//
//import (
//	"flag"
//	log "k8s.io/klog/v2"
//)
//
//func init() {
//	// 初始化 klog
//	flag.Set("logtostderr", "true") // 日志输出到标准错误
//	flag.Parse()
//}
//
//func main() {
//	log.Info("This is an info message")
//	log.Warning("This is a warning message")
//	log.Error("This is an error message")
//
//	// 确保所有日志都被刷新
//	log.Flush()
//}
//```
//
//在这个示例中，`log.Flush()` 确保在程序结束前，所有日志消息都被写入到标准错误输出。
