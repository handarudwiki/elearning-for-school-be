package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/handarudwiki/config"
)

type JWTService interface {
	GenerateAccessToken(id uint, role int) (token string, err error)
	GenrateRefreshToken(id uint, role int) (token string, err error)
	ValidateToken(token string) (claims jwt.MapClaims, err error)
}

type jwtService struct {
	jwt config.JWT
}

func NewJWT(jwt config.JWT) JWTService {
	return &jwtService{jwt}
}

func (s *jwtService) GenerateAccessToken(id uint, role int) (token string, err error) {
	expiredTime := time.Now().Add(time.Hour * 15)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": id,
		"role":   role,
		"exp":    expiredTime.Unix(),
	})

	return claims.SignedString([]byte(s.jwt.AccessSecretKey))
}
func (s *jwtService) GenrateRefreshToken(id uint, role int) (token string, err error) {
	expiredTime := time.Now().Add(time.Hour * 24 * 7)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": id,
		"role":   role,
		"exp":    expiredTime.Unix(),
	})

	return claims.SignedString([]byte(s.jwt.RefreshSecretKey))
}

func (s *jwtService) ValidateToken(token string) (claims jwt.MapClaims, err error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}

		return []byte(s.jwt.AccessSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
