package main

import (
	"bufio"
	"fmt"
)

func readLine(r *bufio.Reader) (s string, err error) {
	for {
		line, isPrefix, err := r.ReadLine();
		if err != nil {
			return "", err
		}
		s += string(line)
		if !isPrefix {
			break
		}
		fmt.Fprintf(Stderr, "!isPrefix after %d runes in string!\n", len(s))
	}
	return
}

func rpl() {
	r := bufio.NewReader(Stdin)
	for {
		line, err := readLine(r)
		if err != nil {
			fmt.Fprintf(Stderr, "err=%v\n", err)
			return
		}
		if len(line) > 0 && line[0] == 'E' {
			fmt.Fprintf(Stderr, "to stderr: %s\n", line[1:])
		} else if line == "." {
			return
		} else {
			fmt.Fprintf(Stdout, "to stdout: %s\n", line)
		}
	}
}
