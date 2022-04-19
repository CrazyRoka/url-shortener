package main

import (
	"net/url"
	"strings"
	"testing"
)

func TestGenerateShort(t *testing.T) {
	for n := uint(1); n <= 10; n++ {
		short := GenerateShort(n)
		verifyShort(t, short, n)
	}
}

func TestBuildShortenLink(t *testing.T) {
	url, _ := url.Parse("www.google.com")
	link := BuildShortenLink(url)
	verifyShortenLink(t, link, url)
}

func verifyShortenLink(t *testing.T, link ShortLink, url *url.URL) {
	if link.Url != url.String() {
		t.Errorf("Expected url %s, but got %s", url.String(), link.Url)
	}

	verifyShort(t, link.Short, SHORT_LENGTH)
}

func verifyShort(t *testing.T, short string, expectedLen uint) {
	if uint(len(short)) != expectedLen {
		t.Errorf("Expected generated short to be of length %d, but was %d", expectedLen, len(short))
	}

	for _, char := range short {
		if !strings.ContainsRune(string(letters), char) {
			t.Errorf("Unexpected char %c", char)
		}
	}
}
