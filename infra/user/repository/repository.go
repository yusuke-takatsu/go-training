package repository

import (
	"database/sql"
	"github.com/yusuke-takatsu/go-training/domain/user/entity"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) Repository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(user *entity.User) error {
	query := `insert into users (id, email, password, image, status, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.Exec(query,
		user.ID,
		user.Email.Value(),
		user.Password.Value(),
		user.Image,
		user.Status,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}
