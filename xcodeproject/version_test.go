package xcodeproject

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const filename = "regextest"

func TestUpdatePodspecVersion(t *testing.T) {
	testData := map[string]string{
		`  s.version          = "1.3.0"`: `  s.version          = "6.6.6"`,
		`  s.version          = '1.3.0'`: `  s.version          = '6.6.6'`,
		`s.version = "v1.3.0"`:           `s.version = "6.6.6"`,
		`s.version = 		"something"`: `s.version = 		"6.6.6"`,
		`s.version = 'whatever'`: `s.version = '6.6.6'`,
	}

	for original, expected := range testData {
		data := []byte(original)
		ioutil.WriteFile(filename, data, 0666)
		UpdatePodspecVersion(filename, "6.6.6")
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Error("Error not nil:", err)
		}
		gotten := string(file)
		assertEqual(t, gotten, expected)
	}
	os.Remove(filename)
}

func assertEqual(t *testing.T, gotten string, expected string) {
	if gotten != expected {
		fmt.Println()
		fmt.Println("Expected:", expected)
		fmt.Println("Gotten:  ", gotten)
		fmt.Println()
		t.Fail()
	}
}
