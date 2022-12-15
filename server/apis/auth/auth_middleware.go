package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

const sessionKey = "auth"
const sessionVerifyTime = 12 * 3600
const TokenExpires = 24 * time.Hour * 90

const secretKey = "devlog_jjang"

type SessionAuthInfo struct {
	UserID string `json:"user_id"`
	//Auth   interface{} `json:"auth"`
}

type AdminToken struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func CreateSession(c echo.Context, userID string) (SessionAuthInfo, error) {
	auth := SessionAuthInfo{
		UserID: userID,
	}
	sess, err := session.Get(sessionKey, c)
	if err != nil {
		fmt.Errorf("session get failure : %s", err)
		return auth, err
	}
	logrus.Debugf("session get : %+v", sess)

	sess.Options = &sessions.Options{
		Path:   "/api/v1",
		Domain: ".yjproject.blog:3000",
		//Domain:   "localhost:3000",
		MaxAge:   sessionVerifyTime,
		Secure:   true,
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
	}

	//b, _ := json.Marshal(&auth)
	sess.Values["auth"] = userID
	if err = sess.Save(c.Request(), c.Response()); err != nil {
		logrus.WithError(err).Error("session save failed.")
		return auth, err
	}

	fmt.Println("session : ", sess)

	logrus.Debugf("session : %+v", sess)

	//auth.Auth = sess.Values["auth"]
	return auth, nil
}

func CreateJwtToken(userId string) (string, int64) {
	admin := &AdminToken{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpires).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, admin)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		logrus.WithError(err).Errorf("token signed failure - admin:%v", admin)
	}
	return signedToken, admin.StandardClaims.ExpiresAt
}

func MiddlewareJWT() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Request().RequestURI, "/api/v1/admin")
		},
		SigningKey: []byte(secretKey),
		Claims:     &AdminToken{},
	}
	return middleware.JWTWithConfig(config)
}
