package utils

import (
	"authorizationservice/internal/models"
	"authorizationservice/pkg/Error"
	"authorizationservice/pkg/logger"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const (
	ctxKeyReqID uint8 = 1
)

func Respond(w http.ResponseWriter, requestId uint64, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			RespondError(w, requestId, err, 500)
			return
		}
	}
	logger.LoggingResponse(requestId, code)
}

func RespondError(w http.ResponseWriter, requestId uint64, err error, errorCode int) {
	logger.LoggingError(requestId, err)
	httpError := &Error.Error{}
	if errors.As(err, &httpError) {
		Respond(w, requestId, errorCode, httpError.ErrorDescription)
		return
	}
	Respond(w, requestId, http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
}

func RespondCSRF() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Context().Value(ctxKeyReqID).(uint64)

		logger.LoggingError(reqID, errors.New("Invalid CSRF token"))
		Respond(w, reqID, http.StatusForbidden, map[string]interface{}{
			"error": "Invalid CSRF token",
		})
	})
}

func CreateCookie(session *models.Session, w http.ResponseWriter) {
	cookies := []http.Cookie{
		{
			Name:     "sessionID",
			Value:    session.SessionID,
			Expires:  time.Now().AddDate(0, 1, 0),
			HttpOnly: true,
		},
	}
	for _, cookie := range cookies {
		http.SetCookie(w, &cookie)
	}
}

func RemoveCookies(cookies []*http.Cookie) {
	for i, _ := range cookies {
		cookies[i].Expires = time.Now().AddDate(0, 0, -1)
		cookies[i].HttpOnly = true
	}
}
