package repositories

import (
	"awesomeProject/domain"
	"context"
	"database/sql"
	"github.com/gofrs/uuid"
	"log"
	"time"
)

type (
	UserEntity struct {
		ID        uuid.UUID
		Username  string
		Email     string
		Password  string
		CreatedAt time.Time `db:"created_at"`
		//UpdatedAt time.Time `db:"updated_at"`
	}

	UsersRepository interface {
		CreateUser(ctx context.Context, username, email, password string) (*UserEntity, error)
		GetUserById(ctx context.Context, id uuid.UUID) (*UserEntity, error)
		GetUserByUsernameAndEmail(ctx context.Context, username, email string) (*UserEntity, error)
		GetUserByUsername(ctx context.Context, username string) (*UserEntity, error)
		SearchUsers(ctx context.Context, username, email string) ([]UserEntity, error)
	}
	usersRepository struct {
		db *sql.DB
	}
)

func (u *UserEntity) ToUser(token string) *domain.User {
	return &domain.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Token:    token,
	}
}

func NewUsersRepository(db *sql.DB) UsersRepository {
	return &usersRepository{
		db: db,
	}
}

func (r *usersRepository) SearchUsers(ctx context.Context, username, email string) ([]UserEntity, error) {
	var users []UserEntity
	query := `
			SELECT * 
			FROM users 
			WHERE "Username" = $1 OR "Email" = $2`

	//if err := r.db.SelectContext(ctx, &user, query, username, email); err != nil {
	//	return user, err
	//}
	rows, err := r.db.QueryContext(ctx, query, username, email)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user UserEntity
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *usersRepository) GetUserById(ctx context.Context, id uuid.UUID) (*UserEntity, error) {
	var user UserEntity
	const query string = "SELECT * FROM users WHERE id = $1 LIMIT 1"

	err := r.db.QueryRowContext(ctx, query, id).Scan(&user)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (r *usersRepository) GetUserByUsernameAndEmail(ctx context.Context, username, email string) (*UserEntity, error) {
	var user UserEntity

	query := `SELECT * FROM users WHERE ("Username", "Email") = ($1, $2)`

	err := r.db.QueryRowContext(ctx, query, username, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (r *usersRepository) GetUserByUsername(ctx context.Context, username string) (*UserEntity, error) {
	var user UserEntity
	const query string = `SELECT * FROM users WHERE "Username" = $1`

	err := r.db.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (r *usersRepository) CreateUser(ctx context.Context, username string, email string, password string) (*UserEntity, error) {
	const command string = `INSERT INTO users (id, "Username", "Email", "Password", "created_at") VALUES (gen_random_uuid(), $1, $2, $3, NOW())`

	if _, err := r.db.ExecContext(ctx, command, username, email, password); err != nil {
		return &UserEntity{}, err
	}

	return r.GetUserByUsernameAndEmail(ctx, username, email)
}
