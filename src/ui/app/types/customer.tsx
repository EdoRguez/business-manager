export type Customer = {
  id: number;
  firstName: string;
  lastName: string;
  identificationNumber: string;
  identificationType: string;
  phone: string;
  email: string;
}

export type CreateCustomer = {
  companyId: number;
  firstName: string;
  lastName: string;
  identificationNumber: string;
  identificationType: string;
  phone: string;
  email: string;
}

export type GetCustomer = {
  companyId: number;
  name: string;
  lastName: string;
  identificationNumber: string;
  limit: number;
  offset: number;
}

export type SearchCustomer = {
  name: string;
  lastName: string;
  identificationNumber: string;
}

export type DeleteCustomer = {
  id: number;
}
