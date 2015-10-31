package version

import (
	"fmt"
	"os"
	"strings"
)

var (
	Version   string
	BuildTime string
	GitCommit string
)

func Print() {
	fmt.Println("Application version:", Version)
	fmt.Println("UTC Build Time:", BuildTime)
	fmt.Println("Git Commit Hash:", GitCommit)
	os.Exit(0)
}

func String() string {
	s := []string{Version, "\nUTC Build Time:", BuildTime, "\nGit Commit Hash:", GitCommit}
	return strings.Join(s, " ")
}
