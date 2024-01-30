import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { Observable } from "rxjs";
import { User } from '../models/user.model';
import { environment } from 'src/environments/environment';
import { PotvrdaSmrti } from '../models/potvrdaSmrti.model';
import { AddPotvrdaSmrti } from '../dto/addPotvrdaSmrti';

@Injectable({
    providedIn: 'root'
})
export class RegistrarService {

    private url = "registrar";

    constructor(
        private http: HttpClient
    ) { }

    public DoctorCreateUser(user: User): Observable<User> {
        return this.http.post<User>(
            `${environment.baseApiUrl}/${this.url}/doctorCreateUser`, user
        );
    }

    public ParentCreateUser(user: User): Observable<User> {
        return this.http.post<User>(
            `${environment.baseApiUrl}/${this.url}/parentCreateUser`, user
        );
    }

    public GetUserJMBG(jmbg: string): Observable<User> {
        return this.http.get<User>(
            `${environment.baseApiUrl}/${this.url}/getUserJMBG/` + jmbg
        );
    }

    public GetUserID(id: string): Observable<User> {
        return this.http.get<User>(
            `${environment.baseApiUrl}/${this.url}/getUserID/` + id
        );
    }

    public GetNewbornsByMotherJMBG(jmbg: string): Observable<User[]> {
        return this.http.get<User[]>(
            `${environment.baseApiUrl}/${this.url}/getNewbornsByMotherJMBG/` + jmbg
        );
    }

    public PostPotvrdaSmrti(potvrda: AddPotvrdaSmrti): Observable<AddPotvrdaSmrti> {
        return this.http.post<AddPotvrdaSmrti>(
            `${environment.baseApiUrl}/${this.url}/postPotvrdaSmrti`, potvrda
        );
    }

    public IsPotvrdaExistJMBG(jmbg: string): Observable<boolean> {
        return this.http.get<boolean>(
            `${environment.baseApiUrl}/${this.url}/isPotvrdaExist/` + jmbg
        );
    }

    public GetPotvrdaSmrtiJMBG(jmbg: string): Observable<PotvrdaSmrti> {
        return this.http.get<PotvrdaSmrti>(
            `${environment.baseApiUrl}/${this.url}/getPotvrdaSmrtiJMBG/` + jmbg
        );
    }

}