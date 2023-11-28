export class Invaliditet {
    id: number = 0
    naziv: string = ""
    opis: string = ""
    ozbiljnost: string = ""


    Invaliditet(id: number, naziv: string, opis: string, ozbiljnost: string) {
        this.id = id
        this.naziv = naziv
        this.opis = opis
        this.ozbiljnost = ozbiljnost
    }
}