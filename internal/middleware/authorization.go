package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID int
}

const TOKENEXP = time.Hour * 3
const SECRETKEY = "supersecretkey"

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID, err := getUserIDFromCookie(c)
		if err != nil {
			code := http.StatusBadRequest
			if errors.Is(err, http.ErrNoCookie) {
				code = http.StatusUnauthorized
			}
			contentType := c.Request.Header.Get("Content-Type")
			if contentType == "application/json" {
				var errorMassages []map[string]interface{}
				errorMassage := map[string]interface{}{
					"message": fmt.Sprintf("Unauthorized %s", err),
					"code":    code,
				}
				errorMassages = append(errorMassages, errorMassage)
				answer, _ := json.Marshal(errorMassages)
				c.Data(code, "application/json", answer)
			} else {
				c.String(code, fmt.Sprintf("Unauthorized %s", err))
			}
			c.Abort()
			return
		}
		c.Set("userID", userID)
	}
}

func getUserIDFromCookie(c *gin.Context) (string, error) {
	userID, err := c.Cookie("userID")

	if err != nil {
		if c.Request.RequestURI == "/api/user/urls" {
			return "", err
		}
		var newUserID string
		newUserID, err = BuildJWTString()
		if err != nil {
			return "", err
		}
		c.SetCookie("userID", newUserID, 3600, "/", "localhost", false, true)
		c.Set("userID", newUserID)
		return newUserID, nil
	} else {
		_, err = GetUserId(userID)
		if err != nil {
			return "", err
		}
	}

	return userID, nil
}

func BuildJWTString() (string, error) {
	// создаём новый токен с алгоритмом подписи HS256 и утверждениями — Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// когда создан токен
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKENEXP)),
		},
		// собственное утверждение
		UserID: 2,
	})

	// создаём строку токена
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		return "", err
	}

	// возвращаем строку токена
	return tokenString, nil
}

func GetUserId(tokenString string) (int, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(SECRETKEY), nil
		})
	if err != nil {
		return -1, fmt.Errorf("token is not valid")
	}

	if !token.Valid {
		return -1, fmt.Errorf("token is not valid")
	}

	return claims.UserID, nil
}
