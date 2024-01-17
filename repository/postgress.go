package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/PawelK2012/go-crud/models"
	_ "github.com/lib/pq"
)

var (
	DB_USER     = os.Getenv("POSTGRES_USER_CRUDAPP")
	DB_PASSWORD = os.Getenv("POSTGRES_PASSWORD_CRUDAPP")
)

const (
	NOTES_TBL = "notes"
)

type Postgress struct {
	db *sql.DB
}

func NewPostgress() (Repository, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", DB_USER, DB_USER, DB_PASSWORD)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Postgress{
		db: db,
	}, err
}

func (s *Postgress) Init(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS notes (
		id SERIAL PRIMARY KEY,
		author VARCHAR(200),
		title VARCHAR(250),
		description VARCHAR(5000),
		tags VARCHAR(250),
		created_at TIMESTAMP
	)`
	_, err := s.db.ExecContext(ctx, query)
	return err
}

// TODO: think about better approche for CREATE operations
func (s *Postgress) Create(ctx context.Context, n models.Note) (int64, error) {
	var id int64
	query := `INSERT INTO ` + NOTES_TBL + `
	(author, title, description, tags, created_at)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := s.db.QueryRowContext(ctx, query, n.Author, n.Title, n.Desc, n.Tags, n.CreatedAt).Scan(&id)
	if err != nil {
		log.Println("creating note failed with error:", err)
		return id, err
	}
	n.Id = id
	return id, nil
}

func (s *Postgress) GetAll(ctx context.Context) ([]models.Note, error) {

	query := `SELECT * FROM ` + NOTES_TBL
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		log.Println("get all docs failed:", err)
		return nil, err
	}
	defer rows.Close()
	var all []models.Note
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.Id, &note.Author, &note.Title, &note.Desc, &note.Tags, &note.CreatedAt); err != nil {
			return nil, err
		}
		all = append(all, note)
	}
	return all, nil
}

func (s *Postgress) GetById(ctx context.Context, id int) (models.Note, error) {
	log.Println("getting note by id", id)
	var note models.Note
	query := `SELECT * FROM ` + NOTES_TBL + ` WHERE id = $1`
	row := s.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&note.Id, &note.Title, &note.Desc, &note.Author, &note.Tags, &note.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return note, fmt.Errorf("note %d not found", id)
		}
		return note, err
	}

	return note, nil
}
