package main

import (
	"context"
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

// NewApp cria uma nova instância do aplicativo
func NewApp() *App {
	return &App{}
}

// ListarVendedores retorna a lista de vendedores
func (a *App) ListarVendedores() []Vendedor {
	return vendedores
}

// startup é chamado quando o app inicia, e o contexto é salvo
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// AdicionarRegistro chama a função para adicionar um registro
func (a *App) AdicionarRegistro(data string, orcRec, orcSemRec, solar, ajuste, entrega, assistencia, vendedorID int, anotacoes string) error {
	return AdicionarRegistro(data, orcRec, orcSemRec, solar, ajuste, entrega, assistencia, vendedorID, anotacoes)
}

// ListarRegistros chama a função para listar os registros
func (a *App) ListarRegistros() []map[string]any {
	return ListarRegistros()
}
