package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"gopkg.in/mgo.v2"
	"./services"
	"./daos"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	session, err := mgo.Dial("server1.example.com,server2.example.com")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return origin == "http://my.other.host"
		},
		AllowedMethods: []string{"GET", "POST", "PUT"},
		AllowedHeaders: []string{
			"Accept", "Content-Type", "X-Custom-Header", "Origin"},
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})

	userDAO := daos.NewUserDAO()
	router, err := rest.MakeRouter(
		rest.Get("/user", services.NewUserService(userDAO)),
		rest.Post("/user", services.PostUser(userDAO)),
	)

	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}