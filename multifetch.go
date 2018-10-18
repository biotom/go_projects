package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	w, err := os.Create("multifetch.txt")
	if err != nil {
		fmt.Printf("Problem opening file")
	}
	//w := bufio.NewWriter(f)
	for range os.Args[1:] {
		fmt.Fprintf(w, <-ch)
	}
	fmt.Fprintf(w, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()

	if err != nil {
		ch <- fmt.Sprintf("error while creating file")
		return
	}
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}
