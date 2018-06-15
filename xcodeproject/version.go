package xcodeproject

import (
	"io/ioutil"
	"regexp"
	"strings"

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
	var re = regexp.MustCompile(`version\s*=\s*("|')(.*)("|')`)
	input, err := ioutil.ReadFile(path)
	sh.Check(err)
	currentVersion := re.FindAllStringSubmatch(string(input), -1)[0][2]
	output := strings.Replace(string(input), currentVersion, version, -1)
	err = ioutil.WriteFile(path, []byte(output), 0666)
	sh.Check(err)
}
