package main

import "github.com/swanwish/go-common/logs"

type commonGenerator struct {
}

func (g commonGenerator) GenerateFile() {
	g.GetText()
}

func (g commonGenerator) GetText() {
	logs.Debugf("get text")
}

type V2Generator struct {
	commonGenerator
}

func (g V2Generator) GetText() {
	logs.Debugf("get v2 text")
}

type V3Generator struct {
	commonGenerator
}

func (g V3Generator) GetText() {
	logs.Debugf("get v3 text")
}

func main() {
	generator := V2Generator{}
	generator.GenerateFile()
}
