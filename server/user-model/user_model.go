/*
 * Copyright 2019 Oleg Borodin  <borodin@unix7.org>
 */


package userModel

import (
    "log"
    "github.com/jmoiron/sqlx"
    _ "github.com/jackc/pgx/v4/stdlib"
)

const schema = `
    DROP TABLE IF EXISTS users;
    CREATE TABLE IF NOT EXISTS users (
        id          INTEGER PRIMARY KEY,
        username    VARCHAR(255) NOT NULL UNIQUE,
        password    VARCHAR(255) NOT NULL,
        isadmin     BOOLEAN DEFAULT FALSE
    );`

type Model struct {
    db *sqlx.DB
}

type User struct {
    Id          int     `db:"id"        json:"id"`
    Username    string  `db:"username"  json:"username"`
    Password    string  `db:"password"  json:"password,omitempty"`
    IsAdmin     bool    `db:"isadmin"   json:"isadmin,omitempty"`
    Limit       int     `db:"-"         json:"limit,omitempty"`
    Offset      int     `db:"-"         json:"offset,omitempty"`
}

type Page struct {
    Total       int     `json:"total"`
    Offset      int     `json:"offset"`
    Limit       int     `json:"limit"`
    UserPattern string  `json:"user_pattern"`
    Users       *[]User `json:"users,omitempty"`
}

func (this *Model) Migrate() error {
    _, err := this.db.Exec(schema)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

func (this *Model) List(page *Page) (error) {
    var request string
    var err error
    var total int

    userPattern := "%" + page.UserPattern + "%"
    request = `SELECT COUNT(id) as total FROM users WHERE username LIKE $1`
    err = this.db.QueryRow(request, userPattern).Scan(&total)
    if err != nil {
        log.Println(err)
        return err
    }
    page.Total = total

    var users []User
    request = `SELECT id, username, password, isadmin FROM users
                WHERE username LIKE $1
                ORDER BY username LIMIT $2 OFFSET $3`

    err = this.db.Select(&users, request, userPattern, page.Limit, page.Offset)
    if err != nil {
        log.Println(err)
        return err
    }
    page.Users = &users
    return nil
}

func (this *Model) Create(user User) error {
    request := `INSERT INTO users(username, password, isadmin) VALUES ($1, $2, $3)`
    _, err := this.db.Exec(request, user.Username, user.Password, user.IsAdmin)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

func (this *Model) Delete(user User) error {
    request := `DELETE FROM users WHERE id = $1`
    _, err := this.db.Exec(request, user.Id)
    if err != nil {
        log.Println(err)
        return err

    }
    return nil
}

func (this *Model) Find(user User) (User, error) {
    request := `SELECT id, username, password, isadmin FROM users WHERE username = $1 LIMIT 1`
    var out User
    err := this.db.Get(&out, request, user.Username)
    if err != nil {
        log.Println(err)
        return out, err
    }
    return out, nil
}

func (this *Model) Update(user User) error {
    request := `UPDATE users SET username = $1, password = $2, isadmin = $3 WHERE id = $4`
    _, err := this.db.Exec(request, user.Username, user.Password, user.IsAdmin, user.Id)
    if err != nil {
        log.Println(err)
        return err
    }
    return nil
}

func (this *Model) Check(user User) error {
    request := `SELECT * FROM users WHERE username = $1 and password = $2 LIMIT 1`
    err := this.db.Get(&user, request, user.Username, user.Password)
    if err != nil {
        log.Println(err)
    }
    return err
}

func New(db *sqlx.DB) *Model {
    model := Model{
        db: db,
    }
    return &model
}
