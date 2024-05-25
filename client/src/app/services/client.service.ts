import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { Observable, concatMap, take } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})

export class ClientService {

  constructor(public auth: AuthService, private http: HttpClient) { }

  getClients(sortBy: string): Observable<any> {

    return this.auth.user$.pipe(
      take(1),
      concatMap(() =>
        this.http.get(
          encodeURI(environment.apiURL + '/clients/sort/' + sortBy),
        )
      )
    )
  }

  getClient(id: number): Observable<any> {
    return this.auth.user$.pipe(
      take(1),
      concatMap(() =>
        this.http.get(
          encodeURI(environment.apiURL + '/clients/' + id)
        )
      )
    )
  }

  searchClient(data: any): Observable<any> {
    return this.auth.user$.pipe(
      take(1),
      concatMap(() =>
        this.http.post(
          encodeURI(environment.apiURL + '/clients/search'),
          data,
        )
      )
    )
  }

  createClient(data: any): Observable<any> {
    return this.auth.user$.pipe(
      take(1),
      concatMap(() => {
        return this.http.post(
          encodeURI(environment.apiURL + '/clients/create'),
          data
        )
      })
    )
  }

  updateClient(data: any, id: number): Observable<any> {
    return this.auth.user$.pipe(
      take(1),
      concatMap(() => {
        return this.http.put(
          encodeURI(environment.apiURL + "/clients/update/" + id),
          data
        )
      })
    )
  }

  updateClientStar(data: any, id: number): Observable<any> {
    return this.auth.user$.pipe(
      take(1),
      concatMap(() => {
        return this.http.put(
          encodeURI(environment.apiURL + "/clients/star/" + id),
          data
        )
      })
    )
  }

  deleteClient(id: any): Observable<any> {
    return this.auth.user$.pipe(
      take(1),
      concatMap(() => {
        return this.http.delete(
          encodeURI(environment.apiURL + '/clients/remove/' + id)
        )
      })
    )
  }
}
