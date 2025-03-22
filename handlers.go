package main

import "fmt"

// AdicionarRegistro manipula a criação de um novo registro no banco
func AdicionarRegistro(data string, orcRec, orcSemRec, solar, ajuste, entrega, assistencia, vendedorID int, anotacoes string) error {
	// Verificar se o vendedor existe
	var vendedor Vendedor
	result := DB.First(&vendedor, vendedorID)
	if result.Error != nil {
		return fmt.Errorf("vendedor com ID %d não encontrado", vendedorID)
	}

	// Criar um novo registro
	registro := Registro{
		Data:                data,
		OrcamentoReceita:    orcRec,
		OrcamentoSemReceita: orcSemRec,
		OculosSolar:         solar,
		Ajuste:              ajuste,
		Entrega:             entrega,
		Assistencia:         assistencia,
		VendedorID:          uint(vendedorID),
		Vendedor:            vendedor, // Associar corretamente o vendedor
		Anotacoes:           anotacoes,
	}

	// Salvar o registro no banco
	result = DB.Create(&registro)
	return result.Error
}

// ListarRegistros retorna todos os registros do banco
func ListarRegistros() []map[string]interface{} {
	var registros []Registro
	DB.Preload("Vendedor").Find(&registros) // Carregar o vendedor associado

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
			"VendedorNome":        r.Vendedor.Nome, // Nome do vendedor
			"Anotacoes":           r.Anotacoes,
		})
	}

	if len(resultado) == 0 {
		resultado = []map[string]interface{}{}
	}

	return resultado
}
