package token

import (
	"errors"
	// "log"
	"time"

	pb "github.com/Mubinabd/restaurant_api-gateway/genproto"
	"github.com/golang-jwt/jwt"
)

const (
	signingKey = "Secret key for restaurant reservation service"
)


func GenereteJWTToken(user_id string, username string) (*pb.Token,error) {

	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["user_id"] = user_id
	claims["username"] = username
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(60 * time.Minute).Unix()
	access, err := accessToken.SignedString([]byte(signingKey))
	if err != nil {
		return nil,errors.New("error while genereting access token : " + err.Error())
	}

	rftclaims := refreshToken.Claims.(jwt.MapClaims)
	rftclaims["user_id"] = user_id
	rftclaims["username"] = username
	rftclaims["iat"] = time.Now().Unix()
	rftclaims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	refresh, err := refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		return nil,errors.New("error while genereting refresh token : " + err.Error())
	}
	return &pb.Token{
        AccessToken: access,
        RefreshToken: refresh,
    },nil
	}



func ValidateToken(token string) (bool, error) {
	_, err := ExtractClaim(token)
	if err != nil {
		return false, err
	}
	// if claims["username"] == "admin" {
	// 	return true, nil
	// }

	return true,nil
}

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	}
	token, err = jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
