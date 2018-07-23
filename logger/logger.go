package logger

import (
	"fmt"

	"github.com/fatih/color"
)

const arrow = "➜ "

func Log(a ...interface{}) {
	LogColor(color.FgHiCyan, a...)
}

func Success(a ...interface{}) {
	LogColor(color.FgHiGreen, a...)
}

func LogColor(c color.Attribute, a ...interface{}) {
	color.Set(c)
	msg := append([]interface{}{arrow}, a...)
	fmt.Println(msg...)
	color.Unset()
}

func Error(e error) {
	LogColor(color.FgHiRed, e.Error())
}

func Fatal(e error) {
	Error(e)
	panic("exiting...")
}
