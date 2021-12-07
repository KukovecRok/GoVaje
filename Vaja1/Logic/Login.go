package Logic

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
	"todorokvaja1/DataStructures"
)

func (c *Controller) InsertUser(ctx context.Context, user DataStructures.User) (err error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 13)

	if err != nil {
		return
	}
	user.Password = string(hashed)

	return c.db.InsertUser(ctx, user)

}

func (c *Controller) GetUserByName(ctx context.Context, username string) (user DataStructures.User, err error) {

	return c.db.GetUserByName(ctx, username)
}
func (c *Controller) Login(ctx context.Context, userLogin DataStructures.UserLogin) (tokenString string, err error) {

	user, err := c.db.GetUserByName(ctx, userLogin.Username)
	if err != nil {
		return
	}

	err = CheckPasswordHash(userLogin.Password, user.Password)
	if err != nil {
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     user.Id.Hex(),
		"valid_until": time.Now().Add(30 * time.Minute),
	})
	tokenString, err = token.SignedString(c.secret)
	return
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
