package token

import (
	"errors"
	"fmt"
	"time"

	redisdao "github.com/CRL-Studio/AuthServer/src/dao/redis"
	rediskeysetdao "github.com/CRL-Studio/AuthServer/src/dao/redis/redisKeySetDao"
	"github.com/CRL-Studio/AuthServer/src/utils/config"
	"github.com/CRL-Studio/AuthServer/src/utils/hash"
	"github.com/CRL-Studio/AuthServer/src/utils/keybuilder"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type claims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}

// AccessToken set claims from parameters and config file and returns access token built
func AccessToken(params map[string]string) (string, error) {
	jti := uuid.NewV4().String()
	now := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		Account: params["account"],
		StandardClaims: jwt.StandardClaims{
			Issuer:    config.Get("jwt.issuer").(string),
			IssuedAt:  now,
			NotBefore: now,
			ExpiresAt: now + int64(config.Get("jwt.access_token_exp").(int)),
			Id:        jti,
		},
	})

	// issues : https://github.com/dgrijalva/jwt-go/issues/65
	// `SignedString` is receiving `[]byte`, but it declares it accepts `interface{}`
	result, err := token.SignedString([]byte(config.Get("jwt.secret").(string)))
	if err != nil {
		return "", err
	}

	redis := redisdao.Redis()
	defer redis.Close()

	key := keybuilder.Jti(params["account"])
	if err = redis.SetEx(key, config.Get("jwt.access_token_exp").(int), jti); err != nil {
		return "", err
	}
	if err := rediskeysetdao.Add(params["account"], key, config.Get("jwt.refresh_token_exp").(int)); err != nil {
		return "", err
	}

	return result, nil
}

// RefreshToken combine user account and timestamp to generate refresh token
func RefreshToken(account string) (string, error) {
	redis := redisdao.Redis()
	defer redis.Close()

	result := hash.New(account + time.Now().String())
	key := keybuilder.RefreshToken(account)

	if err := redis.SetEx(key, config.Get("jwt.refresh_token_exp").(int), result); err != nil {
		return "", err
	}
	if err := rediskeysetdao.Add(account, key, config.Get("jwt.refresh_token_exp").(int)); err != nil {
		return "", err
	}

	return result, nil
}

// Parse validates token gotten from request and returns claims if it's legal
func Parse(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.Get("jwt.secret").(string)), nil
	})

	if err != nil {
		return nil, err
	}

	claims, err := validate(token)

	if err != nil {
		return nil, err
	}
	return claims, err
}

func validate(token *jwt.Token) (map[string]interface{}, error) {
	redis := redisdao.Redis()
	defer redis.Close()

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		if claims["iss"] != config.Get("jwt.issuer") {
			return nil, errors.New("token issuer mismatch")
		}

		key := keybuilder.Jti(claims["account"].(string))
		jti, _ := redis.Get(key)
		if jti != claims["jti"].(string) {
			return nil, errors.New("token id mismatch")
		}

	} else {
		return nil, errors.New("get token claims failed")
	}

	return map[string]interface{}{
		"account": claims["account"],
	}, nil
}
