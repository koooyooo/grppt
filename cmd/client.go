package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"sync"
	"time"

	"github.com/koooyooo/grppt/core/client"
)

func main() {
	fmt.Println("Run Client")
	st := time.Now()
	conn, cli, err := client.CreateClient("localhost:5051")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	var threadCount = 10
	var callCount = 300
	var wg sync.WaitGroup
	wg.Add(threadCount * callCount)
	for t := 0; t < threadCount; t++ {
		go func() {
			for c := 0; c < callCount; c++ {
				st := time.Now()
				req, err := http.NewRequest("GET" /*"https://httpbin.org/get"*/, "http://localhost:80", nil)
				if err != nil {
					log.Fatal(err)
				}
				resp, err := client.RunHttpClient(*cli, req)
				if err != nil {
					log.Fatal(err)
				}
				ed := time.Now()
				logResponseLine(os.Stdout, ed.Sub(st), resp)
				wg.Done()
			}
		}()
	}
	wg.Wait()
	ed := time.Now()
	fmt.Printf("Total-Time: %d \n", ed.Sub(st).Milliseconds())
}

func logResponseDump(w io.Writer, d time.Duration, resp *http.Response) {
	dump, err := httputil.DumpResponse(resp, true)
	if err == nil {
		w.Write(dump)
	}
}

func logResponseLine(w io.Writer, d time.Duration, resp *http.Response) {
	fmt.Printf("%d %v %v\n", resp.StatusCode, d, resp.Header)
}
