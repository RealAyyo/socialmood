package jwt

import (
	"errors"
	"fmt"
	"os"
	"socialmood/internal/config"
	"strconv"
	"strings"
	"time"

	jwtValidator "github.com/dgrijalva/jwt-go"
)

var (
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrJwtSecretIsEmpty        = errors.New("jwt secret is empty")
	ErrInvalidToken            = errors.New("invalid token")
)

type JWT struct {
	secret             string
	accessTokenExpired string
}

type Tokens struct {
	AccessToken string `json:"access_token"`
}

func New(conf *config.JWTConf) *JWT {
	return &JWT{
		secret:             os.Getenv("JWT_SECRET"),
		accessTokenExpired: conf.AccessTokenExpired,
	}
}

func (j *JWT) GenerateToken(userID string) (*Tokens, error) {
	accessToken := jwtValidator.New(jwtValidator.SigningMethodHS256)
	claims := accessToken.Claims.(jwtValidator.MapClaims)

	parsedExp, err := parseDuration(j.accessTokenExpired)
	if err != nil {
		return nil, err
	}
	exp := time.Now().Add(parsedExp)

	claims["exp"] = exp.Unix()
	claims["sub"] = userID
	claims["iat "] = time.Now().Unix()
	claims["iss"] = "SocialMood"
	claims["aud"] = "SocialMoodAPP"

	accessTokenString, err := accessToken.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	return &Tokens{AccessToken: accessTokenString}, nil
}

func (j *JWT) ParseToken(token string) (string, error) {
	tkn, err := jwtValidator.Parse(token, func(token *jwtValidator.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtValidator.SigningMethodHMAC); !ok {
			return "", ErrUnexpectedSigningMethod
		}

		if j.secret == "" {
			return "", ErrJwtSecretIsEmpty
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return "", err
	}

	if !tkn.Valid {
		return "", ErrInvalidToken
	}

	claims, ok := tkn.Claims.(jwtValidator.MapClaims)
	if !ok {
		return "", ErrInvalidToken
	}

	userID := claims["sub"].(string)
	return userID, nil
}

func parseDuration(s string) (time.Duration, error) {
	units := map[string]time.Duration{
		"h": time.Hour,
		"m": time.Minute,
		"s": time.Second,
	}

	parts := strings.SplitN(s, "", 2)
	if len(parts) < 2 {
		return 0, fmt.Errorf("unknown format: %s", s)
	}

	number, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("error parse int: %v", err)
	}

	unit, ok := units[string(s[len(s)-1])]
	if !ok {
		return 0, fmt.Errorf("unknown time unit: %s", string(s[len(s)-1]))
	}

	return time.Duration(number) * unit, nil
}
