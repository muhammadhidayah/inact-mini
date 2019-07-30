package main

import (
	"log"
	"net/http"

	_middleware "github.com/muhammadhidayah/inact-mini/middleware"
	_userHttpDeliver "github.com/muhammadhidayah/inact-mini/users/delivery/http"
	_userRepo "github.com/muhammadhidayah/inact-mini/users/repository"
	_userUsecase "github.com/muhammadhidayah/inact-mini/users/usecase"
)

func main() {
	setting := NewSetting()
	dbConn, err := setting.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := _userRepo.NewPgUsersRepository(dbConn)

	userUA := _userUsecase.NewUserUsecase(userRepo)

	customMux := new(_middleware.DefaultMiddleware)
	customMux.RegisterMiddlewareDefault(customMux.CORS)
	customMux.RegisterMiddlewareDefault(_middleware.MiddlewareJWTAuth)

	_userHttpDeliver.NewUserHandler(customMux, userUA)

	server := new(http.Server)
	server.Addr = ":9001"
	server.Handler = customMux
	server.ListenAndServe()

}
