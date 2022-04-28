package main

import (
	"Reddit/pkg/handler"
	"Reddit/pkg/repository"
	"Reddit/pkg/service"
	"log"
)

func main() {
	serverInstance := new(Server)
	postgresConfing := repository.Confing{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "12",
		DBName:   "Reddit",
		SSLMode:  "disable",
	}
	database, err := repository.NewPostgresDB(postgresConfing)
	if err != nil {
		log.Fatal(err.Error())
	}
	repos := repository.NewRepositories(database)
	servicesis := service.NewService(repos)
	handlear := handler.NewHandler(servicesis)
	runServer(serverInstance, handlear)

}
func runServer(serverInstance *Server, handlear *handler.Handler) {
	port := "8080"
	router := handlear.GetRouter()
	if err := serverInstance.Run(port, router); err != nil {
		log.Fatal(err.Error())
	}
	log.Print("server started")
}
