import { ApplicationConfig } from '@angular/core';
import { provideRouter } from '@angular/router';
import { provideAuth0 } from '@auth0/auth0-angular';
import { routes } from './app.routes';
import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { authHttpInterceptorFn } from '@auth0/auth0-angular';

export const appConfig: ApplicationConfig = {
  providers: [
    provideRouter(routes),
    provideHttpClient(withInterceptors([authHttpInterceptorFn])),
    provideAuth0({
      domain: 'dev-ibgzfy6dgnkl2z1i.us.auth0.com',
      clientId: 'D0TaywjhHsG8shTREGyveHD9ck9TNCbB',
      cacheLocation: 'localstorage', // Set cache location to "localstorage"
      useRefreshTokens: true,
      authorizationParams: {
        redirect_uri: "https://192.168.1.15:4200/dashboard",
        // Request this audience at user authentication time
        audience: 'https://192.168.1.15:80',
      },

      // Specify configuration for the interceptor
      httpInterceptor: {
        allowedList: [
          'https://192.168.1.15:80/*',
          'https://192.168.1.15:80/user/*',
          'https://192.168.1.15:80/invoices/*',
          'https://192.168.1.15:80/clients/search',
        ]
      }
    })
  ]
};
