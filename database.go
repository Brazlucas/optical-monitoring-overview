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

	// Inserir vendedores no banco se ainda não existirem
	var count int64
	DB.Model(&Vendedor{}).Count(&count)
	if count == 0 {
		vendedores := []Vendedor{
			{Nome: "Gabrielli"},
			{Nome: "Matheus"},
			{Nome: "Wagner"},
			{Nome: "Michelly"},
		}
		DB.Create(&vendedores) // Insere os vendedores no banco
	}
}

type Registro struct {
	ID                  uint `gorm:"primaryKey"`
	Data                string
	LenteContato        int
	Fluxo               int
	Venda               int
	OrcamentoReceita    int
	OrcamentoSemReceita int
	OculosSolar         int
	Ajuste              int
	Entrega             int
	Assistencia         int
	VendedorID          uint
	Vendedor            Vendedor `gorm:"foreignKey:VendedorID"` // Adicionando relação correta
	Anotacoes           string
}

type Vendedor struct {
	ID   uint `gorm:"primaryKey"`
	Nome string
}
