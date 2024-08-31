package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/ytnsym/receivernamelen"
)

func main() { unitchecker.Main(receivernamelen.Analyzer) }
