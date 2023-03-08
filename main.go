package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Please include a web address to scrape as an argument")
        os.Exit(1)
    }
    url := os.Args[1]

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    lw := logWriter{}
    io.Copy(lw, resp.Body)
}

func (lw logWriter) Write(bs []byte) (int, error) {
    byteLength := len(bs)

    fmt.Println(string(bs))
    fmt.Printf("Just wrote %v bytes", byteLength)

    return byteLength, nil
}
