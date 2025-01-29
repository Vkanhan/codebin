package models

import (
	"database/sql"
)

type Gist struct {
	ID        string
	Title     string
	Content   string
	CreatedAt string
	Expires   string
}

type GistModel struct {
	DB *sql.DB
}
