package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

var Session gocqlx.Session

func InitDatabase(db_host string) gocqlx.Session {
	log.Info("Initializing database")

	cluster := gocql.NewCluster(db_host)
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.WithError(err).Fatal("Error creating database session")
	}
	log.WithField("session", session).Info("Created database session")

	err = session.ExecStmt(`CREATE KEYSPACE IF NOT EXISTS shortening WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`)
	if err != nil {
		log.WithError(err).Fatal("Error creating keyspace")
	}
	log.Info("Created database keyspace")

	err = session.ExecStmt(`CREATE TABLE IF NOT EXISTS shortening.links (
		url text,
		short text PRIMARY KEY)`)
	if err != nil {
		log.WithError(err).Fatal("Error creating table")
	}
	log.Info("Created database tables")

	log.Info("Database initialized successfully")
	Session = session
	return session
}
