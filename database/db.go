package database

import (
	"log"

	"github.com/devopscodeck/Ginutility/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=54322 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
