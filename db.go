package utiles

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

func NewDB(host string, port int, user, password, dbname string) (*DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &DB{
		Conn: db,
	}, nil
}

func (d *DB) Close() error {
	return d.Conn.Close()
}

func (d *DB) InsertMessage(clientID, message string) error {
	stmt, err := d.Conn.Prepare("INSERT INTO messages (client_id, message) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(clientID, message)
	if err != nil {
		return err
	}

	return nil
}
