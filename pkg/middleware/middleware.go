package middleware

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
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
		response := &models.Response{Status: http.StatusUnauthorized, Body: nil}
		timeNow := time.Now()

		session, err := r.Cookie("session_id")
		if errors.Is(err, http.ErrNoCookie) {
			httpResponse.SendResponse(w, r, response, timeNow, lg)
			return
		}

		userId, err := core.GetUserId(r.Context(), session.Value)
		if err != nil {
			lg.Errorf("auth check error: %s", err.Error())
			httpResponse.SendResponse(w, r, response, timeNow, lg)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), UserIDKey, userId))
		if userId == 0 {
			httpResponse.SendResponse(w, r, response, timeNow, lg)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func MethodCheck(next http.Handler, method string, lg *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeNow := time.Now()

		if r.Method != method {
			response := &models.Response{Status: http.StatusMethodNotAllowed, Body: nil}
			httpResponse.SendResponse(w, r, response, timeNow, lg)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetRole(next http.Handler, core Core, lg *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, isAuth := r.Context().Value(UserIDKey).(uint64)
		if !isAuth {
			next.ServeHTTP(w, r)
			return
		}

		result, err := core.GetRole(r.Context(), userId)
		if err != nil {
			lg.Errorf("auth check error: %s", err)
			next.ServeHTTP(w, r)
			return
		}

		if result != "admin" {
			next.ServeHTTP(w, r)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), UserRoleKey, result))

		next.ServeHTTP(w, r)
	})
}

func CheckRole(next http.Handler, core Core, lg *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeNow := time.Now()

		userId, isAuth := r.Context().Value(UserIDKey).(uint64)
		if !isAuth {
			response := &models.Response{Status: http.StatusUnauthorized, Body: nil}
			httpResponse.SendResponse(w, r, response, timeNow, lg)
			return
		}

		result, err := core.GetRole(r.Context(), userId)
		if err != nil {
			lg.Errorf("auth check error: %s", err)
			response := &models.Response{Status: http.StatusUnauthorized, Body: nil}
			httpResponse.SendResponse(w, r, response, timeNow, lg)
			return
		}

		if result != "admin" {
			response := &models.Response{Status: http.StatusForbidden, Body: nil}
			httpResponse.SendResponse(w, r, response, timeNow, lg)
			return
		}

		next.ServeHTTP(w, r)
	})
}
