package xcodeproject

import (
	"os/exec"

	"github.com/AlTavares/go/sh"
)

func IsCarthage() bool {
	return sh.FileExists("Cartfile")
}

func IsCocoapods() bool {
	return sh.FileExists("Podfile")
}

func IsGitTreeClean() bool {
	cmd := exec.Command("git", "git diff-index --quiet HEAD --")
	return cmd.Run() == nil
}
