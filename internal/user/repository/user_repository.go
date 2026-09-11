package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ix1ax/social-network-backend/internal/user/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {

	query :=
		`
		INSERT INTO users (name,surname,email,password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
		`

	err := r.db.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Surname,
		user.Email,
		user.PasswordHash,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return err
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {

	query :=
		`
		SELECT id, name, surname, email, password_hash, created_at, updated_at
		FROM users WHERE email = $1
		`

	var user entity.User

	err := r.db.GetContext(ctx, &user, query, email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {

	query :=
		`
		SELECT id, name, surname, email, password_hash, created_at, updated_at
		FROM users WHERE id = $1
		`

	var user entity.User

	err := r.db.GetContext(ctx, &user, query, id)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
