package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"post/internal/app/models"
	responseUseCase "post/internal/app/response/usecase"
	"post/pkg/Error"
	"post/pkg/httputils"
	"strconv"
)


const (
	ctxKeySession uint8 = 3
	ctxKeyReqID   uint8 = 1
	ctxUserInfo   uint8 = 2
)

type Handlers struct {
	useCase responseUseCase.UseCase
}

func NewHandler(useCase responseUseCase.UseCase) *Handlers {
	return &Handlers{
		useCase: useCase,
	}
}

func (h *Handlers) CreatePostResponse(w http.ResponseWriter, r *http.Request) {
	reqID := r.Context().Value(ctxKeyReqID).(uint64)
	response := &models.Response{}
	if err := json.NewDecoder(r.Body).Decode(response); err != nil {
		httputils.RespondError(w, reqID, err, http.StatusInternalServerError)
		return
	}
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		httputils.RespondError(w, reqID, err, http.StatusInternalServerError)
		return
	}
	response.PostID = id
	response.VacancyResponse = r.URL.String() == "/api/vacancy/"+strconv.FormatUint(id, 10)+"/response"
	response.OrderResponse = r.URL.String() == "/api/order/"+strconv.FormatUint(id, 10)+"/response"
	response, err = h.useCase.Create(*response)
	if err != nil {
		httpErr := &Error.Error{}
		errors.As(err, &httpErr)
		if httpErr.InternalError {
			httputils.RespondError(w, reqID, err, http.StatusInternalServerError)
		} else {
			httputils.RespondError(w, reqID, err, http.StatusBadRequest)
		}
		return
	}
	httputils.Respond(w, reqID, http.StatusCreated, response)
}

func (h *Handlers) GetAllPostResponses(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reqID := r.Context().Value(ctxKeyReqID).(uint64)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		httputils.RespondError(w, reqID, err, http.StatusInternalServerError)
		return
	}
	vacancyResponse := r.URL.String() == "/api/vacancy/"+strconv.FormatUint(id, 10)+"/response"
	orderResponse := r.URL.String() == "/api/order/"+strconv.FormatUint(id, 10)+"/response"
	responses, err := h.useCase.FindByPostID(id, orderResponse, vacancyResponse)
	if err != nil {
		httpErr := &Error.Error{}
		errors.As(err, &httpErr)
		if httpErr.InternalError {
			httputils.RespondError(w, reqID, err, http.StatusInternalServerError)
		} else {
			httputils.RespondError(w, reqID, err, http.StatusBadRequest)
		}
		return
	}

	httputils.Respond(w, reqID, http.StatusOK, responses)
}

func (h *Handlers) ChangePostResponse(w http.ResponseWriter, r *http.Request) {
	response := &models.Response{}
	reqID := r.Context().Value(ctxKeyReqID).(uint64)

	if err := json.NewDecoder(r.Body).Decode(response); err != nil {
		httputils.RespondError(w, reqID, err, http.StatusInternalServerError)
		return
	}
	params := mux.Vars(r)
	var err error
	response.PostID, err = strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		httputils.RespondError(w, reqID, err, http.StatusInternalServerError)
		return
	}
	response.UserID = r.Context().Value(ctxUserInfo).(uint64)
	response.VacancyResponse = r.URL.String() == "/api/vacancy/"+strconv.FormatUint(response.PostID, 10)+"/response"
	response.OrderResponse = r.URL.String() == "/api/order/"+strconv.FormatUint(response.PostID, 10)+"/response"
	responses, err := h.useCase.Change(*response)

	if err != nil {
		httpErr := &Error.Error{}
		errors.As(err, &httpErr)
		if httpErr.InternalError {
			httputils.RespondError(w, reqID, err, http.StatusInternalServerError)
		} else {
			httputils.RespondError(w, reqID, err, http.StatusBadRequest)
		}
		return
	}

	httputils.Respond(w, reqID, http.StatusOK, responses)
}

func (h *Handlers) DelPostResponse(w http.ResponseWriter, r *http.Request) {
	response := &models.Response{}
	params := mux.Vars(r)
	reqID := r.Context().Value(ctxKeyReqID).(uint64)

	var err error
	response.PostID, err = strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		httputils.RespondError(w, reqID, err, http.StatusInternalServerError)
		return
	}

	response.UserID = r.Context().Value(ctxUserInfo).(uint64)
	response.VacancyResponse = r.URL.String() == "/api/vacancy/"+strconv.FormatUint(response.PostID, 10)+"/response"
	response.OrderResponse = r.URL.String() == "/api/order/"+strconv.FormatUint(response.PostID, 10)+"/response"
	err = h.useCase.Delete(*response)

	if err != nil {
		httpErr := &Error.Error{}
		errors.As(err, &httpErr)
		if httpErr.InternalError {
			httputils.RespondError(w, reqID, err, http.StatusInternalServerError)
		} else {
			httputils.RespondError(w, reqID, err, http.StatusBadRequest)
		}
		return
	}
	var emptyInterface interface{}

	httputils.Respond(w, reqID, http.StatusOK, emptyInterface)
}
