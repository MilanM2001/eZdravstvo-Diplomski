import { Vakcina } from '../models/vakcina.model';

export class AddPregled {
  tipPregleda: string = '';
  pocetakPregleda: number = 0;
  zavrsetakPregleda: number = 0;
  vakcinaID: string = '';

  AddPregled(
    tipPregleda: string,
    pocetakPregleda: number,
    zavrsetakPregleda: number,
    vakcinaID: string
  ) {
    this.tipPregleda = tipPregleda;
    this.pocetakPregleda = pocetakPregleda;
    this.zavrsetakPregleda = zavrsetakPregleda;
    this.vakcinaID = vakcinaID;
  }
}
