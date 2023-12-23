package main

import (
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

type Store interface {
	GetMenu(id int) (*models.Menu, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgresStore, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", DB_USER, DB_USER, DB_PASSWORD)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateMenuTable()
}

func (s *PostgresStore) CreateMenuTable() error {
	query := `CREATE TABLE IF NOT EXISTS menu (
		id serial primary key,
		menu_name varchar(50) DEFAULT 'Standard',
		breakfast varchar(50) DEFAULT 'Full Irish breakfast',
		large_breakfast varchar(50),
		lunch varchar(50),
		large_lunch varchar(50),
		dinner varchar(50),
		large_dinner varchar(50),
		kids_menu varchar(50),
		desert varchar(50),
		drink varchar(50),
		sides varchar(50),
		created_at timestamp
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) GetMenu(id int) (*models.Menu, error) {
	// For simplicity of this example we have only 1 menu
	// I could easly create more menus ie Lunch, set menu, evening menu etc.
	// rows, err := s.db.Query(`select * from menu where id = $1`, id)
	rows, err := s.db.Query(`select * from menu`)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoMenu(rows)
	}

	return nil, fmt.Errorf("menu %d not found", id)
}

func scanIntoMenu(rows *sql.Rows) (*models.Menu, error) {
	menu := new(models.Menu)
	err := rows.Scan(
		&menu.ID,
		&menu.MenuName,
		&menu.Breakfast,
		&menu.LargeBreakfast,
		&menu.Lunch,
		&menu.LargeLunch,
		&menu.Dinner,
		&menu.LargeDinner,
		&menu.KidsMenu,
		&menu.Desert,
		&menu.Drink,
		&menu.Sides)

	log.Print("===>>. %+v\n", menu)
	return menu, err
}
