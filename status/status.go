package status

import (
	"fmt"
	"net/http"
)

type Checker []string

type StatusChecker interface {
    Make(links []string) []string
    Check(links []string)
}

func (s *Checker) Make(links []string) []string {
    for _, link := range links {
        *s = append(*s, link)
    }

    return *s
}

func (*Checker) Check(links []string)  {
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
