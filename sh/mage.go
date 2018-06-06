package sh

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlTavares/go/logger"

	"github.com/fatih/color"
)

var DryRun = false

func Run(command ...string) {
	c := strings.Join(command, " ")
	logger.LogColor(color.FgHiGreen, c)
	if IsDryRun() {
		return
	}
	cmd := exec.Command("sh", "-c", c)
	cmd.Stdout = os.Stdout
	var errorBuffer bytes.Buffer
	cmd.Stderr = &errorBuffer
	err := cmd.Run()
	if err != nil {
		fmt.Println()
		logger.LogColor(color.FgHiRed, "Error running the following command:")
		logger.LogColor(color.FgHiGreen, "\t", c)
		rawError := strings.TrimSpace(errorBuffer.String())
		logger.LogColor(color.FgHiRed, rawError)
		Check(err)
	}
}

func RunAt(path string, command ...string) {
	cmd := append([]string{"cd", path, "&&"}, command...)
	Run(cmd...)
}

func Check(e error) {
	if e != nil {
		logger.Error(e)
	}
}

func IsDryRun() bool {
	return DryRun || strings.EqualFold(os.Getenv("dryrun"), "true")
}
