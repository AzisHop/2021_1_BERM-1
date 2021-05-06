package errortools

import (
	customError "authorizationservice/pkg/error"
	"net/http"
)

func ErrorHandle(err error) (interface{}, int) {
	if respBody, code, ok := sqlErrorHandle(err); ok {
		return respBody, code
	}
	if respBody, code, ok := validationErrorHandle(err); ok {
		return respBody, code
	}
	if respBody, code, ok := grpcErrorHandle(err); ok {
		return respBody, code
	}
	return map[string]interface{}{
		"message": customError.InternalServerErrorMsg,
	}, http.StatusInternalServerError
}
