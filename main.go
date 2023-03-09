package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

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
        answer = strings.Replace(answer, "\n", "", -1)

        switch {
        case answer == "n" || answer == "no":
            fmt.Println("k Bye")
        case answer == "y" || answer == "yes":
            fmt.Print("Enter url to scrape: ")
            u, _ := reader.ReadString('\n')
            url := strings.Replace(u, "\n", "", -1) 
            getUrl(url)
        default:
            continue
        }

        // lw := logWriter{}
        // io.Copy(lw, response.Body)
        break
    }


}

func getUrl(url string) string {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    // ReadString
    r := bufio.NewReader(resp.Body)
    for {
        line, err := r.ReadString('\n')
        if len(line) == 0 && err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println("Error:", err)
            os.Exit(1)
        }
        line = strings.TrimSuffix(line, "\n")

        fmt.Println(line)
        return line
    }
    return "string"
}

func (lw logWriter) Write(bs []byte) (int, error) {
    byteLength := len(bs)

    fmt.Println(string(bs))
    fmt.Printf("Just wrote %v bytes", byteLength)

    return byteLength, nil
}
