export class Alergija {
    id: number = 0
	naziv: string = ""
    ozbiljnost: string = ""


	Alergija(id: number , naziv: string, ozbiljnost: string ) {
		this.id = id
		this.naziv = naziv
        this.ozbiljnost = ozbiljnost
	}
}