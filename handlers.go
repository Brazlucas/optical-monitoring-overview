package main

func AdicionarRegistro(data string, orcRec, orcSemRec, solar, ajuste, entrega, assistencia, vendedorID int, anotacoes string) error {
	registro := Registro{
		Data:                data,
		OrcamentoReceita:    orcRec,
		OrcamentoSemReceita: orcSemRec,
		OculosSolar:         solar,
		Ajuste:              ajuste,
		Entrega:             entrega,
		Assistencia:         assistencia,
		VendedorID:          uint(vendedorID),
		Anotacoes:           anotacoes,
	}
	result := DB.Create(&registro)
	return result.Error
}

func ListarRegistros() []Registro {
	var registros []Registro
	DB.Preload("Vendedor").Find(&registros)
	return registros
}
