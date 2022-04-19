package main

import (
	"math/rand"
	"net/url"
)

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

const SHORT_LENGTH uint = 6

func GenerateShort(n uint) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func BuildShortenLink(uri *url.URL) ShortLink {
	return ShortLink{
		Url:   uri.String(),
		Short: GenerateShort(SHORT_LENGTH),
	}
}
