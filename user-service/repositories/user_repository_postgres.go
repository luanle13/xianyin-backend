package repositories

import (
	"context"
	"database/sql"

	"github.com/luanle13/xianyin-backend/user-service/models"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, "SELECT id, username, email, password FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.QueryRowContext(ctx,
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id",
		user.Username, user.Email, user.Password,
	).Scan(&user.ID)
}

func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	_, err := r.db.ExecContext(ctx,
		"UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4",
		user.Username, user.Email, user.Password, user.ID,
	)
	return err
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	return err
}
