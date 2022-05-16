package dbops

import (
	"context"

	"github.com/zapponejosh/jellyfish/internal/models"
)

func (d DB) GetUser(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	row := d.db.QueryRow(ctx, "SELECT * FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at, &user.Profile_image)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (d DB) CreateUser(ctx context.Context, user *models.User) (int, error) {
	lastInsertId := 0

	row := d.db.QueryRow(ctx, "INSERT INTO users (name, email, password, profile_image) VALUES ($1, $2, $3, $4) RETURNING id", user.Name, user.Email, user.Password, user.Profile_image) // using queryRow to get ID since postgres doesn't support lastInsertId
	err := row.Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func (d DB) DeleteUser(ctx context.Context, id int) (int, error) {
	deletedUserId := 0

	row := d.db.QueryRow(ctx, "DELETE FROM users WHERE id = $1 RETURNING id", id)
	err := row.Scan(&deletedUserId)
	if err != nil {
		return 0, err
	}
	return deletedUserId, nil
}
