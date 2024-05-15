package domain

type User struct {
	ID             int     `db:"id"`
	Email          string  `db:"email"`
	Username       string  `db:"username"`
	HashedPassword string  `db:"hashed_password"`
	Bio            *string `db:"bio"`
	Image          *string `db:"image"`
}
