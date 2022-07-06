package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	log.SetFlags(0)
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	if len(args) < 2 {
		log.Println("usage: foreachline file.txt command [command args]" +
			"\nCommand will be called for each line, with line passed over stdin." +
			"\nEmpty lines are skipped.")
		os.Exit(2)
	}
	f, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	var lineno int
	for sc.Scan() {
		lineno++
		if len(sc.Bytes()) == 0 {
			continue
		}
		cmd := exec.Command(args[1], args[2:]...)
		cmd.Stdin = bytes.NewReader(sc.Bytes())
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("line %d: %w", lineno, err)
		}
	}
	return sc.Err()
}
