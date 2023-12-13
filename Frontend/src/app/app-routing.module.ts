import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/auth/login/login.component';
import { RegisterComponent } from './components/auth/register/register.component';
import { WelcomeComponent } from './components/welcome/welcome.component';
import { MarriageComponent } from './components/eMaticar/marriage/marriage.component';
import { ChooseServiceComponent } from './components/auth/choose-service/choose-service.component';
import { RegularOrAdminComponent } from './components/auth/regular-or-admin/regular-or-admin.component';
import { ViewMyRegistryComponent } from './components/eMaticar/view-my-registry/view-my-registry.component';
import { UserDiedComponent } from './components/eZdravstvo/user-died/user-died.component';
import { PreglediLekarComponent } from './components/eZdravstvo/pregled/pregledi-lekar/pregledi-lekar.component';
import { PreglediGradjaninComponent } from './components/eZdravstvo/pregled/pregledi-gradjanin/pregledi-gradjanin.component';
import { PregledAddComponent } from './components/eZdravstvo/pregled/pregled-add/pregled-add.component';
import { VakcinaAddComponent } from './components/eZdravstvo/vakcina/vakcina-add/vakcina-add.component';
import { VakcineComponent } from './components/eZdravstvo/vakcina/vakcine/vakcine.component';
import { VakcinaViewComponent } from './components/eZdravstvo/vakcina/vakcina-view/vakcina-view.component';
import { PregledViewComponent } from './components/eZdravstvo/pregled/pregled-view/pregled-view.component';
import { InvaliditetiAdminComponent } from './components/eAdmin/invaliditet/invaliditeti-admin/invaliditeti-admin.component';
import { InvaliditetAddComponent } from './components/eAdmin/invaliditet/invaliditet-add/invaliditet-add.component';
import { AlergijeAdminComponent } from './components/eAdmin/alergija/alergije-admin/alergije-admin.component';
import { AlergijaAddComponent } from './components/eAdmin/alergija/alergija-add/alergija-add.component';
import { KartoniComponent } from './components/eZdravstvo/karton/kartoni/kartoni.component';
import { KartonViewComponent } from './components/eZdravstvo/karton/karton-view/karton-view.component';
import { KartonEditAlergijeComponent } from './components/eZdravstvo/karton/karton-edit-alergije/karton-edit-alergije.component';
import { KartonEditInvaliditetiComponent } from './components/eZdravstvo/karton/karton-edit-invaliditeti/karton-edit-invaliditeti.component';
import { KartonViewMyComponent } from './components/eZdravstvo/karton/karton-view-my/karton-view-my.component';
import { GradjaninAddDoctorComponent } from './components/eZdravstvo/gradjanin/gradjanin-add-doctor/gradjanin-add-doctor.component';

const routes: Routes = [
  {
    path: 'Login',
    component: LoginComponent,
  },
  {
    path: 'Register',
    component: RegisterComponent,
  },

  {
    path: 'Pregledi-Lekar',
    component: PreglediLekarComponent,
  },
  {
    path: 'Pregledi-Gradjanin',
    component: PreglediGradjaninComponent,
  },
  {
    path: 'Pregled-Add',
    component: PregledAddComponent,
  },
  {
    path: 'Pregled-View/:id',
    component: PregledViewComponent,
  },
  {
    path: 'Vakcine',
    component: VakcineComponent,
  },
  {
    path: 'Vakcina-Add',
    component: VakcinaAddComponent,
  },
  {
    path: 'Vakcina-View/:id',
    component: VakcinaViewComponent,
  },
  {
    path: 'Invaliditeti-Admin',
    component: InvaliditetiAdminComponent
  },
  {
    path: 'Invaliditet-Add',
    component: InvaliditetAddComponent
  },
  {
    path: 'Alergije-Admin',
    component: AlergijeAdminComponent
  },
  {
    path: 'Alergija-Add',
    component: AlergijaAddComponent
  },
  {
    path: 'Kartoni',
    component: KartoniComponent
  },
  {
    path: 'Karton-View/:jmbg',
    component: KartonViewComponent,
  },
  {
    path: 'My-Karton',
    component: KartonViewMyComponent
  },
  {
    path: 'Karton-Edit-Alergije/:jmbg',
    component: KartonEditAlergijeComponent
  },
  {
    path: 'Karton-Edit-Invaliditeti/:jmbg',
    component: KartonEditInvaliditetiComponent
  },
  {
    path: 'Add-User-Doctor',
    component: GradjaninAddDoctorComponent
  },


  {
    path: 'Welcome',
    component: WelcomeComponent,
  },
  {
    path: 'Marriage',
    component: MarriageComponent,
  },
  {
    path: 'choose-service',
    component: ChooseServiceComponent,
  },
  {
    path: 'regular-or-admin',
    component: RegularOrAdminComponent,
  },
  {
    path: 'view-my-registry',
    component: ViewMyRegistryComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule { }
