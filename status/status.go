package status

type Links []string

func (l Links) Make(links []string)  {
    for _, link := range links {
        l = append(l, link)
    }
}

