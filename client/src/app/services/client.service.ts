import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { Observable, concatMap, take } from 'rxjs';

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
          encodeURI('https://192.168.1.15:80/clients/sort/' + sortBy),
        )
      )
    )
  }

  getClient(id: number): Observable<any> {
    return this.auth.user$.pipe(
      take(1),
      concatMap(() =>
        this.http.get(
          encodeURI('https://192.168.1.15:80/clients/' + id)
        )
      )
    )
  }

  searchClient(data: any): Observable<any> {
    return this.auth.user$.pipe(
      take(1),
      concatMap(() =>
        this.http.post(
          encodeURI('https://192.168.1.15:80/clients/search'),
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
          encodeURI('https://192.168.1.15:80/clients/create'),
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
          encodeURI("https://192.168.1.15:80/clients/update/" + id),
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
          encodeURI("https://192.168.1.15:80/clients/star/" + id),
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
          encodeURI('https://192.168.1.15:80/clients/remove/' + id)
        )
      })
    )
  }
}
