package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/syumai/go-utils/netutil"
)

func main() {
	portStr := os.Args[1]
	port, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	ln, err := netutil.TryListen("tcp", int(port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("listening address:", ln.Addr().String())
	_ = ln
}
