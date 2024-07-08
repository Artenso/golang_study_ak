package main

import (
	"fmt"
	"strings"
)

const (
	Read    = "Read"
	Write   = "Write"
	Execute = "Execute"
)

func main() {
	num := 755
	fmt.Println(getFilePermissions(num))

}

func getFilePermissions(flag int) string {
	owner := flag / 100
	group := (flag - owner*100) / 10
	other := flag % 10

	getPerm := func(num int) string {
		perm := []string{Read, Write, Execute}
		for i := 0; i < 3; i++ {
			if num>>(2-i)&1 > 0 {
				continue
			}
			perm[i] = "-"
		}
		return strings.Join(perm, ",")
	}

	return fmt.Sprintf("Owner:%s Group:%s Other:%s", getPerm(owner), getPerm(group), getPerm(other))
}
