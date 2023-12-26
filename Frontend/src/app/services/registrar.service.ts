import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { Observable } from "rxjs";
import { User } from '../models/user.model';
import { environment } from 'src/environments/environment';

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

}