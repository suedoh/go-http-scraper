package status

type Links []string

type StatusChecker interface {
    Make([]string) []string
}

func (l Links) Make(links []string) []string {
    for _, link := range links {
        l = append(l, link)
    }

    return l
}

