package postgres

import (
    "database/sql"

    "github.com/Go11Group/at_lesson/lesson35/model"
)

type UserStorage struct {
    Db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
    return &UserStorage{Db: db}
}

func (u *UserStorage) Create(user *model.User) error {
    _, err := u.Db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
    return err
}

func (u *UserStorage) Get(id int) (*model.User, error) {
    var user model.User
    err := u.Db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
    return &user, err
}

func (u *UserStorage) Update(user *model.User) error {
    _, err := u.Db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, user.ID)
    return err
}

func (u *UserStorage) Delete(id int) error {
    _, err := u.Db.Exec("DELETE FROM users WHERE id = $1", id)
    return err
}
