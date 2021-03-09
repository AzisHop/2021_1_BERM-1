package apiserver

import (
	"encoding/json"
	"errors"
	"fl_ru/model"
	"fl_ru/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)


const(
	cookiesSalt = "ajsh468Slasdl*6%%8"
)


type server struct{
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store: store,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP (w http.ResponseWriter, r *http.Request){
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter(){
	s.router.HandleFunc("/signup",  s.handleSignUp()).Methods("POST")
	s.router.HandleFunc("/signin",  s.handleSignIn()).Methods("POST")
	s.router.HandleFunc("/profile/{id:[0-9]+}",  s.authenticateUser(s.handleCreateProfile())).Methods("POST")

}

func (s *server) handleSignUp() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		u := &model.User{}
		if err := json.NewDecoder(r.Body).Decode(u) ;err != nil{
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := u.Validate(); err != nil{
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := u.BeforeCreate(); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		err := s.store.User().Create(u)
		if  err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u.Sanitize()
		cookies, err := s.createCookies(u)
		if err != nil{
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		for _, cookie := range cookies {
			http.SetCookie(w, &cookie)
		}
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleSignIn() http.HandlerFunc {
	type Request struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request){
		request := &Request{}
		if err := json.NewDecoder(r.Body).Decode(request) ;err != nil{
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			Email: request.Email,
			Password: request.Password,
		}
		if err := s.store.User().FindByEmail(u); err != nil {
			s.error(w, r, http.StatusConflict, err)
			return
		}
		if u.ComparePassword(request.Password) == false{
			s.error(w, r, http.StatusConflict, errors.New("Bad password"))
			return
		}
		u.Sanitize()
		cookies, err := s.createCookies(u)
		if err != nil{
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		for _, cookie := range cookies {
			http.SetCookie(w, &cookie)
		}
		s.respond(w, r, http.StatusAccepted, u)
	}
}
//profile/1
func (s *server) handleCreateProfile() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request){
		//Дернуть id из контекста
		w.WriteHeader(http.StatusOK)
	}
}

func (s *server) authenticateUser(next http.Handler) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := mux.Vars(r)["id"]
			cookie, err := r.Cookie("session")
			if err != nil {
				s.error(w, r, http.StatusUnauthorized, errors.New("Unauthorized"))
				return
			}
			session := &model.Session{
				SessionId: cookie.Value,
			}
			if err = s.store.Session().Find(session); err != nil {
				s.error(w, r, http.StatusUnauthorized, errors.New("Unauthorized"))
				return
			}
			if id != strconv.FormatUint(session.UserId, 10) {
				s.error(w, r, http.StatusUnauthorized, errors.New("Bad id"))
				return
			}
			//next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "id", s)))
	})

}



func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error){
	s.respond(w, r, code, map[string]string{"error" : err.Error()})
}

func (s* server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}){
	w.WriteHeader(code)
	if data != nil{
		_ = json.NewEncoder(w).Encode(data)

	}
}

func (s *server)createCookies(u *model.User) ([]http.Cookie, error){

	session := &model.Session{
		SessionId: u.Email + time.Now().String(),
		UserId: u.Id,
	}
	session.BeforeCreate()
	if err := s.store.Session().Create(session); err != nil{
		return nil, err
	}
	cookie := http.Cookie{
		Name: "session",
		Value: session.SessionId,
	}
	cookies := []http.Cookie{
		cookie,
		{
			Name:  "id",
			Value: strconv.FormatUint(u.Id, 10),
		},
		{
			Name: "executor",
			Value: strconv.FormatBool(u.Executor),
		},
	}
	return cookies, nil
}
