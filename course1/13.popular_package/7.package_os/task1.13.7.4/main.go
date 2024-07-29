package main

import (
	"fmt"
	"os/exec"
)

func ExecBin(binPath string, args ...string) string {
	cmd := exec.Command(binPath, args...)
	data, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error executing binary: %s", err.Error())
	}
	return string(data)
}

func main() {
	fmt.Println(ExecBin("ls", "-la"))
	fmt.Println(ExecBin("nonexistent-binary"))
}
