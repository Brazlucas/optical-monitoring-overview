package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("controle_vendas.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	// Criando a estrutura do banco
	DB.AutoMigrate(&Registro{}, &Vendedor{})
}

type Registro struct {
	ID                  uint `gorm:"primaryKey"`
	Data                string
	OrcamentoReceita    int
	OrcamentoSemReceita int
	OculosSolar         int
	Ajuste              int
	Entrega             int
	Assistencia         int
	VendedorID          uint
	Anotacoes           string
}

type Vendedor struct {
	ID   uint `gorm:"primaryKey"`
	Nome string
}
