package logger

import (
	"fmt"

	"github.com/fatih/color"
)

const arrow = "➜ "

func Log(a ...interface{}) {
	LogColor(color.FgHiCyan, a...)
}

func LogColor(c color.Attribute, a ...interface{}) {
	color.Set(c)
	msg := append([]interface{}{arrow}, a...)
	fmt.Println(msg...)
	color.Unset()
}

func Error(e error) {
	color.Set(color.FgHiRed)
	fmt.Print(arrow, " ")
	panic(color.HiRedString(e.Error()))
}
