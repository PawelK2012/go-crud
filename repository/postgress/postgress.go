package postgress

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PawelK2012/go-crud/models"
	"github.com/PawelK2012/go-crud/repository"
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

func NewPostgress() (repository.Repository, error) {
	log.Println("db0:", os.Getenv("POSTGRES_USER_CRUDAPP"), os.Getenv("POSTGRES_PASSWORD_CRUDAPP"))
	log.Println("db1:", DB_USER, DB_PASSWORD)
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", DB_USER, DB_USER, DB_PASSWORD)

	m := make(map[string]string)
	for _, e := range os.Environ() {
		if i := strings.Index(e, "="); i >= 0 {
			log.Printf("env: %+v\n", i, e)
			m[e[:i]] = e[i+1:]
		}
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("db connection failed", err)
		log.Fatal(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Println("ping failed", err)
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
	query := `SELECT * FROM notes WHERE id=$1`
	row := s.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&note.Id, &note.Author, &note.Title, &note.Desc, &note.Tags, &note.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return note, fmt.Errorf("note %d not found", id)
		}
		return note, err
	}

	return note, nil
}

func (s *Postgress) Create(ctx context.Context, n models.Note) (int64, error) {
	var id int64
	query := `INSERT INTO ` + NOTES_TBL + `
	(author, title, description, tags, created_at)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := s.db.QueryRowContext(ctx, query, n.Author, n.Title, n.Desc, n.Tags, n.CreatedAt).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *Postgress) Update(ctx context.Context, id string, n models.Note) (models.Note, error) {
	query := "UPDATE notes SET author = $1, title = $2, description = $3, tags = $4 WHERE id = $5"
	res, err := s.db.ExecContext(ctx, query, n.Author, n.Title, n.Desc, n.Tags, id)

	if err != nil {
		return n, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return n, err
	}

	if rowsAffected == 0 {
		return n, repository.ErrUpdateFailed
	}

	return n, nil
}

func (s *Postgress) Delete(ctx context.Context, id int) (int, error) {
	query := "DELETE FROM notes WHERE id = $1"
	res, err := s.db.ExecContext(ctx, query, id)

	if err != nil {
		return id, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return id, err
	}

	if rowsAffected == 0 {
		return id, repository.ErrDeleteFailed
	}

	return id, nil
}
