package repository

import (
	"context"
	"database/sql"
	"patijournal/pkg/pb"

	_ "github.com/lib/pq"
)

type EntryPostgresRepository struct {
	db *sql.DB
}

func NewEntryPostgresRepository(connString string) (*EntryPostgresRepository, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	// Ensure the database connection is valid
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	// Create entries table if it doesn't exist
	createTable := `CREATE TABLE IF NOT EXISTS entries (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		body TEXT NOT NULL,
		image TEXT
	);`

	if _, err := db.Exec(createTable); err != nil {
		db.Close()
		return nil, err
	}

	return &EntryPostgresRepository{db: db}, nil
}

func (r *EntryPostgresRepository) GetAll(ctx context.Context) ([]*pb.Entry, error) {
	query := "SELECT id, title, image FROM entries"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var entries []*pb.Entry
	for rows.Next() {
		entry := &pb.Entry{}
		err := rows.Scan(&entry.Id, &entry.Title, &entry.Image)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (r *EntryPostgresRepository) GetByID(ctx context.Context, id int32) (*pb.Entry, error) {
	query := "SELECT id, title, body, image FROM entries WHERE id=$1"
	row := r.db.QueryRow(query, id)

	entry := &pb.Entry{}
	err := row.Scan(&entry.Id, &entry.Title, &entry.Body, &entry.Image)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *EntryPostgresRepository) Create(ctx context.Context, entry *pb.Entry) (*pb.Entry, error) {
	query := "INSERT INTO entries (title, body, image) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, entry.Title, entry.Body, entry.Image).Scan(&entry.Id)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *EntryPostgresRepository) Update(ctx context.Context, entry *pb.Entry) (*pb.Entry, error) {
	query := "UPDATE entries SET title=$1, body=$2, image=$3 WHERE id=$4"
	_, err := r.db.Exec(query, entry.Title, entry.Body, entry.Image, entry.Id)
	if err != nil {
		return nil, err
	}
	return entry, nil
}
