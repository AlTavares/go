package sh

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type CommandOutput struct {
	Output      string
	ErrorOutput string
	Error       error
}

func Command(command ...string) *exec.Cmd {
	return CommandWithOutput(os.Stdout, os.Stderr, command...)
}

func Run(command ...string) (output CommandOutput) {
	outWriter := NewProxyWriter(os.Stdout)
	errWriter := NewProxyWriter(os.Stderr)
	cmd := CommandWithOutput(outWriter, errWriter, command...)
	output.Error = cmd.Run()
	output.Output = outWriter.String()
	output.ErrorOutput = errWriter.String()
	return output
}

func CommandWithOutput(stdout, stderr io.Writer, command ...string) *exec.Cmd {
	c := strings.Join(command, " ")
	fmt.Println(c)
	cmd := exec.Command("sh", "-c", c)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	return cmd
}

func RunAt(path string, command ...string) {
	cmd := append([]string{"cd", path, "&&"}, command...)
	Run(cmd...)
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
