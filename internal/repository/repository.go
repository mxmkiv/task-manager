package repository

import (
	"context"
	"errors"
	"log"
	"task-manager/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const DUPLICATE_ERROR_CODE = "23505"

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(user *model.User) error {
	query := "INSERT INTO users (login, password_hash, role) VALUES ($1, $2, $3) RETURNING id;"

	err := u.db.QueryRow(context.Background(), query, user.Login, user.PasswordHash, user.Role).Scan(&user.Id)
	if err != nil {
		var dbErr *pgconn.PgError
		if errors.As(err, &dbErr) {
			if dbErr.Code == DUPLICATE_ERROR_CODE {
				return errors.New("user with this login already exists")
			}
		}
		log.Println("err: ", err)
		return err
	}

	return nil
}

func (u *UserRepository) GetByLogin(login string) (*model.User, error) {

	query := "SELECT * FROM users WHERE login = $1;"

	user := &model.User{}
	err := u.db.QueryRow(context.Background(), query, login).Scan(
		&user.Id, &user.Login, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("no user found")
		}
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetAllUsers() (*[]model.User, error) {

	query := "SELECT * FROM users LIMIT 50 OFFSET 0;"

	rows, err := u.db.Query(context.Background(), query)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("no user founds")
		}
		return nil, err
	}
	defer rows.Close()

	res, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.User])
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (u *UserRepository) GetUserById(id int) (*model.User, error) {
	query := "SELECT * FROM users WHERE id = $1"

	user := &model.User{}
	err := u.db.QueryRow(context.Background(), query, id).Scan(
		&user.Id, &user.Login, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("no user founds")
		}
		return nil, err
	}

	return user, nil
}
