package status

import (
	"fmt"
	"net/http"
)

type Links []string

type StatusChecker interface {
    Make(links []string) []string
    Check(links []string)
}

func (l Links) Make(links []string) []string {
    for _, link := range links {
        l = append(l, link)
    }

    return l
}

func (l Links) Check(links []string)  {
    for _, link := range links {
         checkStatus(link)
    }
}

func checkStatus(link string)  {
    _, err := http.Get(link)
    if err != nil {
        fmt.Println(link, "is down!")
        return
    }

    fmt.Println(link, "is up :)")
}
