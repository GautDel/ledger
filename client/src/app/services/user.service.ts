import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { Observable, concatMap, take } from 'rxjs';

@Injectable({
  providedIn: 'root'
})

export class UserService {

  constructor(public auth: AuthService, private http: HttpClient) { }

  getUser(): Observable<any> {
    return this.auth.user$.pipe(
        take(1),
        concatMap(() =>
          this.http.get(
            encodeURI(`https://192.168.1.15:80/user/`)
          )
        ),
      )
  }

  updateUser(data: any): Observable<any> {
    return this.auth.user$.pipe(
      take(1),
      concatMap(() => {
        return this.http.put(
          encodeURI(`https://192.168.1.15:80/user/update`),
          data
        )
      })
    )
  }
}
