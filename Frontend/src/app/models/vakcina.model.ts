export class Vakcina {
    id: string = ""
    naziv: string = ""
    kompanija: string = ""


	TipVkacine(id: string, naziv: string , kompanija: string ) {
		this.id = id
		this.naziv = naziv
		this.kompanija = kompanija
	}
    
}