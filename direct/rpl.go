package main

import (
	"bufio"
	"fmt"
	"os"
)

func rpl() {
	rd := bufio.NewScanner(os.Stdin)
	for rd.Scan() {
		line := rd.Text()
		if len(line) > 0 && line[0] == 'E' {
			fmt.Fprintf(os.Stderr, "to stderr: %s\n", line[1:])
		} else if line == "." {
			return
		} else {
			fmt.Printf("to stdout: %s\n", line)
		}
	}
	if err := rd.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "err=%v\n", err)
	}
}
