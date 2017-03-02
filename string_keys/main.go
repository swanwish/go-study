package main

import (
	"flag"
	"io/ioutil"

	"strings"

	"fmt"

	"github.com/swanwish/go-common/logs"
	"github.com/swanwish/go-common/utils"
)

var (
	path string
)

func parseCmdLineArgs() {
	flag.StringVar(&path, "path", "", "The path to parse")
	flag.Parse()
}

type KeyPair struct {
	Define string
	Value  string
}

func main() {
	parseCmdLineArgs()
	if path == "" {
		logs.Errorf("The path is not specified")
		return
	}
	if !utils.FileExists(path) {
		logs.Errorf("The file %s does not exists", path)
		return
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		logs.Errorf("Failed to read file %s, the error is %v", path, err)
		return
	}
	lines := strings.Split(string(content), "\n")
	defPairs := []KeyPair{}
	maxLength := 0
	results := []string{}
	for _, line := range lines {
		line = strings.Trim(line, " \t")
		if line != "" {
			parts := strings.Split(line, "\t")
			if len(parts) != 2 {
				logs.Errorf("Invalid line %s, it should be separated by \\n", line)
				continue
			}
			defPairs = append(defPairs, KeyPair{Define: parts[0], Value: parts[1]})
			if len(parts[0]) > maxLength {
				maxLength = len(parts[0])
			}
		} else {
			if len(defPairs) > 0 {
				for _, pair := range defPairs {
					results = append(results, pair.GetDefinition(maxLength))
				}
				defPairs = []KeyPair{}
				maxLength = 0
				results = append(results, "")
			}
		}
	}
	if len(defPairs) > 0 {
		for _, pair := range defPairs {
			results = append(results, pair.GetDefinition(maxLength))
		}
		defPairs = []KeyPair{}
		maxLength = 0
	}

	fmt.Printf("%s\n", strings.Join(results, "\n"))
}

func (pair *KeyPair) GetDefinition(maxLength int) string {
	paddingLength := maxLength - len(pair.Define)
	padding := ""
	for i := 0; i < paddingLength; i++ {
		padding += " "
	}
	return fmt.Sprintf(`#define %s%s    @"%s"`, pair.Define, padding, pair.Value)
}
