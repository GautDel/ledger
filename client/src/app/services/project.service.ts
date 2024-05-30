import { Injectable } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { Observable, concatMap, take } from 'rxjs';
import { environment } from '../../environments/environment';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})

export class ProjectService {

  constructor(public auth: AuthService, private http: HttpClient) { }

  getProjects(sortBy: string): Observable<any> {
    return this.auth.user$.pipe(
      take(1),
      concatMap(() =>
        this.http.get(encodeURI(environment.apiURL + "/projects/sort/" + sortBy))
      )
    )
  }


  searchProjects(data: any): Observable<any> {
    return this.auth.user$.pipe(
      take(1),
      concatMap(() =>
        this.http.post(
          encodeURI(environment.apiURL + '/projects/search'),
          data,
        )
      )
    )
  }

  createProject(data: any): Observable<any> {
    return this.auth.user$.pipe(
      take(1),
      concatMap(() =>
        this.http.post(
          encodeURI(environment.apiURL + '/projects/create'),
          data,
        )
      )
    )
  }
}
