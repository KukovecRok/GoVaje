package MongoDB

import (
	"context"
	"errors"
	"github.com/getsentry/sentry-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
	"todorokvaja1/DataStructures"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (dbo *MongoDB) InsertUser(ctx context.Context, user DataStructures.User) (err error) {

	user.Password, err = HashPassword(user.Password)
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	var uporabnik DataStructures.User
	err = dbo.Client.Database(dbo.Database).Collection("user").FindOne(ctx, bson.M{"username": user.Username}).Decode(&uporabnik)
	if err == nil {
		log.Printf("Najdel userja s tem nicknameom %s", err)
		sentry.CaptureException(err)
		return
	} else {
		log.Printf("Poteka registracija novega uporabnika")
	}

	_, err = dbo.Client.Database(dbo.Database).Collection("user").InsertOne(ctx, user)
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	return
}
func (dbo *MongoDB) GetUserByName(ctx context.Context, username string) (user DataStructures.User, err error) {

	err = dbo.Client.Database(dbo.Database).Collection("user").FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return
	}

	return
}
func (dbo *MongoDB) Login(ctx context.Context, user DataStructures.User) (pravilno bool, err error) {

	gesloZaPrimerjavo := user.Password
	err = dbo.Client.Database(dbo.Database).Collection("user").FindOne(ctx, bson.M{"username": user.Username}).Decode(&user)

	if err != nil {
		sentry.CaptureException(err)
		return
	}

	pravilno = CheckPasswordHash(gesloZaPrimerjavo, user.Password)
	if pravilno {
		log.Printf("%s Pravo geslo", user.Username)
	} else {
		log.Printf("%s Napacno geslo", user.Username)
		err = errors.New("Unauthorized - napacen username/geslo")
		sentry.CaptureException(err)
		return
	}

	return
}
