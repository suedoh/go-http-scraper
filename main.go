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

        var response *http.Response
        // change to switch
        switch {
        case answer == "n" || answer == "no":
            fmt.Println("k Bye")
            os.Exit(1)
        case answer == "y" || answer == "yes":
            fmt.Print("Enter url to scrape: ")
            url, _ := reader.ReadString('\n')
            response = getUrl(url)
        }

        lw := logWriter{}
        io.Copy(lw, response.Body)
        break
    }


}

func getUrl(url string) *http.Response {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    return resp
}

func (lw logWriter) Write(bs []byte) (int, error) {
    byteLength := len(bs)

    fmt.Println(string(bs))
    fmt.Printf("Just wrote %v bytes", byteLength)

    return byteLength, nil
}
