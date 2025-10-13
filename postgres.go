package database

import (
    "database/sql"
    _ "github.com/lib/pq"
)

func ConnectPostgres() (*sql.DB, error) {
    return sql.Open("postgres", "user=postgres password=yourpassword dbname=deliverydb sslmode=disable")
}
