package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_middleware "github.com/muhammadhidayah/inact-mini/middleware"
	"github.com/muhammadhidayah/inact-mini/models"
	"github.com/muhammadhidayah/inact-mini/users"
)

type UserHandler struct {
	ArticleUC users.Usecase
}

func NewUserHandler(mux *http.ServeMux, ua users.Usecase) {
	handler := &UserHandler{
		ArticleUC: ua,
	}

	middleware := _middleware.InitMidleware()

	mux.HandleFunc("/user", middleware.ApplyMiddleware(handler.FetchUser, middleware.Method("GET")))
	mux.HandleFunc("/postuser", middleware.ApplyMiddleware(handler.Store, middleware.Method("POST")))
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
