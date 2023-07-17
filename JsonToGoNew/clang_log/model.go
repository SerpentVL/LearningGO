package clang_log

import "fmt"

func Info(msg string, fields ...interface{}) {
	fmt.Printf("Info: " + msg + "\n", fields...)
}

func Debug(msg string, fields ...interface{}) {
	fmt.Printf("Debug: " + msg + "\n", fields...)
}

func Error(msg string, fields ...interface{}) {
	fmt.Printf("Error: " + msg + "\n", fields...)
}

func Clang(msg string) {
	fmt.Printf("Clang out: %s\n", msg)
}
