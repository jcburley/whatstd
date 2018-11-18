package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] != "" {
		replSocket := os.Args[1]
		l, err := net.Listen("tcp", replSocket)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot start srepl listening on %s: %s\n",
				replSocket, err.Error())
			os.Exit(12)
		}
		defer l.Close()

		fmt.Printf("Joker repl listening at %s...\n", l.Addr())
		conn, err := l.Accept() // Wait for a single connection
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot start repl accepting on %s: %s\n",
				l.Addr(), err.Error())
			os.Exit(13)
		}

		oldStdin := os.Stdin
		oldStdout := os.Stdout
		oldStderr := os.Stderr

		/* This does not work: "cannot use conn (type net.Conn) as type *osFile in assignment". */
		os.Stdin = conn
		/* (Casting doesn't work, either.) */
		os.Stdout = os.File(conn)
		os.Stderr = (*os.File)(conn)

		defer func() {
			conn.Close()
			os.Stdin = oldStdin
			os.Stdout = oldStdout
			os.Stderr = oldStderr
		}()

		fmt.Printf("Client at %s. Use '.', or close the connection, to exit.\n", conn.RemoteAddr())
	}
	rpl()
	fmt.Printf("final stdout\n")
	fmt.Fprintf(os.Stderr, "final stderr\n")
}
