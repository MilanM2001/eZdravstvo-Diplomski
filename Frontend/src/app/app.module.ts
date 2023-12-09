import { NgModule } from '@angular/core';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { BrowserModule } from '@angular/platform-browser';
import { MatToolbarModule } from '@angular/material/toolbar';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './components/header/header.component';
import { AuthInterceptor } from './services/auth.interceptor';
import { RegisterComponent } from './components/auth/register/register.component';
import { LoginComponent } from './components/auth/login/login.component';
import { WelcomeComponent } from './components/welcome/welcome.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatCardModule } from '@angular/material/card';
import { MatSelectModule } from '@angular/material/select';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from "@angular/material/form-field";
import { MatIconModule } from '@angular/material/icon'
import {MatDatepickerModule} from '@angular/material/datepicker';
import {MatNativeDateModule} from '@angular/material/core';
import { MarriageComponent } from './components/eMaticar/marriage/marriage.component';
import { ChooseServiceComponent } from './components/auth/choose-service/choose-service.component';
import { RegularOrAdminComponent } from './components/auth/regular-or-admin/regular-or-admin.component'; 
import {MatSnackBarModule} from "@angular/material/snack-bar";
import { AddPersonRegistryComponent } from './components/eZdravstvo/add-person-registry/add-person-registry.component'; 
import { ViewMyRegistryComponent } from './components/eMaticar/view-my-registry/view-my-registry.component';
import { ZdravstvenoStanjeItemComponent } from './components/eZdravstvo/zdravstveno-stanje/zdravstveno-stanje-item/zdravstveno-stanje-item.component';
import { ZdravstvenoStanjeListComponent } from './components/eZdravstvo/zdravstveno-stanje/zdravstveno-stanje-list/zdravstveno-stanje-list.component';
import { ZdravstvenoStanjeAddComponent } from './components/eZdravstvo/zdravstveno-stanje/zdravstveno-stanje-add/zdravstveno-stanje-add.component';
import { ZdravstvenaStanjaDoctorComponent } from './components/eZdravstvo/zdravstveno-stanje/zdravstvena-stanja-doctor/zdravstvena-stanja-doctor.component';
import { ZdravstvenoStanjeViewMyComponent } from './components/eZdravstvo/zdravstveno-stanje/zdravstveno-stanje-view-my/zdravstveno-stanje-view-my.component';
import { UserDiedComponent } from './components/eZdravstvo/user-died/user-died.component';
import { PregledAddComponent } from './components/eZdravstvo/pregled/pregled-add/pregled-add.component';
import { PregledItemComponent } from './components/eZdravstvo/pregled/pregled-item/pregled-item.component';
import { PregledListComponent } from './components/eZdravstvo/pregled/pregled-list/pregled-list.component';
import { PregledViewComponent } from './components/eZdravstvo/pregled/pregled-view/pregled-view.component';
import { PreglediLekarComponent } from './components/eZdravstvo/pregled/pregledi-lekar/pregledi-lekar.component';
import { PreglediGradjaninComponent } from './components/eZdravstvo/pregled/pregledi-gradjanin/pregledi-gradjanin.component';
import { VakcinaAddComponent } from './components/eZdravstvo/vakcina/vakcina-add/vakcina-add.component';
import { VakcineComponent } from './components/eZdravstvo/vakcina/vakcine/vakcine.component';
import { VakcinaViewComponent } from './components/eZdravstvo/vakcina/vakcina-view/vakcina-view.component';
import { InvaliditetiAdminComponent } from './components/eAdmin/invaliditet/invaliditeti-admin/invaliditeti-admin.component';
import { InvaliditetItemComponent } from './components/eAdmin/invaliditet/invaliditet-item/invaliditet-item.component';
import { InvaliditetListComponent } from './components/eAdmin/invaliditet/invaliditet-list/invaliditet-list.component';
import { InvaliditetViewComponent } from './components/eAdmin/invaliditet/invaliditet-view/invaliditet-view.component';
import { InvaliditetAddComponent } from './components/eAdmin/invaliditet/invaliditet-add/invaliditet-add.component';
import { AlergijaAddComponent } from './components/eAdmin/alergija/alergija-add/alergija-add.component';
import { AlergijaItemComponent } from './components/eAdmin/alergija/alergija-item/alergija-item.component';
import { AlergijaListComponent } from './components/eAdmin/alergija/alergija-list/alergija-list.component';
import { AlergijaViewComponent } from './components/eAdmin/alergija/alergija-view/alergija-view.component';
import { AlergijeAdminComponent } from './components/eAdmin/alergija/alergije-admin/alergije-admin.component';
import { KartoniComponent } from './components/eZdravstvo/karton/kartoni/kartoni.component';
import { KartonViewComponent } from './components/eZdravstvo/karton/karton-view/karton-view.component';
import { KartonItemComponent } from './components/eZdravstvo/karton/karton-item/karton-item.component';
import { KartonListComponent } from './components/eZdravstvo/karton/karton-list/karton-list.component';
import { KartonEditAlergijeComponent } from './components/eZdravstvo/karton/karton-edit-alergije/karton-edit-alergije.component';
import { KartonEditInvaliditetiComponent } from './components/eZdravstvo/karton/karton-edit-invaliditeti/karton-edit-invaliditeti.component';
import { KartonViewMyComponent } from './components/eZdravstvo/karton/karton-view-my/karton-view-my.component';



@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    RegisterComponent,
    LoginComponent,
    WelcomeComponent,
    MarriageComponent,
    ChooseServiceComponent,
    RegularOrAdminComponent,
    AddPersonRegistryComponent,
    ViewMyRegistryComponent,
    ZdravstvenoStanjeItemComponent,
    ZdravstvenoStanjeListComponent,
    ZdravstvenoStanjeAddComponent,
    ZdravstvenaStanjaDoctorComponent,
    ZdravstvenoStanjeViewMyComponent,
    UserDiedComponent,
    PregledAddComponent,
    PregledItemComponent,
    PregledListComponent,
    PregledViewComponent,
    PreglediLekarComponent,
    PreglediGradjaninComponent,
    VakcinaAddComponent,
    VakcineComponent,
    VakcinaViewComponent,
    InvaliditetiAdminComponent,
    InvaliditetItemComponent,
    InvaliditetListComponent,
    InvaliditetViewComponent,
    InvaliditetAddComponent,
    AlergijaAddComponent,
    AlergijaItemComponent,
    AlergijaListComponent,
    AlergijaViewComponent,
    AlergijeAdminComponent,
    KartoniComponent,
    KartonViewComponent,
    KartonItemComponent,
    KartonListComponent,
    KartonEditAlergijeComponent,
    KartonEditInvaliditetiComponent,
    KartonViewMyComponent,
  ],
  imports: [
    BrowserModule,
    MatToolbarModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    MatCardModule,
    MatSelectModule,
    MatButtonModule,
    MatIconModule,
    MatDatepickerModule,
    MatNativeDateModule,
    ReactiveFormsModule,
    BrowserAnimationsModule,
    MatSnackBarModule,
    MatInputModule,
    MatFormFieldModule,
  ],
  providers: [{
    provide: HTTP_INTERCEPTORS,
    useClass: AuthInterceptor,
    multi: true
  }],
  bootstrap: [AppComponent]
})
export class AppModule { }
