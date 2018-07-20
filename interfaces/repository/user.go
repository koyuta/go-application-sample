package repository

import (
	"context"

	"github.com/koyuta/go-application-sample/domain"
)

type User struct {
	MySQLHandler
}

func (u *User) Store(ctx context.Context, user domain.User) (int64, error) {
	q := `INSERT INTO users (id, name, description, updated_at, created_at)
VALUES (?, ?, ?, NOW(), NOW())`
	result, err := u.ExecuteContext(ctx, q,
		user.ID,
		user.Name,
		user.Description,
		user.UpdatedAt,
		user.CreatedAt,
	)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func (u *User) FindByID(ctx context.Context, id int64) (domain.User, error) {
	row, err := u.QueryContext(ctx, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return domain.User{}, err
	}
	defer row.Close()

	var d domain.User
	row.Next()
	if err = row.Scan(
		&d.ID,
		&d.Name,
		&d.Description,
		&d.UpdatedAt,
		&d.CreatedAt,
	); err != nil {
		return d, err
	}
	return d, nil
}

func (u *User) FindAll(ctx context.Context) ([]domain.User, error) {
	row, err := u.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return []domain.User{}, err
	}
	defer row.Close()

	var s []domain.User
	for row.Next() {
		var d domain.User
		if err = row.Scan(
			&d.ID,
			&d.Name,
			&d.Description,
			&d.UpdatedAt,
			&d.CreatedAt,
		); err != nil {
			return []domain.User{}, err
		}
		s = append(s, d)
	}
	return s, nil
}
