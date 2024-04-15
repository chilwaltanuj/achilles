package model

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// ClientRDBMS holds the database client and its configuration.
type ClientRDBMS struct {
	DB     *sqlx.DB          // The SQLx database client
	Config ClientRDBMSConfig // The configuration for the RDBMS client
	Logger *logrus.Logger    // Logger instance for logging database operations
}

type RequestClientRDBMS struct {
	Query     string
	Arguments []interface{}
}
