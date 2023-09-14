import { TipVakcine } from "./tipVakcine.model";
import { User } from "./user.model";

export class Pregled {
    id: number = 0
	pocetakPregleda: number = 0
    zavrsetakPregleda: number = 0
    tipVakcine: TipVakcine = new TipVakcine()
    tipPregleda: string = ""
    gradjanin: User = new User;
    lekar: User = new User;


	Pregled(id: number , pocetakPregleda: number , zavrsetakPregleda: number , tipVakcine: TipVakcine , tipPregleda: string , gradjanin: User , lekar: User ) {
		this.id = id;
		this.pocetakPregleda = pocetakPregleda;
		this.zavrsetakPregleda = zavrsetakPregleda;
		this.tipVakcine = tipVakcine;
		this.tipPregleda = tipPregleda;
		this.gradjanin = gradjanin;
		this.lekar = lekar;
	}
}