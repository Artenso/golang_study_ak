package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// printTree prints dir contents
func printTree(path string, prefix string, isLast bool, depth int) {
	// stop printing
	if isLast || depth <= 0 {
		return
	}

	// get files and folders (entries) from current dir
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("failed to read dir: %s", err.Error())
	}

	// get file or folder (f) form files and folders
	for i, entry := range entries {
		fmt.Print(prefix)

		// create new prefix for next call
		newPrefix := prefix

		// update new prefix depends on is element last in folder
		if i == len(entries)-1 {
			fmt.Print("└── ")
			newPrefix += "    "
		} else {
			fmt.Print("│── ")
			newPrefix += "│   "
		}

		fmt.Println(entry.Name())

		// create new path to look deeper
		newPath := filepath.Join(path, entry.Name())

		// stop printing than entry is not folder and all files are printed
		if !entry.IsDir() {
			isLast = true
		}

		// call printTree with new args
		printTree(newPath, newPrefix, isLast, depth-1)

	}
}

func main() {
	var depth int
	// get depth flag
	flag.IntVar(&depth, "n", -1, "tree depth")
	flag.Parse()

	// check path specified
	if flag.NArg() == 0 {
		log.Fatal("no path to directory specified")
	}

	// get path
	path := flag.Arg(0)

	// convert relative path to absolute
	if !strings.HasPrefix(path, "/") {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		path = filepath.Join(wd, path)
	}

	printTree(path, "", false, depth)
}

/*
EXAMPLE n = 5:
go run main.go -n 5 course1/16.terminal_app
│── 1.terminal_app_args
│   └── task1.16.1.1
│       │── main.go
│       └── task1.16.1.1
└── 2.terminal_app_flag
    └── task1.16.2.1
        │── main.go
        └── task1.16.2.1.

EXAMPLE n = 5:
go run main.go -n 2 course1/16.terminal_app
│── 1.terminal_app_args
│   └── task1.16.1.1
└── 2.terminal_app_flag
    └── task1.16.2.1
*/
