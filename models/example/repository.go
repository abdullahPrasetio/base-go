package example

import (
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"time"
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	Create(param RequestCreateEmployee, ctx context.Context) (Employee, error)
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(param RequestCreateEmployee, ctx context.Context) (Employee, error) {
	newctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()
	var SQL string
	var query sq.InsertBuilder
	var result Employee

	query = sq.Insert("employee").Columns("name", "age", "address", "created_at", "updated_at").
		Values(param.Name, param.Age, param.Address)
	SQL, _, err := query.ToSql()
	res, err := r.db.ExecContext(newctx, SQL)

	if err != nil {
		return result, err
	}
	no, err := res.RowsAffected()
	if err != nil {
		return result, err
	}

	fmt.Println(no)
	return result, nil
}
