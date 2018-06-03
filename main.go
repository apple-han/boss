package main

import (
	"learn/job/boss/parser"
	"learn/job/engine"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://www.zhipin.com/?ka=header-home-logo",
		ParserFunc: parser.PrintPositionList,
	})
}
