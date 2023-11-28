export class AddInvaliditet {
    naziv: string = ""
    opis: string = ""
    ozbiljnost: string = ""

    AddInvaliditet(naziv: string, opis: string, ozbiljnost: string) {
        this.naziv = naziv
        this.opis = opis
        this.ozbiljnost = ozbiljnost
    }
}