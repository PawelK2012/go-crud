package postgress

import (
	"context"
	"database/sql"
	"log"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PawelK2012/go-crud/models"
	_ "github.com/lib/pq"
)

func TestPostgress_GetById(t *testing.T) {

	n := createNote(1, time.Now())

	rows := sqlmock.NewRows([]string{"id", "author", "title", "desc", "tags", "created_at"}).
		AddRow(n.Id, n.Author, n.Title, n.Desc, n.Tags, n.CreatedAt)

	rowsErr := sqlmock.NewRows([]string{"id", "author", "title", "desc", "tags", "created_at"})

	db, mock := newMock()

	type fields struct {
		db *sql.DB
	}
	testFields := fields{
		db: db,
	}
	type args struct {
		ctx   context.Context
		id    int
		query string
		rows  *sqlmock.Rows
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Note
		wantErr bool
	}{
		{name: "GetById happy flow", fields: testFields, args: args{id: int(n.Id), ctx: context.Background(), query: "SELECT (.+) FROM notes WHERE", rows: rows}, want: n},
		{name: "GetById sad flow", fields: testFields, args: args{id: 2222, ctx: context.Background(), query: "SELECT (.+) FROM notes WHERE", rows: rowsErr}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			repo := &Postgress{
				db: tt.fields.db,
			}
			mock.ExpectQuery(tt.args.query).
				WithArgs(tt.args.id).
				WillReturnRows(rows)

			got, err := repo.GetById(tt.args.ctx, tt.args.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("Postgress.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Postgress.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgress_GetAll(t *testing.T) {
	db, mock := newMock()

	ti := time.Now()

	var notes []models.Note
	n := createNote(1, ti)
	n1 := createNote(2, ti)
	n2 := createNote(3, ti)
	notes = append(notes, n, n1, n2)

	rows := sqlmock.NewRows([]string{"id", "author", "title", "desc", "tags", "created_at"})

	for _, r := range notes {
		rows.AddRow(r.Id, r.Author, r.Title, r.Desc, r.Tags, r.CreatedAt)
	}

	type fields struct {
		db *sql.DB
	}
	testFields := fields{
		db: db,
	}
	type args struct {
		ctx   context.Context
		query string
		rows  *sqlmock.Rows
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Note
		wantErr bool
	}{
		{name: "GetAll happy flow", fields: testFields, args: args{ctx: context.Background(), query: "SELECT (.+) FROM notes", rows: rows}, want: notes},
		{name: "GetAll sad flow", fields: testFields, args: args{ctx: context.Background(), query: "SELECT (.+) FROM notes where error", rows: rows}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Postgress{
				db: tt.fields.db,
			}
			mock.ExpectQuery(tt.args.query).
				WillReturnRows(rows)

			got, err := s.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Postgress.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Postgress.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgress_Create(t *testing.T) {
	n := createNote(22, time.Now())
	db, mock := newMock()
	type fields struct {
		db *sql.DB
	}
	testFields := fields{
		db: db,
	}
	type args struct {
		ctx   context.Context
		n     models.Note
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{name: "Create happy flow", fields: testFields, args: args{ctx: context.Background(), query: "INSERT INTO notes (author, title, description, tags, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id", n: n}, want: 22},
		{name: "Create sad flow", fields: testFields, args: args{ctx: context.Background(), query: "INSERT INTO badquery (author, description, tags, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id", n: n}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Postgress{
				db: tt.fields.db,
			}

			//using regexp.QuoteMeta() as otherwise SQL query doesn't match
			mock.ExpectQuery(regexp.QuoteMeta(tt.args.query)).
				WithArgs(tt.args.n.Author, tt.args.n.Title, tt.args.n.Desc, tt.args.n.Tags, tt.args.n.CreatedAt).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(22))

			got, err := s.Create(tt.args.ctx, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Postgress.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Postgress.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgress_Update(t *testing.T) {
	db, mock := newMock()
	type fields struct {
		db *sql.DB
	}
	testFields := fields{
		db: db,
	}
	type args struct {
		ctx          context.Context
		id           string
		n            models.Note
		lastInsertID int64
		rowsAffected int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Note
		wantErr bool
	}{
		{name: "GetAll happy flow", fields: testFields, args: args{ctx: context.Background(), lastInsertID: 1, rowsAffected: 1}},
		{name: "GetAll sad flow", fields: testFields, args: args{ctx: context.Background(), lastInsertID: 0, rowsAffected: 0}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Postgress{
				db: tt.fields.db,
			}

			mock.ExpectExec(regexp.QuoteMeta("UPDATE notes")).WillReturnResult(sqlmock.NewResult(tt.args.lastInsertID, tt.args.rowsAffected))

			got, err := s.Update(tt.args.ctx, tt.args.id, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Postgress.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Postgress.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgress_Init(t *testing.T) {
	db, mock := newMock()
	type fields struct {
		db *sql.DB
	}
	testFields := fields{
		db: db,
	}
	type args struct {
		ctx          context.Context
		lastInsertID int64
		rowsAffected int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "GetAll happy flow", fields: testFields, args: args{ctx: context.Background(), lastInsertID: 1, rowsAffected: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Postgress{
				db: tt.fields.db,
			}
			mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS notes")).WillReturnResult(sqlmock.NewResult(tt.args.lastInsertID, tt.args.rowsAffected))

			if err := s.Init(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Postgress.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func newMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func createNote(id int64, t time.Time) models.Note {
	return models.Note{
		Id:        id,
		Author:    "mr x",
		Title:     "test note",
		Desc:      "some desc",
		Tags:      "#test",
		CreatedAt: t,
	}
}
