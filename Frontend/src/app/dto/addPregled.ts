import { Vakcina } from "../models/vakcina.model"

export class AddPregled {
    tipPregleda: string = ""
    pocetakPregleda: number = 0
    zavrsetakPregleda: number = 0
    vakcina: Vakcina = new Vakcina()

    AddPregled(tipPregleda: string, pocetakPregleda: number, zavrsetakPregleda: number, vakcina: Vakcina) {
        this.tipPregleda = tipPregleda
        this.pocetakPregleda = pocetakPregleda
        this.zavrsetakPregleda = zavrsetakPregleda
        this.vakcina = vakcina
    }
    
}