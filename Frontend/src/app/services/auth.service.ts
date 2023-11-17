import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { environment } from "src/environments/environment";
import {Credentials} from "../models/credentials";

@Injectable({
    providedIn: 'root'
    })
    export class AuthService {
        private url = "auth";
        constructor(private http: HttpClient) { }

    public Registration(credentials: Credentials): Observable<any> {
        return this.http.post(`${environment.baseApiUrl}/${this.url}/registration`, credentials);
    }

    public Login(credentials: Credentials): Observable<any> {
      return this.http.post(`${environment.baseApiUrl}/${this.url}/login`, credentials, {responseType: 'text'});
    }


}
