package repositories

import (
	"context"

	"github.com/luanle13/xianyin-backend/user-service/models"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int) error
}
