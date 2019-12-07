package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/koooyooo/grppt/core/client"
)

func main() {
	fmt.Println("Run Client")
	st := time.Now()
	conn, cli, err := client.CreateClient()
	defer conn.Close()
	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.RunHttpClient(*cli, req)
	if err != nil {
		log.Fatal(err)
	}
	ed := time.Now()
	dumpBytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dumpBytes))
	fmt.Println("Total-Time: %d", ed.Sub(st).Milliseconds())
}
