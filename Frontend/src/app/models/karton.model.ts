import { Alergija } from "./alergija.model"
import { Invaliditet } from "./invaliditet.model"

export class Karton {
    id: number = 0
    jmbg: string = ""
    alergije: Alergija[] = []
    invaliditeti: Invaliditet[] = []

    Karton(id: number, jmbg: string) {
        this.id = id
        this.jmbg = jmbg
    }
}