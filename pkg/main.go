package main

import (
	"math/rand"
	"bufio"
	"log"
	"net"
	"strings"
	"time"
)

const (
	addr = "0.0.0.0:12345" // Сетевой адрес
	proto = "tcp4" // Протокол сетевой службы
)

func main() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	go func() {
		for{
			time.Sleep(time.Second * 3)
	     	writeProverb(conn)
		}
	}()

	reader := bufio.NewReader(conn)
	b, err := reader.ReadBytes('\n')
	if err != nil {
		log.Println(err)
		return
	}

	msg := strings.TrimSuffix(string(b), "\n")
	msg = strings.TrimSuffix(msg, "\r")

	if msg == "bye" {
		conn.Write([]byte("Goodbye!\n"))
	}
}

func writeProverb(conn net.Conn) {
	conn.Write([]byte(proverbs[rand.Intn(len(proverbs))]))
}

var proverbs = []string{
"Don't communicate by sharing memory, share memory by communicating.\n",
"Concurrency is not parallelism.\n",
"Channels orchestrate; mutexes serialize.\n",
"The bigger the interface, the weaker the abstraction.\n",
"Make the zero value useful.\n",
"interface{} says nothing.\n",
"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.\n",
"A little copying is better than a little dependency.\n",
"Syscall must always be guarded with build tags.\n",
"Cgo must always be guarded with build tags.\n",
"Cgo is not Go.\n",
"With the unsafe package there are no guarantees.\n",
"Clear is better than clever.\n",
"Reflection is never clear.\n",
"Errors are values.\n",
"Don't just check errors, handle them gracefully.\n",
"Design the architecture, name the components, document the details.\n",
"Documentation is for users.\n",
"Don't panic.\n",
}