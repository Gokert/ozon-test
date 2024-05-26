package middleware

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	utils "ozon-test/pkg"
	"ozon-test/pkg/models"
	httpResponse "ozon-test/pkg/response"
	"time"
)

type contextKey string

const UserIDKey contextKey = "userId"
const UserRoleKey contextKey = "userRole"

type Core interface {
	GetUserId(ctx context.Context, sid string) (uint64, error)
	GetRole(ctx context.Context, userId uint64) (string, error)
}

func AuthCheck(next http.Handler, core Core, lg *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session_id")
		if errors.Is(err, http.ErrNoCookie) {
			next.ServeHTTP(w, r)
			return
		}

		userId, err := core.GetUserId(r.Context(), session.Value)
		if err != nil {
			lg.Errorf("auth check error: %s", err.Error())
			next.ServeHTTP(w, r)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), UserIDKey, userId))
		if userId == 0 {
			next.ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func MethodCheck(next http.Handler, method string, lg *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := &models.Response{Status: http.StatusMethodNotAllowed, Body: nil}
		timeNow := time.Now()

		if r.Method != method {
			httpResponse.SendLog(http.StatusMethodNotAllowed, utils.GetPost, timeNow, lg)
			httpResponse.SendResponse(w, r, response, timeNow, lg)
			return
		}
		next.ServeHTTP(w, r)
	})
}
