package main

import "os"

func getArgs() []string {
	if len(os.Args) < 2 {
		return nil
	}
	return os.Args[1:]
}
