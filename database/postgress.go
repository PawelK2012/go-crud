package database

import (
	"context"
	"database/sql"
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

type Postgres struct {
	db *sql.DB
}

func NewPostgressClient() (ClientInterface, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", DB_USER, DB_USER, DB_PASSWORD)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{
		db: db,
	}, err
}

func (s *Postgres) Init(ctx context.Context) error {
	return s.CreateNoteTable()
}

func (s *Postgres) CreateNoteTable() error {
	query := `create table if not exists notes (
		id serial primary key,
		author varchar(200),
		title varchar(250),
		description varchar(5000),
		tags varchar(250),
		created_at timestamp
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *Postgres) CreateNote(ctx context.Context, n *models.Note) error {
	query := `insert into notes
	(author, title, description, tags, created_at)
	values ($1, $2, $3, $4, $5)`

	_, err := s.db.Query(query, n.Author, n.Title, n.Desc, n.Tags, n.CreatedAt)
	if err != nil {
		return err
	}
	// log.Printf("resp:  %+v\n", resp())
	return nil
}

func (s *Postgres) GetNoteById(ctx context.Context, id int) (*models.Note, error) {

	rows, err := s.db.Query(`select * from note`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		// return scanIntoMenu(rows)
	}

	return nil, fmt.Errorf("menu %d not found", id)
}

// func scanIntoMenu(rows *sql.Rows) (*models.Menu, error) {
// 	menu := new(models.Menu)
// 	err := rows.Scan(
// 		&menu.ID,
// 		&menu.MenuName,
// 		&menu.Breakfast,
// 		&menu.LargeBreakfast,
// 		&menu.Lunch,
// 		&menu.LargeLunch,
// 		&menu.Dinner,
// 		&menu.LargeDinner,
// 		&menu.KidsMenu,
// 		&menu.Desert,
// 		&menu.Drink,
// 		&menu.Sides)

// 	return menu, err
// }
