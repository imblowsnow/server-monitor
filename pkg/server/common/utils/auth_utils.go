package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"server-monitor/pkg/server/config"
)

// 生成token
func GenerateToken(data map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(data))
	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(config.Config.Secret))
}

// 解析token
func ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.Config.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("token claims error")
	}
}
