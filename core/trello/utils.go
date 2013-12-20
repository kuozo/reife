package trello

import (
	"strconv"
)

const (
	ENDPOINT = "https://api.trello.com/"
	VERSION  = 1
)

type Trello struct {
	Key   string
	Token string
}

func NewTrello(key, token string) (t *Trello) {
	t = new(Trello)
	t.Key = key
	t.Token = token
	return t
}

func (t *Trello) spliceUrl(sub string) string {
	url := ENDPOINT + strconv.Itoa(VERSION) + "/" + sub + ""
	return url
}
