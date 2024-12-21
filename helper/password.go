package helper

import (
 "golang.org/x/crypto/bcrypt"
 "time"
 "github.com/golang-jwt/jwt/v5"
)
var secretKey = []byte("secret-key")

func EncryptPassword(password string) (string, error) {
 bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
 return string(bytes), err
}

func ComparePassword(password, hash string) bool {
 err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
 return err == nil
}



func CreateToken(email string) (string, error) {
 token := jwt.NewWithClaims(jwt.SigningMethodHS256,
  jwt.MapClaims{
   "email": email,
   "exp":   time.Now().Add(time.Hour * 24).Unix(),
  })

 tokenString, err := token.SignedString(secretKey)
 if err != nil {
  return "", err
 }

 return tokenString, nil
}