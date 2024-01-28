export class AddPotvrdaSmrti {
    datumSmrti: number = 0
    mestoSmrti: string = ""
    jmbg: string = ""

    AddPotvrdaSmrti(datumSmrti: number, mestoSmrti: string, jmbg: string) {
        this.datumSmrti = datumSmrti
        this.mestoSmrti = mestoSmrti
        this.jmbg = jmbg
    }
}