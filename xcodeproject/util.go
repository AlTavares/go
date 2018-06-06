package xcodeproject

import (
	"os"
	"os/exec"
)

func IsCarthage() bool {
	return FileExists("Cartfile")
}

func IsCocoapods() bool {
	return FileExists("Podfile")
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func IsGitTreeClean() bool {
	cmd := exec.Command("git", "diff --quiet")
	return cmd.Run() == nil
}
