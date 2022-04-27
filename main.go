package main

import (
	"Reddit/pkg/heandlears"
	"Reddit/pkg/repositories"
	"Reddit/pkg/services"
	"log"
)

func main() {
	serverInstance := new(Server)
	postgresConfing := repositories.Confing{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "12",
		DBName:   "Reddit",
		SSLMode:  "disable",
	}
	database, err := repositories.NewPostgresDB(postgresConfing)
	if err != nil {
		log.Fatal(err.Error())
	}
	repos := repositories.NewRepositories(database)
	servicesis := services.NewService(repos)
	handlear := heandlears.NewHandler(servicesis)
	runServer(serverInstance, handlear)

}
func runServer(serverInstance *Server, handlear *heandlears.Handler) {
	port := "8080"
	router := handlear.GetRouter()
	if err := serverInstance.Run(port, router); err != nil {
		log.Fatal(err.Error())
	}
	log.Print("server started")
}
