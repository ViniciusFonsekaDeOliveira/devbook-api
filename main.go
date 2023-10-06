package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Carregar()

	fmt.Println(config.Dsn)


	fmt.Println("Rodando API!")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Porta), r))
}