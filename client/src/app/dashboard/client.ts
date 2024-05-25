export interface Client {
  ID?: number,
  FirstName: string,
  LastName: string,
  Description: string,
  Email: string,
  Phone: string,
  Address: string,
  Country: string,
  CreatedAt?: string,
  UpdatedAt?: string,
  Projects: any[],
  Starred: boolean,
}
