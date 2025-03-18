export namespace main {
	
	export class Vendedor {
	    ID: number;
	    Nome: string;
	
	    static createFrom(source: any = {}) {
	        return new Vendedor(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Nome = source["Nome"];
	    }
	}

}

