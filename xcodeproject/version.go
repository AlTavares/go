package xcodeproject

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/AlTavares/go/sh"
)

func SetVersion(version string) {
	sh.Run("agvtool new-marketing-version", version)
}

func SetBuild(version string) {
	sh.Run("agvtool new-version -all", version)
}

func IncrementBuildNumber() {
	sh.Run("agvtool next-version -all")
}

func UpdatePodspecVersion(path string, version string) {
	var re = regexp.MustCompile(`version = '.*'`)
	input, err := ioutil.ReadFile(path)
	sh.Check(err)
	newVersion := fmt.Sprintf("version = '%s'", version)
	output := re.ReplaceAllString(string(input), newVersion)
	err = ioutil.WriteFile(path, []byte(output), 0666)
	sh.Check(err)
}
