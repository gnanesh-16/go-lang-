package models

import (
    "database/sql"
    "go-rest-api/config"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func GetUsers() ([]User, error) {
    rows, err := config.DB.Query("SELECT id, name, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    users := []User{}

    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}

func GetUser(id int) (User, error) {
    var user User
    err := config.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            return user, nil
        }
        return user, err
    }
    return user, nil
}

func (u *User) CreateUser() error {
    err := config.DB.QueryRow("INSERT INTO users(name, email) VALUES($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
    if err != nil {
        return err
    }
    return nil
}

func (u *User) UpdateUser() error {
    _, err := config.DB.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", u.Name, u.Email, u.ID)
    if err != nil {
        return err
    }
    return nil
}

func DeleteUser(id int) error {
    _, err := config.DB.Exec("DELETE FROM users WHERE id = $1", id)
    if err != nil {
        return err
    }
    return nil
}
