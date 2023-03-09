package main

import (
    "bufio"
    "fmt"
    "io"
    "net/http"
    "os"

    "github.com/suedoh/go-http-scraper/status"
)

type logWriter struct {}

func main() {
    links := []string{
        "http://duckduckgo.com",
        "http://cht.sh",
        "http://github.com",
        "http://hotmail.com",
        "http://golang.com",
    }

    status := &status.Checker{}
    linksToCheck := status.Make(links)
    status.Check(linksToCheck)

    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Println("Would you like to scrape HTML? (y/n)")
        fmt.Print("Enter answer: ")
        answer, _ := reader.ReadString('\n')

        // change to switch
        if answer == "n" || answer == "no" {
            fmt.Println("k Bye")
            os.Exit(1)
        } else {
            break
        }

    }

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
