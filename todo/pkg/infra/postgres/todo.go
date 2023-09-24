package postgres

import (
	"context"
	"go-micro-sample/todo/pkg/domain/entity"
	"go-micro-sample/todo/pkg/domain/repository"

	"github.com/jmoiron/sqlx"
)

type todoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) repository.ITodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (tr *todoRepository) SelectAll(ctx context.Context, id int64) ([]*entity.Todo, error) {
	var todos []*entity.Todo
	err := tr.db.SelectContext(ctx, &todos, "select * from todos where user_id = $1 and done = false", id)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (tr *todoRepository) Insert(ctx context.Context, title string, id int64) (*entity.Todo, error) {
	todo := &entity.Todo{
		Title:  title,
		UserID: id,
	}
	query := `INSERT INTO todos (title, user_id) VALUES (:title, :user_id) RETURNING id, done`
	row, err := tr.db.NamedQueryContext(ctx, query, todo)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if row.Next() {
		if err := row.StructScan(&todo); err != nil {
			return nil, err
		}
	}
	return todo, nil
}

func (tr *todoRepository) Update(ctx context.Context, id int64) (bool, error) {
	query := `UPDATE todos SET done = $1 WHERE id = $2`
	_, err := tr.db.ExecContext(ctx, query, true, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
