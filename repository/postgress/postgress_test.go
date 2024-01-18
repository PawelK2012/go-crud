package postgress

import (
	"context"
	"database/sql"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PawelK2012/go-crud/models"
	_ "github.com/lib/pq"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestPostgress_GetById(t *testing.T) {

	timeNow := time.Now()

	var n models.Note
	n.Id = 1
	n.Author = "mr x"
	n.Title = "test note"
	n.Desc = "some desc"
	n.Tags = "#test"
	n.CreatedAt = timeNow

	rows := sqlmock.NewRows([]string{"id", "author", "title", "desc", "tags", "created_at"}).
		AddRow(n.Id, n.Author, n.Title, n.Desc, n.Tags, n.CreatedAt)

	rowsErr := sqlmock.NewRows([]string{"id", "author", "title", "desc", "tags", "created_at"})

	db, mock := NewMock()

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
