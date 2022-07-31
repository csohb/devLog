package common

import (
	"devLog/common/api_context"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	SessionName   = "devLog-session"
	SessionPath   = "/api/v1"
	SessionSecret = "devLog-session-secret"
	MaxAge        = 3600
)

type SessionInfo struct {
	UserID  string
	IsLogin bool
}

func (s *SessionInfo) ParseAuthorization(c echo.Context) *api_context.CommonResponse {
	sess, err := session.Get(SessionName, c)
	if err != nil {
		return api_context.FailureJSON(http.StatusUnauthorized, "세션 정보를 찾을 수 없습니다.")
	}

	sess.Options = &sessions.Options{
		Path:     SessionPath,
		MaxAge:   MaxAge,
		HttpOnly: true,
	}

	if val, has := sess.Values["devlog"]; has == false {
		return api_context.FailureJSON(http.StatusUnauthorized, "세션 정보를 찾을 수 없습니다.")
	} else {
		*s = val.(SessionInfo)
	}

	if err = sess.Save(c.Request(), c.Response()); err != nil {
		return api_context.FailureJSON(http.StatusInternalServerError, "세션 생성 실패")
	}

	return nil
}

func (s SessionInfo) Middleware() echo.MiddlewareFunc {
	return session.Middleware(sessions.NewCookieStore([]byte(SessionSecret)))
}

func (s SessionInfo) CreateSession(c echo.Context) error {
	sess, err := session.Get(SessionName, c)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:     SessionPath,
		MaxAge:   MaxAge,
		HttpOnly: true,
	}

	sess.Values["devlog"] = s

	if err = sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return nil
}
