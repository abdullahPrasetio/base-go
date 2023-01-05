package example

/********************************************************************************
* Temancode Example Repository Package                                          *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/abdullahPrasetio/base-go/configs"
)

type repository struct {
	db *sql.DB
}

var (
	SQL         string
	queryInsert sq.InsertBuilder
	querySelect sq.SelectBuilder
	err         error
	arguments   []interface{}
)

type Repository interface {
	Create(param RequestEmployee) (Employee, error)
	GetAll() ([]Employee, error)
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(param RequestEmployee) (Employee, error) {
	var result Employee
	location, err := time.LoadLocation(configs.Configs.TIME_LOCATION)
	CreatedAt := time.Now().In(location).Format("2006-01-02 15:04:05")
	newctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	tx, err := r.db.Begin()
	if err != nil {
		return result, err
	}
	defer tx.Rollback()
	queryInsert = sq.Insert("employee").Columns("name", "age", "address", "created_at", "updated_at").
		Values(param.Name, param.Age, param.Address, CreatedAt, CreatedAt)
	SQL, arguments, err = queryInsert.ToSql()
	if err != nil {
		return result, err
	}
	stmt, err := tx.Prepare(SQL)
	if err != nil {
		return result, err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(newctx, arguments...)
	if err != nil {
		return result, err
	}
	err = tx.Commit()
	result.Name = param.Name
	result.Age = param.Age
	result.Address = param.Address
	result.CreatedAt = CreatedAt
	result.UpdatedAt = CreatedAt
	result.ID, _ = res.LastInsertId()

	return result, nil
}

func (r *repository) GetByID(id string) (Employee, error) {
	return Employee{}, nil
}

func (r *repository) GetAll() ([]Employee, error) {
	var results []Employee
	newctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	querySelect = sq.Select("id", "name", "age", "address", "created_at", "updated_at").
		From("employee").OrderBy("name")
	SQL, arguments, err = querySelect.ToSql()
	if err != nil {
		return results, err
	}
	rows, err := r.db.QueryContext(newctx, SQL)
	if err != nil {
		return results, err
	}
	defer rows.Close()
	for rows.Next() {
		result := Employee{}
		err := rows.Scan(&result.ID, &result.Name, &result.Age, &result.Address, &result.CreatedAt, &result.UpdatedAt)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}

	return results, nil
}
