// Solution to part 1 of the Whispering Gophers code lab.
// This program reads from standard input and writes JSON-encoded messages to
// standard output. For example, this input line:
//	Hello!
// Produces this output:
//	{"Body":"Hello!"}
//
package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type Message struct {
	Body string
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	e := json.NewEncoder(os.Stdin)
	for s.Scan() {
		m := Message{Body: s.Text()}
		if err := e.Encode(m); err != nil {
			log.Fatal(err)
		}
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
}
