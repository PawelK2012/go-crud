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
func (s *Postgres) GetAll(ctx context.Context, table string) (*sql.Rows, error) {
	var all map[string]interface{}
	//all := ClientDBResult{}
	query := `SELECT * FROM ` + table
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		log.Println("get all docs failed:", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		//var s []string
		if err := rows.Scan(all); err != nil {
			log.Println("get all rows.Scan() error:", err)
			return nil, err
		}
		// all = append(all.Documents, Document{
		// 	Id:         *createDocumentResponse.ID,
		// 	Properties: document,
		// })
	}
	fmt.Printf("---> retrived rows 111 %+v\n", rows)
	return rows, nil
}

// TODO: think about better approche for CREATE operations
func (s *Postgres) Create(ctx context.Context, table string, n *models.Note) (int64, error) {
	var id int64
	query := `INSERT INTO ` + table + `
	(author, title, description, tags, created_at)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := s.db.QueryRowContext(ctx, query, n.Author, n.Title, n.Desc, n.Tags, n.CreatedAt).Scan(&id)
	if err != nil {
		log.Println("creating note failed with error:", err)
		return id, err
	}
	n.Id = id
	log.Printf("note id:%s was created", id)
	return id, nil
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
