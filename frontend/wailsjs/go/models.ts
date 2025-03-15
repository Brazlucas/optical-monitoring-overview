export namespace main {
	
	export class Registro {
	    ID: number;
	    Data: string;
	    OrcamentoReceita: number;
	    OrcamentoSemReceita: number;
	    OculosSolar: number;
	    Ajuste: number;
	    Entrega: number;
	    Assistencia: number;
	    VendedorID: number;
	    Anotacoes: string;
	
	    static createFrom(source: any = {}) {
	        return new Registro(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Data = source["Data"];
	        this.OrcamentoReceita = source["OrcamentoReceita"];
	        this.OrcamentoSemReceita = source["OrcamentoSemReceita"];
	        this.OculosSolar = source["OculosSolar"];
	        this.Ajuste = source["Ajuste"];
	        this.Entrega = source["Entrega"];
	        this.Assistencia = source["Assistencia"];
	        this.VendedorID = source["VendedorID"];
	        this.Anotacoes = source["Anotacoes"];
	    }
	}

}

