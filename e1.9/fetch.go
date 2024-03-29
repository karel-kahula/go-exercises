package main

import (
        "fmt"
        "io"
        "net/http"
        "os"
        "strings"
       )

func main() {
    for _, url := range os.Args[1:] {
        if ! strings.HasPrefix(url, "http://") {
            url = "http://" + url
        }
        resp, err := http.Get(url)
            if err != nil {
                fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
                os.Exit(1)
            }
            if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
                fmt.Fprintf(os.Stderr, "fetch %v\n", err)
                os.Exit(1)
            }
            resp.Body.Close()
            if err != nil {
                fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
                os.Exit(1)
            }

            fmt.Fprintf(os.Stdout, "\n\nHTTP Status Code: %s\n", resp.Status)
    }
}

