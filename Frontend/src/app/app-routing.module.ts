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
    path: 'User-Died',
    component: UserDiedComponent,
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
export class AppRoutingModule {}
