// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package db

import (
	"context"
)

const addNewUser = `-- name: AddNewUser :one
INSERT INTO users(username,password,email)VALUES($1,$2,$3) RETURNING id, username, password, email, create_at, is_active, is_staff, is_admin
`

type AddNewUserParams struct {
	Username string
	Password string
	Email    string
}

func (q *Queries) AddNewUser(ctx context.Context, arg AddNewUserParams) (User, error) {
	row := q.queryRow(ctx, q.addNewUserStmt, addNewUser, arg.Username, arg.Password, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreateAt,
		&i.IsActive,
		&i.IsStaff,
		&i.IsAdmin,
	)
	return i, err
}

const disableUser = `-- name: DisableUser :exec
UPDATE users SET is_active=false
`

func (q *Queries) DisableUser(ctx context.Context) error {
	_, err := q.exec(ctx, q.disableUserStmt, disableUser)
	return err
}

const enableUser = `-- name: EnableUser :exec
UPDATE users SET is_active=true
`

func (q *Queries) EnableUser(ctx context.Context) error {
	_, err := q.exec(ctx, q.enableUserStmt, enableUser)
	return err
}

const getOneUSer = `-- name: GetOneUSer :one
SELECT id, username, password, email, create_at, is_active, is_staff, is_admin FROM users WHERE username=$1 and is_active=true LIMIT 1 OFFSET 0
`

func (q *Queries) GetOneUSer(ctx context.Context, username string) (User, error) {
	row := q.queryRow(ctx, q.getOneUSerStmt, getOneUSer, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.CreateAt,
		&i.IsActive,
		&i.IsStaff,
		&i.IsAdmin,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, password, email, create_at, is_active, is_staff, is_admin FROM users LIMIT $1 OFFSET $2
`

type ListUsersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Password,
			&i.Email,
			&i.CreateAt,
			&i.IsActive,
			&i.IsStaff,
			&i.IsAdmin,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
