export class Invaliditet {
    id: number = 0
    naziv: string = ""
    ozbiljnost: string = ""


    Invaliditet(id: number, naziv: string, ozbiljnost: string) {
        this.id = id
        this.naziv = naziv
        this.ozbiljnost = ozbiljnost
    }
}