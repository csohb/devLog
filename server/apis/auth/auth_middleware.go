package auth

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

const sessionKey = "auth"
const sessionVerifyTime = 12 * 3600

type SessionAuthInfo struct {
	UserID string `json:"user_id"`
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
		Path:     "/api/v1",
		MaxAge:   sessionVerifyTime,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}

	//b, _ := json.Marshal(&auth)
	sess.Values["auth"] = userID
	if err = sess.Save(c.Request(), c.Response()); err != nil {
		logrus.WithError(err).Error("session save failed.")
		return auth, err
	}

	fmt.Println("session : ", sess)

	logrus.Debugf("session : %+v", sess)
	return auth, nil
}
