package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

var vendedores = []Vendedor{
	{ID: 1, Nome: "Gabrielli"},
	{ID: 2, Nome: "Matheus"},
	{ID: 3, Nome: "Wagner"},
	{ID: 4, Nome: "Michelly"},
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) ListarVendedores() []Vendedor {
	return vendedores
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) AdicionarRegistro(data string, orcRec, orcSemRec, solar, ajuste, entrega, assistencia, vendedorID int, anotacoes string) error {
	return AdicionarRegistro(data, orcRec, orcSemRec, solar, ajuste, entrega, assistencia, vendedorID, anotacoes)
}

func (a *App) ListarRegistros() []map[string]interface{} {
	var registros []Registro
	DB.Preload("Vendedor").Find(&registros)

	var resultado []map[string]interface{}

	for _, r := range registros {
		resultado = append(resultado, map[string]interface{}{
			"ID":                  r.ID,
			"Data":                r.Data,
			"OrcamentoReceita":    r.OrcamentoReceita,
			"OrcamentoSemReceita": r.OrcamentoSemReceita,
			"OculosSolar":         r.OculosSolar,
			"Ajuste":              r.Ajuste,
			"Entrega":             r.Entrega,
			"Assistencia":         r.Assistencia,
			"VendedorID":          r.VendedorID,
			"VendedorNome":        r.Vendedor.Nome,
			"Anotacoes":           r.Anotacoes,
		})
	}

	return resultado
}
