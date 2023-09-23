package postgres

import (
	"context"
	"go-micro-sample/user/pkg/domain/entity"
	"go-micro-sample/user/pkg/domain/repository"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) SelectAll(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	err := ur.db.Select(&users, "select * from users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) Select(ctx context.Context, id int64) (*entity.User, error) {
	user := &entity.User{}
	err := ur.db.Get(user, "select * from users where id = $1", id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) Insert(ctx context.Context, name string) (*entity.User, error) {
	user := &entity.User{
		Name: name,
	}
	query := `INSERT INTO users (name) VALUES (:name) RETURNING id`
	row, err := ur.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if row.Next() {
		if err := row.StructScan(&user); err != nil {
			return nil, err
		}
	}
	return user, nil
}
