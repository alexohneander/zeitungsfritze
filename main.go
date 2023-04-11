package main

import (
	"github.com/alexohneander/zeitungsfritze/cmd"
)

var (
	version    = "unknown"
	commitHash = "unknown"
)

func main() {
	cmd.Run(version, commitHash)
}
