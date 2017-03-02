package main

import (
	"flag"

	"github.com/swanwish/go-common/logs"
	"github.com/swanwish/go-study/epub/epub_file"
)

var (
	name    string
	srcDir  string
	destDir string
	key     string
)

func parseCmdLineArgs() {
	flag.StringVar(&name, "name", "", "The name of the epub file")
	flag.StringVar(&srcDir, "src", "", `The src directory of the epub file`)
	flag.StringVar(&destDir, "dest", "", "The destination directory for the generated epub file")
	flag.StringVar(&key, "key", "", "The key for encryption content")
	flag.Parse()
}

func main() {
	parseCmdLineArgs()
	if name == "" {
		logs.Errorf("The epub file name is not specified")
		return
	}
	if srcDir == "" {
		logs.Errorf("The source directory is not specified")
		return
	}
	if destDir == "" {
		destDir = srcDir
	}
	ef := epub_file.EPubFile{}
	if err := ef.Create(name, key, srcDir, destDir); err != nil {
		logs.Errorf("Failed to create epub file generator, the error is %v", err)
		return
	}
	if err := ef.Generate(); err != nil {
		logs.Errorf("Failed to generate epub file, the error is %v", err)
		return
	}
}
