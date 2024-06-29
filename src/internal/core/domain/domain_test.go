package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUser(t *testing.T) {
	t.Parallel()
	hashPwd, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	t.Run("should pass check password", func(t *testing.T) {
		u := &User{
			HashedPassword: string(hashPwd),
		}

		assert := assert.New(t)
		assert.True(u.CheckPassword("password"))
	})

	t.Run("should fail check password", func(t *testing.T) {
		u := &User{
			HashedPassword: string(hashPwd),
		}

		assert := assert.New(t)
		assert.False(u.CheckPassword("wrongpassword"))
	})

}
