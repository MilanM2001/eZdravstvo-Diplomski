import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { User } from '../models/user.model';
import { ZdravstvenoStanje } from '../models/zdravstvenoStanje.model';
import { Vakcina } from '../models/vakcina.model';
import { AddVakcina } from '../dto/addVakcina';
import { AddPregled } from '../dto/addPregled';
import { Pregled } from '../models/pregled.model';

@Injectable({
  providedIn: 'root',
})
export class HealthcareService {
  private url = 'healthcare';
  constructor(private http: HttpClient) {}

  public GetSveVakcine(): Observable<Vakcina[]> {
    return this.http.get<Vakcina[]>(
      `${environment.baseApiUrl}/${this.url}/getSveVakcine`
    );
  }

  public GetVakcinaID(id: string): Observable<Vakcina> {
    return this.http.get<Vakcina>(
      `${environment.baseApiUrl}/${this.url}/getVakcinaID/` + id
    );
  }

  public GetPregledID(pregled_id: string): Observable<Pregled> {
    return this.http.get<Pregled>(
      `${environment.baseApiUrl}/${this.url}/getPregledID/` + pregled_id
    );
  }

  public GetMojiPreglediLekar(): Observable<Pregled[]> {
    return this.http.get<Pregled[]>(
      `${environment.baseApiUrl}/${this.url}/getMojiPreglediLekar`
    );
  }

  public GetMojiSlobodniPreglediLekar(): Observable<Pregled[]> {
    return this.http.get<Pregled[]>(
      `${environment.baseApiUrl}/${this.url}/getMojiSlobodniPreglediLekar`
    );
  }

  public GetMojiZauzetiPreglediLekar(): Observable<Pregled[]> {
    return this.http.get<Pregled[]>(
      `${environment.baseApiUrl}/${this.url}/getMojiZauzetiPreglediLekar`
    );
  }

  public PostVakcina(vakcina: AddVakcina): Observable<AddVakcina> {
    return this.http.post<AddVakcina>(
      `${environment.baseApiUrl}/${this.url}/postVakcina`,
      vakcina
    );
  }

  public PutVakcina(vakcina: AddVakcina, id: string): Observable<AddVakcina> {
    return this.http.put<AddVakcina>(
      `${environment.baseApiUrl}/${this.url}/putVakcina/` + id,
      vakcina
    );
  }

  public DeleteVakcinaID(id: string) {
    return this.http.delete(
      `${environment.baseApiUrl}/${this.url}/deleteVakcinaID/` + id
    );
  }

  public PostPregled(pregled: AddPregled): Observable<AddPregled> {
    return this.http.post<AddPregled>(
      `${environment.baseApiUrl}/${this.url}/postPregled`,
      pregled
    );
  }

  public SetAppointment(id: string) {
    return this.http.put(
      `${environment.baseApiUrl}/${this.url}/setAppointment/` + id,
      null
    );
  }

  public DeleteAppointment(id: string) {
    return this.http.delete(
      `${environment.baseApiUrl}/${this.url}/deleteAppointmentByID/` + id
    );
  }

  public DeletePregledID(id: string) {
    return this.http.delete(
      `${environment.baseApiUrl}/${this.url}/deletePregledID/` + id
    );
  }

  public GetMe(): Observable<User> {
    return this.http.get<User>(`${environment.baseApiUrl}/${this.url}/getMe`);
  }

  public GetAllZdravstvenaStanja(): Observable<ZdravstvenoStanje[]> {
    return this.http.get<ZdravstvenoStanje[]>(
      `${environment.baseApiUrl}/${this.url}/allZdravstvenaStanja`
    );
  }

  public GetMyZdravstvenoStanje(): Observable<ZdravstvenoStanje> {
    return this.http.get<ZdravstvenoStanje>(
      `${environment.baseApiUrl}/${this.url}/myZdravstvenoStanje`
    );
  }

  public NewZdravstvenoStanje(
    zdravstvenoStanje: ZdravstvenoStanje
  ): Observable<ZdravstvenoStanje> {
    return this.http.post<ZdravstvenoStanje>(
      `${environment.baseApiUrl}/${this.url}/newZdravstvenoStanje`,
      zdravstvenoStanje
    );
  }

  public AddPersonToRegistry(user: User): Observable<User> {
    return this.http.post<User>(
      `${environment.baseApiUrl}/${this.url}/addPersonToRegistry`,
      user
    );
  }
}
