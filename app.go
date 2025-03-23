package main

import (
	"context"
)

type App struct {
	ctx context.Context
}

var vendedores = []Vendedor{
	{ID: 1, Nome: "Gabrielli"},
	{ID: 2, Nome: "Matheus"},
	{ID: 3, Nome: "Wagner"},
	{ID: 4, Nome: "Michelly"},
}

func NewApp() *App {
	return &App{}
}

func (a *App) ListarVendedores() []Vendedor {
	return vendedores
}

// startup é chamado quando o app inicia, e o contexto é salvo
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) AdicionarRegistro(data string, orcRec, orcSemRec, solar, ajuste, entrega, assistencia, vendedorID int, anotacoes string) error {
	return AdicionarRegistro(data, orcRec, orcSemRec, solar, ajuste, entrega, assistencia, vendedorID, anotacoes)
}

func (a *App) ListarRegistros() []map[string]any {
	return ListarRegistros()
}
