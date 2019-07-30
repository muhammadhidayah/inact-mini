package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	_middleware "github.com/muhammadhidayah/inact-mini/middleware"
	"github.com/muhammadhidayah/inact-mini/models"
	"github.com/muhammadhidayah/inact-mini/users"
)

type UserHandler struct {
	ArticleUC users.Usecase
}

func NewUserHandler(mux *_middleware.DefaultMiddleware, ua users.Usecase) {
	handler := &UserHandler{
		ArticleUC: ua,
	}

	middleware := _middleware.InitMidleware()

	mux.HandleFunc("/user", middleware.ApplyMiddleware(handler.FetchUser, middleware.Method("GET")))
	mux.HandleFunc("/postuser", middleware.ApplyMiddleware(handler.Store, middleware.Method("POST")))
	mux.HandleFunc("/login", middleware.ApplyMiddleware(handler.Login, middleware.Method("POST")))
}

func (uh *UserHandler) FetchUser(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		id64, _ := strconv.Atoi(id)
		res, err := uh.ArticleUC.GetUserById(int64(id64))
		if err != nil {
			OutputJSON(w, err)
			return
		}

		OutputJSON(w, res)
	}
}

func (uh *UserHandler) Store(w http.ResponseWriter, r *http.Request) {
	dataJson := json.NewDecoder(r.Body)
	dataPayload := models.Users{}

	if err := dataJson.Decode(&dataPayload); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	lastID, err := uh.ArticleUC.InsertUser(&dataPayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	msg := fmt.Sprintf("insert user success, id %d", lastID)

	w.Write([]byte(msg))

}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	username, password, isOk := r.BasicAuth()

	if !isOk {
		http.Error(w, "Invalid username and password", http.StatusBadRequest)
		return
	}

	isOk, userInfo := uh.authenticationUser(username, password)

	if !isOk {
		http.Error(w, "Invalid username and password", http.StatusBadRequest)
		return
	}

	claims := _middleware.MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    _middleware.APPLICATION_NAME,
			ExpiresAt: time.Now().Add(_middleware.LOGIN_EXPIRATION_DURATION).Unix(),
		},
		ID:       userInfo["id"].(string),
		Username: userInfo["username"].(string),
	}

	token := jwt.NewWithClaims(
		_middleware.JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString(_middleware.JWT_SIGNATURE_KEY)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString, _ := json.Marshal(_middleware.M{"token": signedToken})
	w.Write([]byte(tokenString))
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.Write([]byte("\n"))
}

func (uh *UserHandler) authenticationUser(username string, password string) (bool, _middleware.M) {
	getUser, err := uh.ArticleUC.GetUserByUsername(username)
	if err != nil {
		return false, nil
	}

	if getUser.Username == username && getUser.Password == password {
		data := _middleware.M{
			"id":       strconv.Itoa(int(getUser.ID)),
			"username": getUser.Username,
		}

		return true, data
	}

	return false, nil
}
