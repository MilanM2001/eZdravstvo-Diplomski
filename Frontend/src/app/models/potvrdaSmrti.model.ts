export class PotvrdaSmrti {
    id: number = 0
    datumSmrti: number = 0
    mestoSmrti: string = ""
    jmbg: string = ""

    PotvrdaSmrti(id: number, datumSmrti: number, mestoSmrti: string, jmbg: string) {
        this.id = id
        this.datumSmrti = datumSmrti
        this.mestoSmrti = mestoSmrti
        this.jmbg = jmbg
    }
}