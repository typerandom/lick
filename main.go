package main

import (
	"github.com/codegangsta/cli"
	"os"
	"path/filepath"
	"strings"
)

func walkPrintLicense(license *LicenseNode, depth int) {
	print(strings.Repeat("\t", depth) + license.Name + "\n")

	depth += 1

	for _, subLicense := range license.Licenses {
		walkPrintLicense(subLicense, depth)
	}
}

func checkDirectoryPath(directoryPath string) {
	rootLicense, err := walkFolderWithLicense(directoryPath, nil)

	if err != nil {
		panic(err)
	}

	walkPrintLicense(rootLicense, 0)
}

func main() {
	app := cli.NewApp()

	app.Name = "check"
	app.Usage = "Check a directory and the licenses in it."
	app.Action = func(c *cli.Context) {
		directoryPath := "./"

		args := c.Args()

		if len(args) > 0 {
			directoryPath = args[0]
		}

		directoryPath, err := filepath.Abs(directoryPath)

		if err != nil {
			panic(err)
		}

		checkDirectoryPath(directoryPath + "/")
	}

	app.Run(os.Args)
}
