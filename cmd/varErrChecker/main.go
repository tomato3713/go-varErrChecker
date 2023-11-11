package main

import (
	"github.com/tomato3713/varErrChecker"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(varErrChecker.Analyzer)
}
