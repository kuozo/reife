package trello

import (
	"strconv"
)

const (
	ENDPOINT = "https://api.trello.com/"
	VERSION  = 1
)

type Trello struct {
	key   string
	token string
}

func NewTrello(key, token string) (t *Trello) {
	t = new(Trello)
	t.key = key
	t.token = token
	return t
}

func (t *Trello) spliceUrl(sub string) string {
	url := ENDPOINT + strconv.Itoa(VERSION) + "/" + sub + ""
	return url
}
