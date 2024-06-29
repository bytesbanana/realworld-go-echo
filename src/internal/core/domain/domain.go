package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID             int     `db:"id"`
	Email          string  `db:"email"`
	Username       string  `db:"username"`
	HashedPassword string  `db:"hashed_password"`
	Bio            *string `db:"bio"`
	Image          *string `db:"image"`
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	return err == nil
}
