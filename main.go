package main

import (
	"fmt"
	"net/http"

	_userHttpDeliver "github.com/muhammadhidayah/inact-mini/users/delivery/http"
	_userRepo "github.com/muhammadhidayah/inact-mini/users/repository"
	_userUsecase "github.com/muhammadhidayah/inact-mini/users/usecase"
)

func main() {
	setting := NewSetting()
	dbConn, err := setting.Connect()

	if err != nil {
		fmt.Println("Oooppss Can't Connect to Database")
		return
	}

	userRepo := _userRepo.NewPgUsersRepository(dbConn)

	userUA := _userUsecase.NewUserUsecase(userRepo)

	mux := http.DefaultServeMux
	var handler http.Handler = mux
	_userHttpDeliver.NewUserHandler(mux, userUA)
	server := new(http.Server)
	server.Addr = ":9000"
	server.Handler = handler
	server.ListenAndServe()

}
