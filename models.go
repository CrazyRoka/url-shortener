package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/scylladb/gocqlx/table"
	"github.com/scylladb/gocqlx/v2"
)

var (
	ShortLinksTable = table.New(table.Metadata{
		Name:    "shortening.links",
		Columns: []string{"url", "short"},
		PartKey: []string{"short"},
	})
)

type ShortLink struct {
	Url   string `db:"url" json:"url"`
	Short string `db:"short" json:"short"`
}

func InsertShortLink(session *gocqlx.Session, shortLink *ShortLink) error {
	query := session.Query(ShortLinksTable.Insert()).BindStruct(shortLink)
	log.WithFields(log.Fields{
		"query":     query,
		"statement": query.Statement(),
		"link":      shortLink,
	}).Info("Inserting link into the database")

	err := query.ExecRelease()
	if err != nil {
		log.WithError(err).WithField("link", shortLink).Warn("Error insering link into the database")
		return err
	}

	log.WithField("link", shortLink).Info("Successfully inserted link")
	return nil
}

func GetShortLink(session *gocqlx.Session, short string) (*ShortLink, error) {
	var shortLink ShortLink
	query := session.Query(ShortLinksTable.Get("url", "short")).Bind(short)
	log.WithFields(log.Fields{
		"query":     query,
		"statement": query.Statement(),
		"short":     short,
	}).Info("Searching url for short")

	err := query.GetRelease(&shortLink)
	if err != nil {
		log.WithError(err).WithField("short", short).Warn("Link not found")
		return nil, err
	}

	return &shortLink, err
}
