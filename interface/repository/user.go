package repository

import "context"

type User struct {
	MySQLHandler
}

func (u *User) Store(ctx context.Context, user domain.User) (int64, error) {
	q := `INSERT INTO users (id, name, description, updated_at, created_at)
VALUES (?, ?, ?, NOW(), NOW())`
	result, err := c.ExecuteContext(ctx, query,
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
	row, err := a.QueryContext(ctx, "SELECT * FROM users WHERE id = ?", id)
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
		&d.UpdateAt,
		&d.CreatedAt,
	); err != nil {
		return d, err
	}
	return d, nil
}

func (u *User) FindAll(ctx context.Context) ([]domain.User, error) {
	row, err := a.QueryContext(ctx, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return domain.User{}, err
	}
	defer row.Close()

	var d domain.User
	for row.Next() {
		if err = row.Scan(
			&d.ID, &d.Name,
			&d.Description,
			&d.UpdateAt,
			&d.CreatedAt,
		); err != nil {
			return d, err
		}
	}
	return d, nil
}
