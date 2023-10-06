package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Dsn = ""
	Host = ""
	Db_nome = ""
	Db_usuario = ""
	Db_senha = ""
	Porta = ""
)

// Importa as configurações definidas nas variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Host = os.Getenv("HOST")
	Db_nome = os.Getenv("DB_NOME")
	Db_usuario = os.Getenv("DB_USUARIO")
	Db_senha = os.Getenv("DB_SENHA")
	Porta = os.Getenv("API_PORT")

	Dsn = fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%s", Host, Db_nome, Db_usuario, Db_senha, Porta)

}