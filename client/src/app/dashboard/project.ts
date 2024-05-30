import { Client } from "./client";

export interface Project {
  ID?: number,
  Name: string,
  Description: string,
  ClientID: number,
  Notes: string,
  Clients: Client[]
}
