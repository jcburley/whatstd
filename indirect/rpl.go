package main

import (
	"bufio"
	"fmt"
)

func rpl() {
	rd := bufio.NewScanner(Stdin)
	for rd.Scan() {
		line := rd.Text()
		if len(line) > 0 && line[0] == 'E' {
			fmt.Fprintf(Stderr, "to stderr: %s\n", line[1:])
		} else if line == "." {
			return
		} else {
			fmt.Fprintf(Stdout, "to stdout: %s\n", line)
		}
	}
	if err := rd.Err(); err != nil {
		fmt.Fprintf(Stderr, "err=%v\n", err)
	}
}
