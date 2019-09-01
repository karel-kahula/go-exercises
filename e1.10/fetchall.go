package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	buf := new(bytes.Buffer)

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()
	nbytes, err := io.Copy(buf, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	n, f, err := createFile(url, start)
	if err != nil {
		ch <- fmt.Sprintf("Could not create file \"%s\": %v", n, err)
	}
	defer f.Close()
	_, err = f.WriteString(buf.String())
	if err != nil {
		ch <- fmt.Sprintf("Could not write to file \"%s\": %v", n, err)
	}
	f.Sync()

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func createFile(n string, t time.Time) (string, *os.File, error) {
	n = strings.Replace(n, "/", "", -1)
	n = strings.Replace(n, "http:", "", -1)
	ts := fmt.Sprintf("%d%d%d", t.Hour(), t.Minute(), t.Second())
	fn := fmt.Sprintf("%s.%s", n, ts)
	f, err := os.Create(fn)
	return fn, f, err
}
