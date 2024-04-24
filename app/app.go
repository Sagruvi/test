package app

import (
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	_ "testproject/docs"
	"testproject/internal/controller"
)

func Run() {
	con := controller.NewController()
	http.HandleFunc("/api/createevent", con.CreateEvent)
	http.HandleFunc("/api/event", con.GetEvent)
	http.HandleFunc("/api/test", con.Test)
	http.HandleFunc("/api/search", con.Search)
	http.HandleFunc("/swagger/*", httpSwagger.WrapHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", "8080"), nil)
}
