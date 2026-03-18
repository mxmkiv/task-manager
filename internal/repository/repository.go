package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
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

func (u *UserRepository) UpdateUserData(updateList *[]model.UpdateFields, requestId int) error {

	query, args, err := UpdateQueryForm(updateList, requestId)
	if err != nil {
		return err
	}

	_, err = u.db.Exec(context.Background(), query, args...)
	if err != nil {
		return fmt.Errorf("[db] query error %s", err)
	}

	return nil
}

func UpdateQueryForm(updateList *[]model.UpdateFields, requestId int) (string, []any, error) {

	if len(*updateList) == 0 {
		return "", nil, errors.New("no fields to update")
	}

	updateParams := make([]string, len(*updateList))
	args := make([]any, len(*updateList))

	counter := 0
	for _, update := range *updateList {
		paramStr := fmt.Sprintf("%s=$%d", update.FieldName, counter+1)
		updateParams[counter] = paramStr

		args[counter] = update.Data
		counter++
	}

	args = append(args, requestId)

	query := fmt.Sprintf("UPDATE users SET %s WHERE id=$%d",
		strings.Join(updateParams, ", "), len(*updateList)+1,
	)

	return query, args, nil

}

func (u *UserRepository) DeleteUser(requestId int) error {

	query := "DELETE FROM users WHERE id=$1"

	_, err := u.db.Exec(context.Background(), query, requestId)
	if err != nil {
		return fmt.Errorf("[db] query error %s", err)
	}

	return nil

}
